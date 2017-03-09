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

package binc

import (
	"reflect"

	"github.com/ugorji/go/codec"
)

var opt codec.BincHandle

func init() {
	opt.Canonical = true
	opt.InternString = true
	opt.CheckCircularRef = false
	opt.AsSymbols = codec.AsSymbolDefault
	opt.SliceType = reflect.TypeOf([]interface{}(nil))
	opt.MapType = reflect.TypeOf(map[string]interface{}(nil))
}
