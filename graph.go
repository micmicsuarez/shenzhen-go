// Copyright 2016 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// Channel models a channel.
type Channel struct {
	Name, Type string
	Cap        int
}

// Graph describes a Go program as a graph.
type Graph struct {
	Name        string
	PackageName string
	PackagePath string
	Imports     []string
	Nodes       map[string]*Node
	Channels    map[string]*Channel
}

func (g *Graph) renderToDot(dst io.Writer) error { return dotTemplate.Execute(dst, g) }
func (g *Graph) renderToGo(dst io.Writer) error  { return goTemplate.Execute(dst, g) }

func (g *Graph) renderNodeEditor(dst io.Writer, n *Node) error {
	return nodeEditorTemplate.Execute(dst, struct {
		Graph *Graph
		Node  *Node
	}{g, n})
}

func (g *Graph) renderChannelEditor(dst io.Writer, e *Channel) error {
	return channelEditorTemplate.Execute(dst, struct {
		Graph   *Graph
		Channel *Channel
	}{g, e})
}

func (g *Graph) handleChannelRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL)
	nm := strings.TrimPrefix(r.URL.Path, "/channel/")

	e, found := g.Channels[nm]
	if nm != "new" && !found {
		http.Error(w, fmt.Sprintf("Channel %q not found", nm), http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		if e == nil {
			e = new(Channel)
		}

		// Parse...
		if err := r.ParseForm(); err != nil {
			log.Printf("Could not parse form: %v", err)
			http.Error(w, "Could not parse", http.StatusBadRequest)
			return
		}

		// ...Validate...
		nn := r.FormValue("Name")
		if !identifierRE.MatchString(nn) {
			msg := fmt.Sprintf("Invalid identifier %q !~ %q", nn, identifierRE)
			log.Printf(msg)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		ci, err := strconv.Atoi(r.FormValue("Cap"))
		if err != nil {
			log.Printf("Capacity is not an integer: %v", err)
			http.Error(w, "Capacity is not an integer", http.StatusBadRequest)
			return
		}
		if ci < 0 {
			log.Printf("Must specify nonnegative capacity [%d < 0]", ci)
			http.Error(w, "Capacity must be non-negative", http.StatusBadRequest)
			return
		}

		// ...update...
		e.Type = r.FormValue("Type")
		e.Cap = ci

		// Do name changes last since they cause a redirect.
		if nn == e.Name {
			break
		}
		delete(g.Channels, e.Name)
		e.Name = nn
		g.Channels[nn] = e
		http.Redirect(w, r, "/channel/"+nn, http.StatusSeeOther) // should cause GET
		return
	}

	if err := g.renderChannelEditor(w, e); err != nil {
		log.Printf("Could not render source editor: %v", err)
		http.Error(w, "Could not render source editor", http.StatusInternalServerError)
		return
	}
	return
}

func (g *Graph) handleNodeRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL)
	nm := strings.TrimPrefix(r.URL.Path, "/node/")
	n, found := g.Nodes[nm]
	if !found {
		http.Error(w, fmt.Sprintf("Node %q not found", nm), http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		// Update the node.
		if err := r.ParseForm(); err != nil {
			log.Printf("Could not parse form: %v", err)
			http.Error(w, "Could not parse", http.StatusBadRequest)
			return
		}

		nm := strings.TrimSpace(r.FormValue("Name"))
		if nm == "" {
			log.Printf("Name invalid [%q == \"\"]", nm)
			http.Error(w, "Name invalid", http.StatusBadRequest)
			return
		}

		n.Wait = (r.FormValue("Wait") == "on")
		n.Code = r.FormValue("Code")

		if err := n.updateChans(g.Channels); err != nil {
			log.Printf("Unable to extract channels used from code: %v", err)
			http.Error(w, "Unable to extract channels used from code", http.StatusBadRequest)
			return
		}

		if nm == n.Name {
			break
		}

		delete(g.Nodes, n.Name)
		n.Name = nm
		g.Nodes[nm] = n

		http.Redirect(w, r, "/node/"+nm, http.StatusSeeOther) // should cause GET
		return
	}

	if err := g.renderNodeEditor(w, n); err != nil {
		log.Printf("Could not render source editor: %v", err)
		http.Error(w, "Could not render source editor", http.StatusInternalServerError)
		return
	}
	return
}

func (g *Graph) outputDotSrc(w http.ResponseWriter) {
	h := w.Header()
	h.Set("Content-Type", "text/vnd.graphviz")
	if err := g.renderToDot(w); err != nil {
		log.Printf("Could not render to dot: %v", err)
		http.Error(w, "Could not render to dot", http.StatusInternalServerError)
	}
}

func (g *Graph) writeGoSrc(w io.Writer) error {
	buf := &bytes.Buffer{}
	if err := g.renderToGo(buf); err != nil {
		return err
	}
	return goimports(w, buf)
}

func (g *Graph) outputGoSrc(w http.ResponseWriter) {
	h := w.Header()
	h.Set("Content-Type", "text/golang")
	if err := g.writeGoSrc(w); err != nil {
		log.Printf("Could not render to Go: %v", err)
		http.Error(w, "Could not render to Go", http.StatusInternalServerError)
	}
}

func (g *Graph) outputJSON(w http.ResponseWriter) {
	h := w.Header()
	h.Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	if err := enc.Encode(g); err != nil {
		log.Printf("Could not encode JSON: %v", err)
		http.Error(w, "Could not encode JSON", http.StatusInternalServerError)
		return
	}
}

func (g *Graph) handleRootRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL)
	q := r.URL.Query()
	if _, t := q["dot"]; t {
		g.outputDotSrc(w)
		return
	}
	if _, t := q["go"]; t {
		g.outputGoSrc(w)
		return
	}
	if _, t := q["json"]; t {
		g.outputJSON(w)
		return
	}
	if _, t := q["run"]; t {
		if err := g.saveBuildAndRun(); err != nil {
			log.Printf("Failed to save, build, run: %v", err)
		}
	}

	var dot, svg bytes.Buffer
	if err := g.renderToDot(&dot); err != nil {
		log.Printf("Could not render to dot: %v", err)
		http.Error(w, "Could not render to dot", http.StatusInternalServerError)
	}
	if err := dotToSVG(&svg, &dot); err != nil {
		log.Printf("Could not render dot to SVG: %v", err)
		http.Error(w, "Could not render dot to SVG", http.StatusInternalServerError)
	}
	if err := rootTemplate.Execute(w, template.HTML(svg.String())); err != nil {
		log.Printf("Could not execute root template: %v", err)
		http.Error(w, "Could not execute root template", http.StatusInternalServerError)
	}
}

func (g *Graph) saveGoSrc() error {
	gopath, ok := os.LookupEnv("GOPATH")
	if !ok || gopath == "" {
		return errors.New("cannot use $GOPATH; empty or undefined")
	}
	pp := filepath.Join(gopath, "src", g.PackagePath)
	if err := os.Mkdir(pp, os.FileMode(0755)); err != nil {
		log.Printf("Could not make path %q, continuing: %v", pp, err)
	}
	mp := filepath.Join(pp, "main.go")
	f, err := os.Create(mp)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := g.writeGoSrc(f); err != nil {
		return err
	}
	return f.Close()
}

func (g *Graph) build() error {
	return exec.Command(`go`, `build`, g.PackagePath).Run()
}

func (g *Graph) saveBuildAndRun() error {
	if err := g.saveGoSrc(); err != nil {
		return err
	}
	if err := g.build(); err != nil {
		return err
	}
	// TODO: Be less lazy about the output binary path
	return open("./" + g.PackageName)
}
