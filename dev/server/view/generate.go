// Copyright 2017 Google Inc.
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

package view

//go:generate embed -p view -v clientResources   -o static-client.go    -base ../../client client.js*
//go:generate embed -p view -v cssResources      -o static-css.go       css/*.css
//go:generate embed -p view -v imageResources    -o static-images.go    images/*
//go:generate embed -p view -v jsResources       -o static-js.go        js/*/*
//go:generate embed -p view -v miscResources     -o static-misc.go      misc/*
//go:generate embed -p view -v templateResources -o static-templates.go templates/*.html
