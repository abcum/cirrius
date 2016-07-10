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

package process

import (
	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
)

func init() {

	orbit.OnInit(func(ctx *orbit.Orbit) {

		process := map[string]interface{}{

			"abort": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"arch": "x64",

			"argv": []string{},

			"chdir": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"config": map[string]interface{}{},

			"connected": false,

			"cpuUsage": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"cwd": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"disconnect": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"env": map[string]interface{}{},

			"emitWarning": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"execArgv": []string{},

			"execPath": otto.UndefinedValue(),

			"exit": func(call otto.FunctionCall) otto.Value {
				ctx.Interrupt <- func() {}
				return otto.UndefinedValue()
			},

			"exitCode": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"getegid": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"geteuid": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"getgid": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"getgroups": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"getuid": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"hrtime": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"initgroups": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"kill": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"mainModule": "",

			"memoryUsage": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"nextTick": func(call otto.FunctionCall) otto.Value {
				obj, _ := ctx.Get("setImmediate")
				obj.Call(obj, call.ArgumentList)
				return otto.UndefinedValue()
			},

			"pid": 0,

			"platform": "linux",

			"release": map[string]interface{}{
				"name":       "Cirrius",
				"lts":        otto.UndefinedValue(),
				"libUrl":     otto.UndefinedValue(),
				"sourceUrl":  otto.UndefinedValue(),
				"headersUrl": otto.UndefinedValue(),
			},

			"setegid": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"seteuid": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"setgid": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"setgroups": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"setuid": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"stderr": otto.UndefinedValue(),

			"stdin": otto.UndefinedValue(),

			"stdout": otto.UndefinedValue(),

			"title": "",

			"umask": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"uptime": func(call otto.FunctionCall) otto.Value {
				return otto.UndefinedValue()
			},

			"version": "1.0.0",

			"versions": map[string]interface{}{},
		}

		ctx.Def("process", process)

	})

}
