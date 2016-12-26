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

package graph

import (
	"text/template"
)

const (
	dotTemplateSrc = `digraph {
	graph[rankdir="UD",fontname="Go"];
	node[shape=box,fontname="Go"];
	{{range .Nodes}}
	"{{.Name}}"[URL="/node/{{.Name}}"];{{end}}
	{{range .Channels}}
	"{{.Name}}" [URL="/channel/{{.Name}}",shape=oval,fontname="Go Mono"];{{end}}
	{{range $n := .Nodes}}{{range $.DeclaredChannels .ChannelsRead}}
	"{{.}}" -> "{{$n.Name}}";{{end}}{{range $.DeclaredChannels .ChannelsWritten}}
	"{{$n.Name}}" -> "{{.}}";{{end}}{{end}}
}`

	goTemplateSrc = `// This binary was automatically generated by Shenzhen Go.
package main

import (
	{{range .Imports}}"{{.}}"
	{{end}}"sync"
)

func main() {
	{{range .Channels}}
	{{.Name}} := make(chan {{.Type}}, {{.Cap}}){{end}}
	var wg sync.WaitGroup
	{{range .Nodes}}
	{{if .Wait}}wg.Add(1)
	{{end}}go func() {
		{{if .Wait}}defer wg.Done()
		{{end}}{{.Impl}}
	}()
	{{end}}
	wg.Wait()
}`
)

var (
	dotTemplate = template.Must(template.New("dot").Parse(dotTemplateSrc))
	goTemplate  = template.Must(template.New("golang").Parse(goTemplateSrc))
)