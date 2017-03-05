// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this info except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package path

import (
	"strings"

	"path/filepath"

	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
)

func init() {

	orbit.Add("path", func(orb *orbit.Orbit) (otto.Value, error) {
		return orb.ToValue(map[string]interface{}{

			"basename": func(path string, ext ...string) (val otto.Value) {
				if len(ext) > 0 {
					val, _ = otto.ToValue(filepath.Base(strings.TrimRight(path, ext[0])))
					return
				}
				val, _ = otto.ToValue(filepath.Base(path))
				return
			},

			"delimiter": ":",

			"dirname": func(path string) (val otto.Value) {
				val, _ = otto.ToValue(filepath.Dir(path))
				return
			},

			"extname": func(path string) (val otto.Value) {
				val, _ = otto.ToValue(filepath.Ext(path))
				return
			},

			"format": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue() // TODO need to implement this method
			},

			"isAbsolute": func(path string) (val otto.Value) {
				val, _ = otto.ToValue(filepath.IsAbs(path))
				return
			},

			"join": func(paths ...string) (val otto.Value) {
				val, _ = otto.ToValue(filepath.Join(paths...))
				return
			},

			"normalize": func(path string) (val otto.Value) {
				val, _ = otto.ToValue(filepath.Clean(path))
				return
			},

			"parse": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue() // TODO need to implement this method
			},

			"posix": otto.UndefinedValue(),

			"relative": func(from, to string) (val otto.Value) {
				rel, _ := filepath.Rel(from, to)
				val, _ = otto.ToValue(rel)
				return
			},

			"resolve": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue() // TODO need to implement this method
			},

			"sep": "/",

			"win32": otto.UndefinedValue(),
		})

	})

}
