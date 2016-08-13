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

	orbit.Add("path", func(ctx *orbit.Orbit) (otto.Value, error) {
		return ctx.ToValue(map[string]interface{}{

			"basename": func(call otto.FunctionCall) (val otto.Value) {
				if len(call.ArgumentList) == 1 {
					if !call.Argument(0).IsString() {
						return ctx.MakeTypeError("`path` must be a string")
					}
					arg1, _ := call.Argument(0).ToString()
					val, _ = otto.ToValue(filepath.Base(arg1))
					return
				}
				if len(call.ArgumentList) == 2 {
					if !call.Argument(0).IsString() {
						return ctx.MakeTypeError("`path` must be a string")
					}
					if !call.Argument(1).IsString() {
						return ctx.MakeTypeError("`ext` must be a string")
					}
					arg1, _ := call.Argument(0).ToString()
					arg2, _ := call.Argument(1).ToString()
					val, _ = otto.ToValue(filepath.Base(strings.TrimRight(arg1, arg2)))
					return
				}
				return ctx.MakeSyntaxError("Wrong argument length")
			},

			"delimiter": ":",

			"dirname": func(call otto.FunctionCall) (val otto.Value) {
				if len(call.ArgumentList) == 1 {
					if !call.Argument(0).IsString() {
						return ctx.MakeTypeError("`path` must be a string")
					}
					arg1, _ := call.Argument(0).ToString()
					val, _ = otto.ToValue(filepath.Dir(arg1))
					return
				}
				return ctx.MakeSyntaxError("Wrong argument length")
			},

			"extname": func(call otto.FunctionCall) (val otto.Value) {
				if len(call.ArgumentList) == 1 {
					if !call.Argument(0).IsString() {
						return ctx.MakeTypeError("`path` must be a string")
					}
					arg1, _ := call.Argument(0).ToString()
					val, _ = otto.ToValue(filepath.Ext(arg1))
					return
				}
				return ctx.MakeSyntaxError("Wrong argument length")
			},

			"format": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue() // TODO need to implement this method
			},

			"isAbsolute": func(call otto.FunctionCall) (val otto.Value) {
				if len(call.ArgumentList) == 1 {
					if !call.Argument(0).IsString() {
						return ctx.MakeTypeError("`path` must be a string")
					}
					arg1, _ := call.Argument(0).ToString()
					val, _ = otto.ToValue(filepath.IsAbs(arg1))
					return
				}
				return ctx.MakeSyntaxError("Wrong argument length")
			},

			"join": func(call otto.FunctionCall) (val otto.Value) {
				var args []string
				for _, v := range call.ArgumentList {
					if !v.IsString() {
						return ctx.MakeTypeError("`path` must be a string")
					}
					str, _ := v.ToString()
					args = append(args, str)
				}
				val, _ = otto.ToValue(filepath.Join(args...))
				return
			},

			"normalize": func(call otto.FunctionCall) (val otto.Value) {
				if len(call.ArgumentList) == 1 {
					if !call.Argument(0).IsString() {
						return ctx.MakeTypeError("`path` must be a string")
					}
					arg1, _ := call.Argument(0).ToString()
					val, _ = otto.ToValue(filepath.Clean(arg1))
					return
				}
				return ctx.MakeSyntaxError("Wrong argument length")
			},

			"parse": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue() // TODO need to implement this method
			},

			"posix": otto.UndefinedValue(),

			"relative": func(call otto.FunctionCall) (val otto.Value) {
				if len(call.ArgumentList) == 2 {
					if !call.Argument(0).IsString() {
						return ctx.MakeTypeError("`from` must be a string")
					}
					if !call.Argument(1).IsString() {
						return ctx.MakeTypeError("`to` must be a string")
					}
					arg1, _ := call.Argument(0).ToString()
					arg2, _ := call.Argument(1).ToString()
					rel, _ := filepath.Rel(arg1, arg2)
					val, _ = otto.ToValue(rel)
					return
				}
				return ctx.MakeSyntaxError("Wrong argument length")
			},

			"resolve": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue() // TODO need to implement this method
			},

			"sep": "/",

			"win32": otto.UndefinedValue(),
		})

	})

}
