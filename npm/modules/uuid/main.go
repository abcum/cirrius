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
	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
	"github.com/satori/go.uuid"
)

func init() {

	orbit.Add("uuid", func(ctx *orbit.Orbit) (otto.Value, error) {
		return ctx.ToValue(map[string]interface{}{
			"namespace": map[string]interface{}{
				"dns":  uuid.NamespaceDNS,
				"url":  uuid.NamespaceURL,
				"oid":  uuid.NamespaceOID,
				"x500": uuid.NamespaceX500,
			},
			"v1": func(call otto.FunctionCall) otto.Value {
				value, _ := ctx.ToValue(uuid.NewV1().String())
				return value
			},
			"v2": func(call otto.FunctionCall) otto.Value {
				arg1 := byte(call.Argument(0).String()[0])
				value, _ := ctx.ToValue(uuid.NewV2(arg1).String())
				return value
			},
			"v3": func(call otto.FunctionCall) otto.Value {
				arg1, _ := uuid.FromString(call.Argument(0).String())
				arg2 := call.Argument(1).String()
				value, _ := ctx.ToValue(uuid.NewV3(arg1, arg2).String())
				return value
			},
			"v4": func(call otto.FunctionCall) otto.Value {
				value, _ := ctx.ToValue(uuid.NewV4().String())
				return value
			},
			"v5": func(call otto.FunctionCall) otto.Value {
				arg1, _ := uuid.FromString(call.Argument(0).String())
				arg2 := call.Argument(1).String()
				value, _ := ctx.ToValue(uuid.NewV5(arg1, arg2).String())
				return value
			},
		})
	})

}
