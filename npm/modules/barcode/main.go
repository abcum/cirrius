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

package uuid

import (
	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
)

func init() {

	orbit.Add("barcode", func(ctx *orbit.Orbit) (otto.Value, error) {
		return ctx.ToValue(map[string]interface{}{
			"codabar": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},
			"code128": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},
			"code39": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},
			"datamatrix": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},
			"ean": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},
			"qr": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},
			"twooffive": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},
		})
	})

}