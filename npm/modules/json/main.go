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

package json

import (
	"bytes"
	"reflect"
	"strings"

	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
	"github.com/ugorji/go/codec"
)

func init() {

	var h codec.JsonHandle

	h.Canonical = true
	h.InternString = true
	h.HTMLCharsAsIs = true
	h.CheckCircularRef = false
	h.AsSymbols = codec.AsSymbolDefault
	h.SliceType = reflect.TypeOf([]interface{}(nil))
	h.MapType = reflect.TypeOf(map[string]interface{}(nil))

	orbit.Add("json", func(orb *orbit.Orbit) (otto.Value, error) {
		return orb.ToValue(map[string]interface{}{
			"encode": func(obj interface{}) (val otto.Value) {
				var enc bytes.Buffer
				if err := codec.NewEncoder(&enc, &h).Encode(obj); err != nil {
					panic(orb.MakeCustomError("Error", err.Error()))
				}
				val, _ = otto.ToValue(enc.String())
				return val
			},
			"decode": func(enc string) (val otto.Value) {
				var out interface{}
				if err := codec.NewDecoder(strings.NewReader(enc), &h).Decode(&out); err != nil {
					panic(orb.MakeCustomError("Error", err.Error()))
				}
				val, _ = otto.ToValue(out)
				return val
			},
		})
	})

}
