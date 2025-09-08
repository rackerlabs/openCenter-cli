// Copyright 2025 Victor Palma <victor.palma@rackspace.com>
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

package provision

import (
	"embed"
	"sync"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

//go:embed all:templates
var templatesFS embed.FS

// Templates holds the parsed templates, ready for execution.
var (
	Templates *template.Template
	once      sync.Once
	initErr   error
)

// Init parses the embedded templates and stores them in the Templates variable.
// It uses a sync.Once to ensure that the templates are parsed only once.
//
// Outputs:
//   - error: An error if one occurred during template parsing.
func Init() error {
	once.Do(func() {
		Templates, initErr = template.New("").Funcs(sprig.TxtFuncMap()).ParseFS(templatesFS, "templates/*.tmpl")
	})
	return initErr
}

func init() {
	if err := Init(); err != nil {
		panic(err)
	}
}
