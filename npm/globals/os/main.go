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

package os

import (
	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
)

func init() {

	orbit.OnInit(func(ctx *orbit.Orbit) {

		os := map[string]interface{}{

			"EOL": "\n",

			"arch": func(call otto.FunctionCall) (val otto.Value) {
				val, _ = otto.ToValue("x64")
				return
			},

			"constants": []string{},

			"cpus": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"endianness": func(call otto.FunctionCall) (val otto.Value) {
				val, _ = otto.ToValue("LE")
				return
			},

			"freemem": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"homedir": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"hostname": func(call otto.FunctionCall) (val otto.Value) {
				val, _ = otto.ToValue("cirrius.local")
				return
			},

			"loadavg": func(call otto.FunctionCall) (val otto.Value) {
				val, _ = otto.ToValue(0)
				return
			},

			"networkInterfaces": func(call otto.FunctionCall) (val otto.Value) {
				val, _ = otto.ToValue("linux")
				return
			},

			"release": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"tmpdir": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"totalmem": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"type": func(call otto.FunctionCall) (val otto.Value) {
				val, _ = otto.ToValue("Linux")
				return
			},

			"userInfo": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},
		}

		ctx.Def("os", os)

	})

}
