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

package histogram

import (
	"github.com/abcum/orbit"
	"github.com/anthonynsimon/bild/histogram"

	"github.com/abcum/cirrius/cpm/image"
)

func init() {
	orbit.Add("image/histogram", New)
}

func New(orb *orbit.Orbit) interface{} {
	return (&Module{orb: orb}).Init
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) Init(img *image.Image) *image.Image {
	return image.NewImage(this.orb, histogram.NewRGBAHistogram(img.Val).Image())
}
