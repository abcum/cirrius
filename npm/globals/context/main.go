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

package context

import (
	"fmt"

	"github.com/abcum/fibre"
	"github.com/abcum/orbit"
	"github.com/aymerick/raymond"
	"github.com/robertkrimen/otto"
)

func init() {

	raymond.RegisterHelper("partial", func(name string, options *raymond.Options) raymond.SafeString {

		ctx := options.DataFrame().Get("ctx").(*orbit.Orbit)

		session := ctx.Vars["fibre"].(*fibre.Context)

		data, _, err := ctx.File(name, "hbs")
		if err != nil {
			session.Code(500)
			ctx.Quit(err)
		}

		code := fmt.Sprintf("%s", data)

		tmpl, err := raymond.Parse(code)
		if err != nil {
			msg := fmt.Sprintf("Unable to parse handlebars template '%s'.", name)
			err := ctx.MakeCustomError("Error", msg)
			session.Code(500)
			ctx.Quit(err)
		}

		priv := options.DataFrame()

		html, err := tmpl.ExecWith(options.Ctx(), priv)
		if err != nil {
			msg := fmt.Sprintf("Unable to render handlebars template '%s'.", name)
			err := ctx.MakeCustomError("Error", msg)
			session.Code(500)
			ctx.Quit(err)
		}

		return raymond.SafeString(html)

	})

	orbit.OnInit(func(ctx *orbit.Orbit) {

		session := ctx.Vars["fibre"].(*fibre.Context)

		context := map[string]interface{}{

			"request": map[string]interface{}{
				"socket": session.Request().Header().Get("Upgrade") == "websocket",
				"origin": session.Request().Header().Get("Origin"),
				"method": session.Request().Method,
				"user":   session.Request().URL().User,
				"pass":   session.Request().URL().Pass,
				"host":   session.Request().URL().Host,
				"path":   session.Request().URL().Path,
				"query":  session.Request().URL().Query,
				"head":   session.Head(),
				"body":   session.Body(),
				"ip":     session.IP(),
			},

			"success": func(call otto.FunctionCall) otto.Value {
				session.Code(200)
				ctx.Quit(nil)
				return otto.UndefinedValue()
			},

			"failure": func(call otto.FunctionCall) otto.Value {
				session.Code(400)
				ctx.Quit(nil)
				return otto.UndefinedValue()
			},

			"status": func(call otto.FunctionCall) otto.Value {

				if len(call.ArgumentList) == 1 {
					code, _ := call.Argument(0).ToInteger()
					session.Code(int(code))
					ctx.Quit(nil)
				}

				if len(call.ArgumentList) >= 2 {
					msg := fmt.Sprintf("Too many arguments to 'context.status'. Expected 1 or 2 arguments, but received %d.", len(call.ArgumentList))
					err := ctx.MakeCustomError("Error", msg)
					session.Code(500)
					ctx.Quit(err)
				}

				return otto.UndefinedValue()

			},

			"display": map[string]interface{}{

				"xml": func(call otto.FunctionCall) otto.Value {

					if len(call.ArgumentList) == 1 {
						data, _ := call.Argument(0).Export()
						session.XML(200, data)
						ctx.Quit(nil)
					}

					if len(call.ArgumentList) == 2 {
						code, _ := call.Argument(0).ToInteger()
						data, _ := call.Argument(1).Export()
						session.XML(int(code), data)
						ctx.Quit(nil)
					}

					if len(call.ArgumentList) >= 3 {
						msg := fmt.Sprintf("Too many arguments to 'context.display.xml'. Expected 1 or 2 arguments, but received %d.", len(call.ArgumentList))
						err := ctx.MakeCustomError("Error", msg)
						session.Code(500)
						ctx.Quit(err)
					}

					return otto.UndefinedValue()

				},

				"text": func(call otto.FunctionCall) otto.Value {

					if len(call.ArgumentList) == 1 {
						data, _ := call.Argument(0).Export()
						session.Text(200, data)
						ctx.Quit(nil)
					}

					if len(call.ArgumentList) == 2 {
						code, _ := call.Argument(0).ToInteger()
						data, _ := call.Argument(1).Export()
						session.Text(int(code), data)
						ctx.Quit(nil)
					}

					if len(call.ArgumentList) >= 3 {
						msg := fmt.Sprintf("Too many arguments to 'context.display.text'. Expected 1 or 2 arguments, but received %d.", len(call.ArgumentList))
						err := ctx.MakeCustomError("Error", msg)
						session.Code(500)
						ctx.Quit(err)
					}

					return otto.UndefinedValue()

				},

				"html": func(call otto.FunctionCall) otto.Value {

					if len(call.ArgumentList) == 1 {
						data, _ := call.Argument(0).Export()
						session.HTML(200, data)
						ctx.Quit(nil)
					}

					if len(call.ArgumentList) == 2 {
						code, _ := call.Argument(0).ToInteger()
						data, _ := call.Argument(1).Export()
						session.HTML(int(code), data)
						ctx.Quit(nil)
					}

					if len(call.ArgumentList) >= 3 {
						msg := fmt.Sprintf("Too many arguments to 'context.display.html'. Expected 1 or 2 arguments, but received %d.", len(call.ArgumentList))
						err := ctx.MakeCustomError("Error", msg)
						session.Code(500)
						ctx.Quit(err)
					}

					return otto.UndefinedValue()

				},

				"json": func(call otto.FunctionCall) otto.Value {

					if len(call.ArgumentList) == 1 {
						data, _ := call.Argument(0).Export()
						session.JSON(200, data)
						ctx.Quit(nil)
					}

					if len(call.ArgumentList) == 2 {
						code, _ := call.Argument(0).ToInteger()
						data, _ := call.Argument(1).Export()
						session.JSON(int(code), data)
						ctx.Quit(nil)
					}

					if len(call.ArgumentList) >= 3 {
						msg := fmt.Sprintf("Too many arguments to 'context.display.json'. Expected 1 or 2 arguments, but received %d.", len(call.ArgumentList))
						err := ctx.MakeCustomError("Error", msg)
						session.Code(500)
						ctx.Quit(err)
					}

					return otto.UndefinedValue()

				},

				"pack": func(call otto.FunctionCall) otto.Value {

					if len(call.ArgumentList) == 1 {
						data, _ := call.Argument(0).Export()
						session.PACK(200, data)
						ctx.Quit(nil)
					}

					if len(call.ArgumentList) == 2 {
						code, _ := call.Argument(0).ToInteger()
						data, _ := call.Argument(1).Export()
						session.PACK(int(code), data)
						ctx.Quit(nil)
					}

					if len(call.ArgumentList) >= 3 {
						msg := fmt.Sprintf("Too many arguments to 'context.display.pack'. Expected 1 or 2 arguments, but received %d.", len(call.ArgumentList))
						err := ctx.MakeCustomError("Error", msg)
						session.Code(500)
						ctx.Quit(err)
					}

					return otto.UndefinedValue()

				},
			},

			"render": func(call otto.FunctionCall) otto.Value {

				if len(call.ArgumentList) >= 3 {
					msg := fmt.Sprintf("Too many arguments to 'context.render'. Expected 1 or 2 arguments, but received %d.", len(call.ArgumentList))
					err := ctx.MakeCustomError("Error", msg)
					session.Code(500)
					ctx.Quit(err)
				}

				file, _ := call.Argument(0).ToString()
				vars, _ := call.Argument(1).Export()

				data, _, err := ctx.File(file, "hbs")
				if err != nil {
					session.Code(500)
					ctx.Quit(err)
				}

				code := fmt.Sprintf("%s", data)

				tmpl, err := raymond.Parse(code)
				if err != nil {
					session.Code(500)
					ctx.Quit(err)
				}

				priv := raymond.NewDataFrame()
				priv.Set("ctx", ctx)

				html, err := tmpl.ExecWith(vars, priv)
				if err != nil {
					session.Code(500)
					ctx.Quit(err)
				}

				session.HTML(200, html)
				ctx.Quit(nil)

				return otto.UndefinedValue()

			},
		}

		ctx.Def("context", context)

		ctx.Def("cirrius", context)

	})

}
