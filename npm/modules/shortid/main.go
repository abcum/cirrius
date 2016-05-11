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

package shortid

import (
	"github.com/SKAhack/go-shortid"
	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
)

func init() {

	orbit.Add("shortid", func(ctx *orbit.Orbit) (otto.Value, error) {

		generator := shortid.Generator()

		return ctx.ToValue(map[string]interface{}{
			"generate": func(call otto.FunctionCall) otto.Value {
				value, _ := ctx.ToValue(generator.Generate())
				return value
			},
			"seed": func(call otto.FunctionCall) otto.Value {
				arg1, _ := call.Argument(0).ToFloat()
				generator.SetSeed(arg1)
				return otto.UndefinedValue()
			},
			"characters": func(call otto.FunctionCall) otto.Value {
				arg1 := call.Argument(0).String()
				generator.SetCharacters(arg1)
				value, _ := ctx.ToValue(arg1)
				return value
			},
		})

	})

}
