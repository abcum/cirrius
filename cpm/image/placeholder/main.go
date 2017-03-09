// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package placeholder

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/abcum/orbit"
	"github.com/golang/freetype"

	"github.com/abcum/cirrius/cpm/colour"
	cpm "github.com/abcum/cirrius/cpm/image"
)

const (
	maxRatioX = 0.64
	maxRatioY = 0.23
)

func init() {
	orbit.Add("image/placeholder", New)
}

func New(orb *orbit.Orbit) interface{} {
	return (&Module{orb: orb}).Init
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) Init(content string, w, h int, cols ...*colour.Colour) *cpm.Image {

	// Specify default foreground
	// and background colours to be
	// used if none are defined.

	var fg = color.RGBA{R: 150, G: 150, B: 150, A: 255}
	var bg = color.RGBA{R: 215, G: 215, B: 215, A: 255}

	switch len(cols) {
	case 1:
		fg = cols[0].ToRGBA()
	case 2:
		fg = cols[0].ToRGBA()
		bg = cols[1].ToRGBA()
	}

	fore := image.NewUniform(fg)
	back := image.NewUniform(bg)
	test := image.NewRGBA(image.Rect(0, 0, 0, 0))

	// If no text has been specified
	// then set the default text as
	// the full image dimensions.

	if content == "" {
		content = fmt.Sprintf("%d x %d", w, h)
	}

	// Load the included arial font
	// so we can use it to output
	// the freetype text.

	font, err := freetype.ParseFont(arial)
	if err != nil {
		this.orb.Quit(err)
	}

	// Setup a new freetype context
	// so that we can write the text
	// and test the text dimensions.

	cont := freetype.NewContext()
	cont.SetDPI(72)
	cont.SetSrc(fore)
	cont.SetDst(test)
	cont.SetFont(font)
	cont.SetFontSize(1.0)
	cont.SetClip(test.Bounds())

	// Start the text a the left edge
	// but ensure that the text is at
	// the bottom of the text size.

	cent := freetype.Pt(0, int(cont.PointToFixed(1.0))>>6)

	// Draw the desired content text
	// onto the freetype context, and
	// panic if there are any errors.

	text, err := cont.DrawString(content, cent)
	if err != nil {
		this.orb.Quit(err)
	}

	// Check that the text stays within
	// the bounds of the image, using
	// the max text ratio values.

	size := math.Min(
		float64(cont.PointToFixed(float64(w)*maxRatioX))/float64(text.X),
		float64(cont.PointToFixed(float64(h)*maxRatioY))/float64(text.Y),
	)

	// Scale the font to the maximum
	// possible size whilst staying
	// within the image ratio bounds.

	cont.SetFontSize(size)

	// Reset the centre as there is no
	// need to measure from the bottom
	// of the text line.

	cent = freetype.Pt(0, 0)

	// Now draw the text onto the
	// freetype context using the max
	// possible font size.

	text, err = cont.DrawString(content, cent)
	if err != nil {
		this.orb.Quit(err)
	}

	// Calculate the centre alignment
	// position of the scaled text
	// for placing on the background.

	cent = freetype.Pt(
		(int(cont.PointToFixed(float64(w)/2.0))-int(text.X/2))>>6,
		(int(cont.PointToFixed(float64(h)/2.0))+int(size/2.6))>>6,
	)

	// Create a new image layer and
	// fill the layer with the chosen
	// background fill colour.

	img := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.Draw(img, img.Bounds(), back, image.ZP, draw.Src)

	// Specify the drawing context for
	// the freetype font, and ensure
	// it is clipped to the bounds.

	cont.SetDst(img)
	cont.SetClip(img.Bounds())

	// Draw the freetype font onto the
	// background image layer, and
	// convert it ot a JS image.

	if _, err := cont.DrawString(content, cent); err != nil {
		this.orb.Quit(err)
	}

	return cpm.NewImage(this.orb, img)

}
