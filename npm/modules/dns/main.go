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

package dns

import (
	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
)

func init() {

	orbit.Add("dns", func(ctx *orbit.Orbit) (otto.Value, error) {
		return ctx.ToValue(map[string]interface{}{

			"getServers": func(call otto.FunctionCall) (val otto.Value) {
				val, _ = otto.ToValue([]string{"8.8.8.8", "8.8.4.4"})
				return
			},

			"lookup": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"lookupService": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"resolve": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"resolve4": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"resolve6": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"resolveCname": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"resolveMx": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"resolveNaptr": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"resolveNs": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"resolveSoa": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"resolveSrv": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"resolvePtr": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"resolveTxt": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"reverse": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"setServers": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},
		})

	})

}
