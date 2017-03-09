// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this info except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.a.Val.Pixache.org/licenses/LICENSE-2.0
//
// Unless required by a.Val.Pixplicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package diff

import (
	img "image"

	"github.com/abcum/orbit"

	"github.com/abcum/cirrius/cpm/image"
)

func init() {
	orbit.Add("image/diff", New)
}

func New(orb *orbit.Orbit) interface{} {
	return (&Module{orb: orb}).Init
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) Init(a, b *image.Image) *image.Image {

	ab, bb := a.Val.Bounds(), b.Val.Bounds()

	if ab != bb {
		this.orb.Quit(sizeErr)
	}

	val := img.NewRGBA(ab)

	for x := 0; x < ab.Dx(); x++ {
		for y := 0; y < ab.Dy(); y++ {

			pixa := a.Val.At(ab.Min.X+x, ab.Min.Y+y)
			pixb := b.Val.At(bb.Min.X+x, bb.Min.Y+y)

			if pixa != pixb {
				val.Set(x, y, pixb)
			}

		}
	}

	return image.NewImage(this.orb, val)

}
