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

package colour

import (
	"github.com/abcum/orbit"
)

func init() {

	orbit.Add("colour", func(ctx *orbit.Orbit) (otto.Value, error) {
		return ctx.ToValue(map[string]interface{}{
			"rgb": func(call otto.FunctionCall) otto.Value {
				r := call.Argument(0).ToInteger()
				g := call.Argument(1).ToInteger()
				b := call.Argument(2).ToInteger()
				return otto.UndefinedValue()
			},
			"lab": func(call otto.FunctionCall) otto.Value {
				l := call.Argument(0).ToInteger()
				a := call.Argument(1).ToInteger()
				b := call.Argument(2).ToInteger()
				return otto.UndefinedValue()
			},
			"hsv": func(call otto.FunctionCall) otto.Value {
				h := call.Argument(0).ToInteger()
				s := call.Argument(1).ToInteger()
				v := call.Argument(2).ToInteger()
				return otto.UndefinedValue()
			},
			"xyz": func(call otto.FunctionCall) otto.Value {
				x := call.Argument(0).ToInteger()
				y := call.Argument(1).ToInteger()
				z := call.Argument(2).ToInteger()
				return otto.UndefinedValue()
			},
			"cmyk": func(call otto.FunctionCall) otto.Value {
				c := call.Argument(0).ToInteger()
				m := call.Argument(1).ToInteger()
				y := call.Argument(2).ToInteger()
				k := call.Argument(3).ToInteger()
				return otto.UndefinedValue()
			},
		})
	})

}
