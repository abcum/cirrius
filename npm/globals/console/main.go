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
	"strconv"
	"strings"
	"time"

	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
)

type logv struct {
	kind string
	fold string
	file string
	line int64
	char int64
	args []otto.Value
	vars []interface{}
}

func save(kind string, call otto.FunctionCall, vars ...interface{}) {

	fold, _ := call.Otto.Get("__dirname")
	fost, _ := fold.ToString()

	file, _ := call.Otto.Get("__filename")
	fist, _ := file.ToString()

	loca := strings.Split(call.CallerLocation()[12:], ":")
	line, _ := strconv.ParseInt(loca[0], 10, 0)
	char, _ := strconv.ParseInt(loca[1], 10, 0)

	line -= 1 // Account for injected module header

	data := logv{
		kind: kind,
		fold: fost,
		file: fist,
		line: line,
		char: char,
		vars: vars,
		args: call.ArgumentList,
	}

	log.Printf("console.%s: %v in %s%s at %d:%d with vars %v", data.kind, data.args, data.fold, data.file, data.line, data.char, data.vars)

}

func init() {

	orbit.OnInit(func(ctx *orbit.Orbit) {

		// TODO need to use fibre request context here

		timers := make(map[string]time.Time)

		console := map[string]interface{}{

			"log": func(call otto.FunctionCall) otto.Value {
				save("log", call)
				return otto.UndefinedValue()
			},

			"info": func(call otto.FunctionCall) otto.Value {
				save("info", call)
				return otto.UndefinedValue()
			},

			"warn": func(call otto.FunctionCall) otto.Value {
				save("warn", call)
				return otto.UndefinedValue()
			},

			"error": func(call otto.FunctionCall) otto.Value {
				save("error", call)
				return otto.UndefinedValue()
			},

			"debug": func(call otto.FunctionCall) otto.Value {
				save("debug", call)
				return otto.UndefinedValue()
			},

			"trace": func(call otto.FunctionCall) otto.Value {
				save("trace", call)
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
				save("time", call, name, amnt)
				return otto.UndefinedValue()
			},

			"timeEnd": func(call otto.FunctionCall) otto.Value {
				name := call.Argument(0).String()
				amnt := time.Since(timers[name])
				delete(timers, name)
				save("time", call, name, amnt)
				return otto.UndefinedValue()
			},
		}

		ctx.Def("console", console)

	})

}
