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

package jpeg

import (
	"io"

	"image"

	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"

	"github.com/disintegration/imaging"
)

func init() {

	orbit.Add("jpeg", func(orb *orbit.Orbit) (otto.Value, error) {
		return orb.ToValue(map[string]interface{}{
			"encode": func(w io.Writer, i image.Image) (val otto.Value) {
				if err := imaging.Encode(w, i, imaging.JPEG); err != nil {
					panic(orb.MakeCustomError("Error", err.Error()))
				}
				return
			},

			"decode": func(r io.Reader) (val otto.Value) {
				img, err := imaging.Decode(r)
				if err != nil {
					panic(orb.MakeCustomError("Error", err.Error()))
				}
				val, _ = orb.ToValue(img)
				return
			},
		})
	})

}
