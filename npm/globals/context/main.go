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

package output

import (
	"github.com/abcum/fibre"
	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
)

func init() {

	orbit.OnInit(func(ctx *orbit.Orbit) {

		session := ctx.Vars["fibre"].(*fibre.Context)

		context := map[string]interface{}{

			"request": map[string]interface{}{
				"user":  session.Request().URL().User,
				"pass":  session.Request().URL().Pass,
				"host":  session.Request().URL().Host,
				"path":  session.Request().URL().Path,
				"query": session.Request().URL().Query,
				"head":  session.Request().Header(),
				"body":  session.Body(),
				"ip":    session.IP(),
			},

			"success": func(call otto.FunctionCall) otto.Value {
				ctx.Interrupt <- func() {}
				session.Code(200)
				return otto.UndefinedValue()
			},

			"failure": func(call otto.FunctionCall) otto.Value {
				ctx.Interrupt <- func() {}
				session.Code(400)
				return otto.UndefinedValue()
			},

			"display": map[string]interface{}{

				"xml": func(call otto.FunctionCall) otto.Value {

					if len(call.ArgumentList) == 1 {
						data, _ := call.Argument(0).Export()
						ctx.Interrupt <- func() {}
						session.XML(200, data)
					}

					if len(call.ArgumentList) == 2 {
						code, _ := call.Argument(0).ToInteger()
						data, _ := call.Argument(1).Export()
						ctx.Interrupt <- func() {}
						session.XML(int(code), data)
					}

					if len(call.ArgumentList) >= 3 {
						ctx.Interrupt <- func() {}
						session.Code(500)
					}

					return otto.UndefinedValue()

				},

				"text": func(call otto.FunctionCall) otto.Value {

					if len(call.ArgumentList) == 1 {
						data, _ := call.Argument(0).Export()
						ctx.Interrupt <- func() {}
						session.Text(200, data)
					}

					if len(call.ArgumentList) == 2 {
						code, _ := call.Argument(0).ToInteger()
						data, _ := call.Argument(1).Export()
						ctx.Interrupt <- func() {}
						session.Text(int(code), data)
					}

					if len(call.ArgumentList) >= 3 {
						ctx.Interrupt <- func() {}
						session.Code(500)
					}

					return otto.UndefinedValue()

				},

				"html": func(call otto.FunctionCall) otto.Value {

					if len(call.ArgumentList) == 1 {
						data, _ := call.Argument(0).Export()
						ctx.Interrupt <- func() {}
						session.HTML(200, data)
					}

					if len(call.ArgumentList) == 2 {
						code, _ := call.Argument(0).ToInteger()
						data, _ := call.Argument(1).Export()
						ctx.Interrupt <- func() {}
						session.HTML(int(code), data)
					}

					if len(call.ArgumentList) >= 3 {
						ctx.Interrupt <- func() {}
						session.Code(500)
					}

					return otto.UndefinedValue()

				},

				"json": func(call otto.FunctionCall) otto.Value {

					if len(call.ArgumentList) == 1 {
						data, _ := call.Argument(0).Export()
						ctx.Interrupt <- func() {}
						session.JSON(200, data)
					}

					if len(call.ArgumentList) == 2 {
						code, _ := call.Argument(0).ToInteger()
						data, _ := call.Argument(1).Export()
						ctx.Interrupt <- func() {}
						session.JSON(int(code), data)
					}

					if len(call.ArgumentList) >= 3 {
						ctx.Interrupt <- func() {}
						session.Code(500)
					}

					return otto.UndefinedValue()

				},

				"pack": func(call otto.FunctionCall) otto.Value {

					if len(call.ArgumentList) == 1 {
						data, _ := call.Argument(0).Export()
						ctx.Interrupt <- func() {}
						session.PACK(200, data)
					}

					if len(call.ArgumentList) == 2 {
						code, _ := call.Argument(0).ToInteger()
						data, _ := call.Argument(1).Export()
						ctx.Interrupt <- func() {}
						session.PACK(int(code), data)
					}

					if len(call.ArgumentList) >= 3 {
						ctx.Interrupt <- func() {}
						session.Code(500)
					}

					return otto.UndefinedValue()

				},
			},
		}

		ctx.Def("context", context)

	})

}
