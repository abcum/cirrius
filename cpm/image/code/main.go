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

package code

import (
	"github.com/abcum/orbit"

	"github.com/skip2/go-qrcode"

	"github.com/abcum/cirrius/cpm/colour"
	"github.com/abcum/cirrius/cpm/image"
)

func init() {
	orbit.Add("image/code", New)
}

func New(orb *orbit.Orbit) interface{} {
	return (&Module{orb: orb}).Init
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) Init(content string, size int, cols ...*colour.Colour) *image.Image {

	var qrc *qrcode.QRCode

	qrc, err := qrcode.New(content, qrcode.Highest)
	if err != nil {
		this.orb.Quit(err)
	}

	switch len(cols) {
	case 1:
		qrc.ForegroundColor = cols[0].ToRGBA()
	case 2:
		qrc.ForegroundColor = cols[0].ToRGBA()
		qrc.BackgroundColor = cols[1].ToRGBA()
	}

	return image.NewImage(this.orb, qrc.Image(size))

}
