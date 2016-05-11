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
	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
)

func init() {

	orbit.OnInit(func(ctx *orbit.Orbit) {

		output := map[string]interface{}{}

		output["text"] = func(call otto.FunctionCall) otto.Value {
			// ctx.fibre.Text(call.Argument(0).Export())
			return otto.UndefinedValue()
		}

		output["html"] = func(call otto.FunctionCall) otto.Value {
			// ctx.fibre.HTML(call.Argument(0).Export())
			return otto.UndefinedValue()
		}

		output["json"] = func(call otto.FunctionCall) otto.Value {
			// ctx.fibre.JSON(call.Argument(0).Export())
			return otto.UndefinedValue()
		}

		output["pack"] = func(call otto.FunctionCall) otto.Value {
			// ctx.fibre.PACK(call.Argument(0).Export())
			return otto.UndefinedValue()
		}

		ctx.Set("output", output)

	})

}
