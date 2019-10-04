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
	"fmt"

	"image/color"

	"github.com/abcum/orbit"
	"github.com/lucasb-eyer/go-colorful"
)

type Colour struct {
	orb *orbit.Orbit
	typ kind
	col colorful.Color
	c   float64
	m   float64
	y   float64
	k   float64
}

func NewColour(orb *orbit.Orbit, typ kind, col colorful.Color) *Colour {
	return &Colour{
		orb: orb,
		col: col,
		typ: typ,
	}
}

func NewColourWithCMYK(orb *orbit.Orbit, typ kind, col colorful.Color, c, m, y, k float64) *Colour {
	return &Colour{
		orb: orb,
		col: col,
		typ: typ,
		c:   c,
		m:   m,
		y:   y,
		k:   k,
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

func (this *Colour) Hcl() [3]float64 {
	h, c, l := this.col.Hcl()
	return [3]float64{fixed(h), fixed(c * 100), fixed(l * 100)}
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
	switch this.typ {
	case cmyk:
		return [4]float64{this.c, this.m, this.y, this.k}
	default:
		r, g, b := this.col.RGB255()
		c, m, y, k := color.RGBToCMYK(r, g, b)
		return [4]float64{
			fixed(float64(c) / 2.55),
			fixed(float64(m) / 2.55),
			fixed(float64(y) / 2.55),
			fixed(float64(k) / 2.55),
		}
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

func (this *Colour) String() string {
	switch this.typ {
	default:
		ru, gu, bu := this.col.RGB255()
		rp, gp, bp := float64(ru), float64(gu), float64(bu)
		return fmt.Sprintf("rgb %v %v %v", fixed(rp/255), fixed(gp/255), fixed(bp/255))
	case cmyk:
		cu, mu, yu, ku := this.c, this.m, this.y, this.k
		cp, mp, yp, kp := float64(cu), float64(mu), float64(yu), float64(ku)
		return fmt.Sprintf("cmyk %v %v %v %v", fixed(cp/100), fixed(mp/100), fixed(yp/100), fixed(kp/100))
	}
}

func (this *Colour) Output() (string, float64, float64, float64, float64) {
	switch this.typ {
	default:
		ru, gu, bu := this.col.RGB255()
		rp, gp, bp := float64(ru), float64(gu), float64(bu)
		return "rgb", fixed(rp / 255), fixed(gp / 255), fixed(bp / 255), 0
	case cmyk:
		cu, mu, yu, ku := this.c, this.m, this.y, this.k
		cp, mp, yp, kp := float64(cu), float64(mu), float64(yu), float64(ku)
		return "cmyk", fixed(cp / 100), fixed(mp / 100), fixed(yp / 100), fixed(kp / 100)
	}
}
