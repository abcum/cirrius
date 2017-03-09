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
	"math"

	"image/color"

	"github.com/abcum/orbit"
	"github.com/lucasb-eyer/go-colorful"
)

type Colour struct {
	orb *orbit.Orbit
	col colorful.Color
}

func NewColour(orb *orbit.Orbit, col colorful.Color) *Colour {
	return &Colour{
		orb: orb,
		col: col,
	}
}

func (this *Colour) Hex() string {
	return this.col.Hex()
}

func (this *Colour) Rgb() [3]uint8 {
	r, g, b := this.col.RGB255()
	return [3]uint8{r, g, b}
}

func (this *Colour) Rgba() [4]uint8 {
	r, g, b := this.col.RGB255()
	return [4]uint8{r, g, b, 255}
}

func (this *Colour) Hsl() [3]float64 {
	h, s, l := this.col.Hsl()
	return [3]float64{fixed(h), fixed(s * 100), fixed(l * 100)}
}

func (this *Colour) Hsv() [3]float64 {
	h, s, v := this.col.Hsv()
	return [3]float64{fixed(h), fixed(s * 100), fixed(v * 100)}
}

func (this *Colour) Lab() [3]float64 {
	l, a, b := this.col.Lab()
	return [3]float64{fixed(l * 100), fixed(a * 100), fixed(b * 100)}
}

func (this *Colour) Luv() [3]float64 {
	l, u, v := this.col.Luv()
	return [3]float64{fixed(l * 100), fixed(u * 100), fixed(v * 100)}
}

func (this *Colour) Xyy() [3]float64 {
	x, y, z := this.col.Xyy()
	return [3]float64{fixed(x * 100), fixed(y * 100), fixed(z * 100)}
}

func (this *Colour) Xyz() [3]float64 {
	x, y, z := this.col.Xyz()
	return [3]float64{fixed(x * 100), fixed(y * 100), fixed(z * 100)}
}

func (this *Colour) Cmyk() [4]float64 {
	r, g, b := this.col.RGB255()
	c, m, y, k := color.RGBToCMYK(r, g, b)
	return [4]float64{
		fixed(float64(c) / 2.55),
		fixed(float64(m) / 2.55),
		fixed(float64(y) / 2.55),
		fixed(float64(k) / 2.55),
	}
}

func (this *Colour) ToRGBA() color.RGBA {
	r, g, b := this.col.RGB255()
	return color.RGBA{R: r, G: g, B: b, A: 255}
}

func (this *Colour) ToCMYK() color.CMYK {
	r, g, b := this.col.RGB255()
	c, m, y, k := color.RGBToCMYK(r, g, b)
	return color.CMYK{C: c, M: m, Y: y, K: k}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func fixed(num float64) float64 {
	output := math.Pow(10, float64(2))
	return float64(round(num*output)) / output
}
