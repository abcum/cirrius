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
	"math"

	"github.com/abcum/orbit"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/robertkrimen/otto"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func fixed(num float64) float64 {
	output := math.Pow(10, float64(2))
	return float64(round(num*output)) / output
}

func eval(col colorful.Color) map[string]interface{} {

	return map[string]interface{}{
		"hex": func() string {
			return col.Hex()
		}(),
		"hsl": func() [3]float64 {
			h, s, l := col.Hsl()
			return [3]float64{fixed(h), fixed(s * 100), fixed(l * 100)}
		}(),
		"hsv": func() [3]float64 {
			h, s, v := col.Hsv()
			return [3]float64{fixed(h), fixed(s * 100), fixed(v * 100)}
		}(),
		"lab": func() [3]float64 {
			l, a, b := col.Lab()
			return [3]float64{fixed(l * 100), fixed(a * 100), fixed(b * 100)}
		}(),
		"luv": func() [3]float64 {
			l, u, v := col.Luv()
			return [3]float64{fixed(l * 100), fixed(u * 100), fixed(v * 100)}
		}(),
		"xyy": func() [3]float64 {
			x, y, z := col.Xyy()
			return [3]float64{fixed(x * 100), fixed(y * 100), fixed(z * 100)}
		}(),
		"xyz": func() [3]float64 {
			x, y, z := col.Xyz()
			return [3]float64{fixed(x * 100), fixed(y * 100), fixed(z * 100)}
		}(),
		"rgb": func() [3]uint8 {
			r, g, b := col.RGB255()
			return [3]uint8{r, g, b}
		}(),
		"cmyk": func() [4]float64 {
			r, g, b := col.RGB255()
			c, m, y, k := color.RGBToCMYK(r, g, b)
			return [4]float64{
				fixed(float64(c) / 2.55),
				fixed(float64(m) / 2.55),
				fixed(float64(y) / 2.55),
				fixed(float64(k) / 2.55),
			}
		}(),
	}

}

func init() {

	orbit.Add("colour", func(ctx *orbit.Orbit) (otto.Value, error) {
		return ctx.ToValue(map[string]interface{}{
			"rgb": func(r, g, b float64) map[string]interface{} {
				r = r / 255
				g = g / 255
				b = b / 255
				return eval(colorful.Color{R: r, G: g, B: b})
			},
			"hcl": func(h, c, l float64) map[string]interface{} {
				c = c / 100
				l = l / 100
				return eval(colorful.Hcl(h, c, l))
			},
			"hsl": func(h, s, l float64) map[string]interface{} {
				s = s / 100
				l = l / 100
				return eval(colorful.Hsl(h, s, l))
			},
			"hsv": func(h, s, v float64) map[string]interface{} {
				s = s / 100
				v = v / 100
				return eval(colorful.Hsl(h, s, v))
			},
			"lab": func(l, a, b float64) map[string]interface{} {
				l = l / 100
				a = a / 100
				b = b / 100
				return eval(colorful.Lab(l, a, b))
			},
			"luv": func(l, u, v float64) map[string]interface{} {
				l = l / 100
				u = u / 100
				v = v / 100
				return eval(colorful.Luv(l, u, v))
			},
			"xyy": func(x, y, z float64) map[string]interface{} {
				x = x / 100
				y = y / 100
				z = z / 100
				return eval(colorful.Xyy(x, y, z))
			},
			"xyz": func(x, y, z float64) map[string]interface{} {
				x = x / 100
				y = y / 100
				z = z / 100
				return eval(colorful.Xyz(x, y, z))
			},
			"cmyk": func(c, m, y, k float64) map[string]interface{} {
				r := (1 - c/100) * (1 - k/100)
				g := (1 - m/100) * (1 - k/100)
				b := (1 - y/100) * (1 - k/100)
				return eval(colorful.Color{R: float64(r), G: float64(g), B: float64(b)})
			},
		})
	})

}
