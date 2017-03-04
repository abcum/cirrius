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

package uuid

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
	"github.com/ugorji/go/codec"
)

func init() {

	var h codec.CborHandle

	h.Canonical = true
	h.InternString = true
	h.CheckCircularRef = false
	h.AsSymbols = codec.AsSymbolDefault
	h.SliceType = reflect.TypeOf([]interface{}(nil))
	h.MapType = reflect.TypeOf(map[string]interface{}(nil))

	orbit.Add("cbor", func(ctx *orbit.Orbit) (otto.Value, error) {
		return ctx.ToValue(map[string]interface{}{
			"encode": func(call otto.FunctionCall) (val otto.Value) {

				if len(call.ArgumentList) > 1 {
					msg := fmt.Sprintf("Too many arguments to 'cbor.encode'. Expected 1 argument, but received %d.", len(call.ArgumentList))
					panic(ctx.MakeCustomError("Error", msg))
				}

				var err error
				var obj interface{}
				var buf bytes.Buffer

				if obj, err = call.Argument(0).Export(); err != nil {
					panic(ctx.MakeCustomError("Error", err.Error()))
				}

				if err = codec.NewEncoder(&buf, &h).Encode(obj); err != nil {
					panic(ctx.MakeCustomError("Error", err.Error()))
				}

				if val, err = otto.ToValue(buf.String()); err != nil {
					panic(ctx.MakeCustomError("Error", err.Error()))
				}

				return val

			},
			"decode": func(call otto.FunctionCall) (val otto.Value) {
				return otto.TrueValue()
			},
		})
	})

}
