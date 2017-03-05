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

package twooffive

import (
	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/twooffive"
)

func init() {

	orbit.Add("twooffive", func(orb *orbit.Orbit) (otto.Value, error) {
		return orb.ToValue(func(content string, w, h int) (val otto.Value) {

			var err error
			var img barcode.Barcode

			if img, err = twooffive.Encode(content, true); err != nil {
				panic(err)
			}

			if img, err = barcode.Scale(img, w, h); err != nil {
				panic(err)
			}

			val, _ = orb.ToValue(img)

			return

		})
	})

}
