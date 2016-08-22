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

package colour

import (
	"image/color"

	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
)

func init() {

	orbit.Add("colour", func(ctx *orbit.Orbit) (otto.Value, error) {
		return ctx.ToValue(map[string]interface{}{
			"rgb2cmyk": func(r, g, b uint8) []int {
				c, m, y, k := color.RGBToCMYK(r, g, b)
				return []int{int(c), int(m), int(y), int(k)}
			},
			"cmyk2rgb": func(c, m, y, k uint8) []int {
				r, g, b := color.CMYKToRGB(c, m, y, k)
				return []int{int(r), int(g), int(b)}
			},
		})
	})

}
