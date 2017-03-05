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

		orb := options.DataFrame().Get("orb").(*orbit.Orbit)

		data, _, err := orb.File(name, "hbs")
		if err != nil {
			msg := fmt.Sprintf("Unable to find handlebars template '%s'.", name)
			panic(orb.MakeCustomError("Error", msg))
		}

		code := fmt.Sprintf("%s", data)

		tmpl, err := raymond.Parse(code)
		if err != nil {
			msg := fmt.Sprintf("Unable to parse handlebars template '%s'.", name)
			panic(orb.MakeCustomError("Error", msg))
		}

		priv := options.DataFrame()

		html, err := tmpl.ExecWith(options.Ctx(), priv)
		if err != nil {
			msg := fmt.Sprintf("Unable to render handlebars template '%s'.", name)
			panic(orb.MakeCustomError("Error", msg))
		}

		return raymond.SafeString(html)

	})

	orbit.OnInit(func(orb *orbit.Orbit) {

		session := orb.Context().Value("fibre").(*fibre.Context)

		context := map[string]interface{}{

			"response": map[string]interface{}{
				"body": session.Response(),
			},

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
				"body":   session.Request().Body,
				"ip":     session.IP().String(),
			},

			"success": func() (val otto.Value) {
				session.Code(200)
				panic(nil)
				return
			},

			"failure": func() (val otto.Value) {
				session.Code(400)
				panic(nil)
				return
			},

			// The following methods enable sending the
			// response data using different encoding
			// methods, with an http status code 200.

			"xml": func(data interface{}) (val otto.Value) {
				session.XML(200, data)
				panic(nil)
				return
			},

			"text": func(data interface{}) (val otto.Value) {
				session.Text(200, data)
				panic(nil)
				return
			},

			"html": func(data interface{}) (val otto.Value) {
				session.HTML(200, data)
				panic(nil)
				return
			},

			"json": func(data interface{}) (val otto.Value) {
				session.JSON(200, data)
				panic(nil)
				return
			},

			"cbor": func(data interface{}) (val otto.Value) {
				session.CBOR(200, data)
				panic(nil)
				return
			},

			"pack": func(data interface{}) (val otto.Value) {
				session.PACK(200, data)
				panic(nil)
				return
			},

			"send": func(data interface{}) (val otto.Value) {
				session.Send(200, data)
				panic(nil)
				return
			},

			// Status enables defining the http status
			// code, with a chained method to send the
			// data using different encoding types.

			"status": func(code int) (val otto.Value) {
				val, _ = orb.ToValue(map[string]interface{}{
					"xml": func(data interface{}) (val otto.Value) {
						session.XML(code, data)
						panic(nil)
						return
					},
					"text": func(data interface{}) (val otto.Value) {
						session.Text(code, data)
						panic(nil)
						return
					},
					"html": func(data interface{}) (val otto.Value) {
						session.HTML(code, data)
						panic(nil)
						return
					},
					"json": func(data interface{}) (val otto.Value) {
						session.JSON(code, data)
						return
					},
					"pack": func(data interface{}) (val otto.Value) {
						session.PACK(code, data)
						panic(nil)
						return
					},
					"send": func(data interface{}) (val otto.Value) {
						session.Send(code, data)
						panic(nil)
						return
					},
				})
				return
			},

			// Render compiles and outputs a handlebars
			// template, and additional partial templates
			// included within the main template file.

			"render": func(call otto.FunctionCall) (val otto.Value) {

				if len(call.ArgumentList) >= 3 {
					msg := fmt.Sprintf("expected 1 or 2 argument(s); got %d.", len(call.ArgumentList))
					panic(orb.MakeRangeError(msg))
				}

				file, _ := call.Argument(0).ToString()
				vars, _ := call.Argument(1).Export()

				data, _, err := orb.File(file, "hbs")
				if err != nil {
					msg := fmt.Sprintf("Unable to find handlebars template '%s'.", file)
					panic(orb.MakeCustomError("Error", msg))
				}

				code := fmt.Sprintf("%s", data)

				tmpl, err := raymond.Parse(code)
				if err != nil {
					msg := fmt.Sprintf("Unable to parse handlebars template '%s'.", file)
					panic(orb.MakeCustomError("Error", msg))
				}

				priv := raymond.NewDataFrame()
				priv.Set("orb", orb)

				html, err := tmpl.ExecWith(vars, priv)
				if err != nil {
					msg := fmt.Sprintf("Unable to render handlebars template '%s'.", file)
					panic(orb.MakeCustomError("Error", msg))
				}

				session.HTML(200, html)
				panic(nil)

				return

			},
		}

		orb.Def("context", context)

		orb.Def("cirrius", context)

	})

}
