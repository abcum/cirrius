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

package imaging

import (
	"image"
	"math"

	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"

	"github.com/disintegration/imaging"
)

func init() {

	orbit.Add("imaging", func(orb *orbit.Orbit) (otto.Value, error) {
		return orb.ToValue(map[string]interface{}{

			"blur": func(img image.Image, sigma float64) (val otto.Value) {
				val, _ = orb.ToValue(imaging.Blur(img, sigma))
				return
			},

			"contrast": func(img image.Image, percentage float64) (val otto.Value) {
				val, _ = orb.ToValue(imaging.AdjustContrast(img, percentage))
				return
			},

			"crop": func(img image.Image, x0, y0, x1, y1 int) (val otto.Value) {
				val, _ = orb.ToValue(imaging.Crop(img, image.Rect(x0, y0, x1, y1)))
				return
			},

			"darken": func(img image.Image, percentage float64) (val otto.Value) {
				val, _ = orb.ToValue(imaging.AdjustBrightness(img, -percentage))
				return
			},

			"fill": func(img image.Image, w, h int) (val otto.Value) {
				val, _ = orb.ToValue(imaging.Fill(img, w, h, imaging.TopLeft, imaging.Lanczos))
				return
			},

			"fit": func(img image.Image, w, h int) (val otto.Value) {
				val, _ = orb.ToValue(imaging.Fit(img, w, h, imaging.Lanczos))
				return
			},

			"flipH": func(img image.Image) (val otto.Value) {
				val, _ = orb.ToValue(imaging.FlipH(img))
				return
			},

			"flipV": func(img image.Image) (val otto.Value) {
				val, _ = orb.ToValue(imaging.FlipV(img))
				return
			},

			"gamma": func(img image.Image, gamma float64) (val otto.Value) {
				val, _ = orb.ToValue(imaging.AdjustGamma(img, math.Abs(gamma)))
				return
			},

			"grayscale": func(img image.Image) (val otto.Value) {
				val, _ = orb.ToValue(imaging.Grayscale(img))
				return
			},

			"invert": func(img image.Image) (val otto.Value) {
				val, _ = orb.ToValue(imaging.Invert(img))
				return
			},

			"lighten": func(img image.Image, percentage float64) (val otto.Value) {
				val, _ = orb.ToValue(imaging.AdjustBrightness(img, math.Abs(percentage)))
				return
			},

			"resize": func(img image.Image, w, h int) (val otto.Value) {
				val, _ = orb.ToValue(imaging.Resize(img, w, h, imaging.Lanczos))
				return
			},

			"rotate180": func(img image.Image) (val otto.Value) {
				val, _ = orb.ToValue(imaging.Rotate180(img))
				return
			},

			"rotate270": func(img image.Image) (val otto.Value) {
				val, _ = orb.ToValue(imaging.Rotate270(img))
				return
			},

			"rotate90": func(img image.Image) (val otto.Value) {
				val, _ = orb.ToValue(imaging.Rotate90(img))
				return
			},

			"sharpen": func(img image.Image, sigma float64) (val otto.Value) {
				val, _ = orb.ToValue(imaging.Sharpen(img, sigma))
				return
			},

			"sigmoid": func(img image.Image, midpoint, factor float64) (val otto.Value) {
				val, _ = orb.ToValue(imaging.AdjustSigmoid(img, midpoint, factor))
				return
			},

			"thumbnail": func(img image.Image, w, h int) (val otto.Value) {
				val, _ = orb.ToValue(imaging.Thumbnail(img, w, h, imaging.Lanczos))
				return
			},

			"transpose": func(img image.Image) (val otto.Value) {
				val, _ = orb.ToValue(imaging.Transpose(img))
				return
			},

			"transverse": func(img image.Image) (val otto.Value) {
				val, _ = orb.ToValue(imaging.Transverse(img))
				return
			},
		})
	})

}
