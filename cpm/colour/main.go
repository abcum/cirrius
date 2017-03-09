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
	"github.com/abcum/orbit"
	"github.com/lucasb-eyer/go-colorful"
)

func init() {
	orbit.Add("color", New)
	orbit.Add("colour", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
	}
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) Hex(h string) *Colour {
	col, err := colorful.Hex(h)
	if err != nil {
		this.orb.Quit(err)
	}
	return NewColour(this.orb, col)
}

func (this *Module) Rgb(r, g, b float64) *Colour {
	r = r / 255
	g = g / 255
	b = b / 255
	return NewColour(this.orb, colorful.Color{R: r, G: g, B: b})
}

func (this *Module) Hcl(h, c, l float64) *Colour {
	c = c / 100
	l = l / 100
	return NewColour(this.orb, colorful.Hcl(h, c, l))
}

func (this *Module) Hsl(h, s, l float64) *Colour {
	s = s / 100
	l = l / 100
	return NewColour(this.orb, colorful.Hsl(h, s, l))
}

func (this *Module) Hsv(h, s, v float64) *Colour {
	s = s / 100
	v = v / 100
	return NewColour(this.orb, colorful.Hsv(h, s, v))
}

func (this *Module) Lab(l, a, b float64) *Colour {
	l = l / 100
	a = a / 100
	a = a / 100
	return NewColour(this.orb, colorful.Lab(l, a, b))
}

func (this *Module) Luv(l, u, v float64) *Colour {
	l = l / 100
	u = u / 100
	v = v / 100
	return NewColour(this.orb, colorful.Luv(l, u, v))
}

func (this *Module) Xyy(x, y, z float64) *Colour {
	x = x / 100
	y = y / 100
	z = z / 100
	return NewColour(this.orb, colorful.Xyy(x, y, z))
}

func (this *Module) Xyz(x, y, z float64) *Colour {
	x = x / 100
	y = y / 100
	z = z / 100
	return NewColour(this.orb, colorful.Xyz(x, y, z))
}

func (this *Module) Cmyk(c, m, y, k float64) *Colour {
	r := float64((1 - c/100) * (1 - k/100))
	g := float64((1 - m/100) * (1 - k/100))
	b := float64((1 - y/100) * (1 - k/100))
	return NewColour(this.orb, colorful.Color{R: r, G: g, B: b})
}
