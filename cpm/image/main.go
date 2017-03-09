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

package image

import (
	"io"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/abcum/orbit"

	_ "github.com/donatj/mpo"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"

	"github.com/disintegration/imaging"

	"github.com/anthonynsimon/bild/channel"
)

func init() {
	orbit.Add("image", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
		// Image channels
		R: channel.Red,
		G: channel.Green,
		B: channel.Blue,
		A: channel.Alpha,
		// Image formats
		BMP:  imaging.BMP,
		GIF:  imaging.GIF,
		PNG:  imaging.PNG,
		JPG:  imaging.JPEG,
		JPEG: imaging.JPEG,
		TIFF: imaging.TIFF,
		// Image anchors
		CENTER:      imaging.Center,
		TOP:         imaging.Top,
		LEFT:        imaging.Left,
		RIGHT:       imaging.Right,
		BOTTOM:      imaging.Bottom,
		TOPLEFT:     imaging.TopLeft,
		TOPRIGHT:    imaging.TopRight,
		BOTTOMLEFT:  imaging.BottomLeft,
		BOTTOMRIGHT: imaging.BottomRight,
		// Image resamplers
		NEAREST:    imaging.NearestNeighbor,
		BOX:        imaging.Box,
		LINEAR:     imaging.Linear,
		HERMITE:    imaging.Hermite,
		MNETRAVALI: imaging.MitchellNetravali,
		CATMULLROM: imaging.CatmullRom,
		BSPLINE:    imaging.BSpline,
		GAUSSIAN:   imaging.Gaussian,
		LANCZOS:    imaging.Lanczos,
		HANN:       imaging.Hann,
		HAMMING:    imaging.Hamming,
		BLACKMAN:   imaging.Blackman,
		BARTLETT:   imaging.Bartlett,
		WELCH:      imaging.Welch,
		COSINE:     imaging.Cosine,
	}
}

type Module struct {
	orb *orbit.Orbit
	// Image channels
	R channel.Channel
	G channel.Channel
	B channel.Channel
	A channel.Channel
	// Image formats
	BMP  imaging.Format
	GIF  imaging.Format
	PNG  imaging.Format
	JPG  imaging.Format
	JPEG imaging.Format
	TIFF imaging.Format
	// Image anchors
	CENTER      imaging.Anchor
	TOP         imaging.Anchor
	LEFT        imaging.Anchor
	RIGHT       imaging.Anchor
	BOTTOM      imaging.Anchor
	TOPLEFT     imaging.Anchor
	TOPRIGHT    imaging.Anchor
	BOTTOMLEFT  imaging.Anchor
	BOTTOMRIGHT imaging.Anchor
	// Image resamplers
	NEAREST    imaging.ResampleFilter
	BOX        imaging.ResampleFilter
	LINEAR     imaging.ResampleFilter
	HERMITE    imaging.ResampleFilter
	MNETRAVALI imaging.ResampleFilter
	CATMULLROM imaging.ResampleFilter
	BSPLINE    imaging.ResampleFilter
	GAUSSIAN   imaging.ResampleFilter
	LANCZOS    imaging.ResampleFilter
	HANN       imaging.ResampleFilter
	HAMMING    imaging.ResampleFilter
	BLACKMAN   imaging.ResampleFilter
	BARTLETT   imaging.ResampleFilter
	WELCH      imaging.ResampleFilter
	COSINE     imaging.ResampleFilter
}

func (this *Module) Load(r io.Reader) *Image {
	img, err := imaging.Decode(r)
	if err != nil {
		this.orb.Quit(err)
	}
	return NewImage(this.orb, img)
}
