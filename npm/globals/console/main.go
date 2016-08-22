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

package console

import (
	"log"
	"path"
	"time"

	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
)

type logv struct {
	kind string
	fold string
	file string
	line int
	char int
	args []interface{}
}

func input(args []otto.Value) (vars []interface{}) {
	for _, v := range args {
		switch {
		case v.IsFunction():
			vars = append(vars, v)
		case v.IsNull():
			vars = append(vars, nil)
		case v.IsUndefined():
			vars = append(vars, nil)
		case v.IsObject():
			obj, _ := v.Export()
			vars = append(vars, obj)
		default:
			vars = append(vars, v.String())
		}
	}
	return
}

func output(kind string, call otto.FunctionCall, args ...interface{}) {

	fold, file := path.Split(call.Otto.Context().Filename)
	line := call.Otto.Context().Line
	char := call.Otto.Context().Column

	line -= 1 // Account for injected module header

	data := logv{
		kind: kind,
		fold: fold,
		file: file,
		line: line,
		char: char,
		args: args,
	}

	log.Printf("console.%s: %v in %s%s at %d:%d", data.kind, data.args, data.fold, data.file, data.line, data.char)

}

func init() {

	orbit.OnInit(func(ctx *orbit.Orbit) {

		// TODO need to use fibre request context here

		timers := make(map[string]time.Time)

		console := map[string]interface{}{

			"log": func(call otto.FunctionCall) otto.Value {
				output("log", call, input(call.ArgumentList)...)
				return otto.UndefinedValue()
			},

			"info": func(call otto.FunctionCall) otto.Value {
				output("info", call, input(call.ArgumentList)...)
				return otto.UndefinedValue()
			},

			"warn": func(call otto.FunctionCall) otto.Value {
				output("warn", call, input(call.ArgumentList)...)
				return otto.UndefinedValue()
			},

			"error": func(call otto.FunctionCall) otto.Value {
				output("error", call, input(call.ArgumentList)...)
				return otto.UndefinedValue()
			},

			"debug": func(call otto.FunctionCall) otto.Value {
				output("debug", call, input(call.ArgumentList)...)
				return otto.UndefinedValue()
			},

			"trace": func(call otto.FunctionCall) otto.Value {
				output("trace", call, input(call.ArgumentList)...)
				return otto.UndefinedValue()
			},

			"time": func(call otto.FunctionCall) otto.Value {
				timers[call.Argument(0).String()] = time.Now()
				return otto.UndefinedValue()
			},

			"stop": func(call otto.FunctionCall) otto.Value {
				name := call.Argument(0).String()
				amnt := time.Since(timers[name])
				delete(timers, name)
				output("time", call, name, amnt)
				return otto.UndefinedValue()
			},

			"timeEnd": func(call otto.FunctionCall) otto.Value {
				name := call.Argument(0).String()
				amnt := time.Since(timers[name])
				delete(timers, name)
				output("time", call, name, amnt)
				return otto.UndefinedValue()
			},
		}

		ctx.Def("console", console)

	})

}
