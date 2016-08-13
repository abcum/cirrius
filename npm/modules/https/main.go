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

package https

import (
	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
)

func init() {

	orbit.Add("https", func(ctx *orbit.Orbit) (otto.Value, error) {
		return ctx.ToValue(map[string]interface{}{

			"get": func(call otto.FunctionCall) (val otto.Value) {
				return otto.UndefinedValue()
			},

			"request": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},
		})
	})

}
