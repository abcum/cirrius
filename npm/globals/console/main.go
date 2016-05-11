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
	"time"

	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
)

func init() {

	orbit.OnInit(func(ctx *orbit.Orbit) {

		// TODO need to use fibre request context here

		timers := make(map[string]time.Time)

		console := map[string]interface{}{}

		console["log"] = func(call otto.FunctionCall) otto.Value {
			log.Printf("console.log: %v on %s", call.ArgumentList, call.CallerLocation())
			return otto.UndefinedValue()
		}

		console["info"] = func(call otto.FunctionCall) otto.Value {
			log.Printf("console.info: %v", call.ArgumentList)
			return otto.UndefinedValue()
		}

		console["warn"] = func(call otto.FunctionCall) otto.Value {
			log.Printf("console.warn: %v", call.ArgumentList)
			return otto.UndefinedValue()
		}

		console["error"] = func(call otto.FunctionCall) otto.Value {
			log.Printf("console.error: %v", call.ArgumentList)
			return otto.UndefinedValue()
		}

		console["debug"] = func(call otto.FunctionCall) otto.Value {
			log.Printf("console.debug: %v", call.ArgumentList)
			return otto.UndefinedValue()
		}

		console["trace"] = func(call otto.FunctionCall) otto.Value {
			log.Printf("console.trace: %v", call.ArgumentList)
			return otto.UndefinedValue()
		}

		console["time"] = func(call otto.FunctionCall) otto.Value {
			timers[call.Argument(0).String()] = time.Now()
			return otto.UndefinedValue()
		}

		console["stop"] = func(call otto.FunctionCall) otto.Value {
			name := call.Argument(0).String()
			amnt := time.Since(timers[name])
			delete(timers, name)
			log.Printf("%s: %s", name, amnt)
			return otto.UndefinedValue()
		}

		console["timeEnd"] = func(call otto.FunctionCall) otto.Value {
			name := call.Argument(0).String()
			amnt := time.Since(timers[name])
			delete(timers, name)
			log.Printf("%s: %s", name, amnt)
			return otto.UndefinedValue()
		}

		ctx.Set("console", console)

	})

}
