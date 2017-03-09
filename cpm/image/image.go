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
	"bytes"
	"image"
	"io"
	"math"

	"encoding/base64"

	"github.com/abcum/orbit"

	"github.com/abcum/cirrius/cpm/stream"

	"github.com/disintegration/imaging"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/blur"
	"github.com/anthonynsimon/bild/channel"
	"github.com/anthonynsimon/bild/clone"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/segment"
	"github.com/anthonynsimon/bild/transform"
)

type Image struct {
	orb *orbit.Orbit
	Val image.Image `console:"-"`
}

func NewImage(orb *orbit.Orbit, val image.Image) *Image {
	return &Image{
		orb: orb,
		Val: val,
	}
}

func (this *Image) Width() int {
	return this.Val.Bounds().Max.X
}

func (this *Image) Height() int {
	return this.Val.Bounds().Max.Y
}

func (this *Image) Blur(radius float64) *Image {
	img := blur.Gaussian(this.Val, radius)
	return NewImage(this.orb, img)
}

func (this *Image) Brightness(percentage float64) *Image {
	img := adjust.Brightness(this.Val, percentage/100)
	return NewImage(this.orb, img)
}

func (this *Image) Channel(c int) *Image {
	img := channel.Extract(this.Val, channel.Channel(c))
	return NewImage(this.orb, img)
}

func (this *Image) Contrast(percentage float64) *Image {
	img := adjust.Contrast(this.Val, percentage/100)
	return NewImage(this.orb, img)
}

func (this *Image) Crop(x0, y0, x1, y1 int) *Image {
	img := imaging.Crop(this.Val, image.Rect(x0, y0, x1, y1))
	return NewImage(this.orb, img)
}

func (this *Image) Darken(percentage float64) *Image {
	img := adjust.Brightness(this.Val, -percentage/100)
	return NewImage(this.orb, img)
}

func (this *Image) Dilate(radius float64) *Image {
	img := effect.Dilate(this.Val, radius)
	return NewImage(this.orb, img)
}

func (this *Image) Edges(radius float64) *Image {
	img := effect.EdgeDetection(this.Val, radius)
	return NewImage(this.orb, img)
}

func (this *Image) Emboss() *Image {
	img := effect.Emboss(this.Val)
	return NewImage(this.orb, img)
}

func (this *Image) Erode(radius float64) *Image {
	img := effect.Erode(this.Val, radius)
	return NewImage(this.orb, img)
}

func (this *Image) Fill(w, h int, anchor int, resampler ...imaging.ResampleFilter) *Image {
	res := imaging.Lanczos
	if len(resampler) > 0 {
		res = resampler[0]
	}
	img := imaging.Fill(this.Val, w, h, imaging.Anchor(anchor), res)
	return NewImage(this.orb, img)
}

func (this *Image) Fit(w, h int, resampler ...imaging.ResampleFilter) *Image {
	res := imaging.Lanczos
	if len(resampler) > 0 {
		res = resampler[0]
	}
	img := imaging.Fit(this.Val, w, h, res)
	return NewImage(this.orb, img)
}

func (this *Image) FlipH() *Image {
	img := imaging.FlipH(this.Val)
	return NewImage(this.orb, img)
}

func (this *Image) FlipV() *Image {
	img := imaging.FlipV(this.Val)
	return NewImage(this.orb, img)
}

func (this *Image) Gamma(gamma float64) *Image {
	img := adjust.Gamma(this.Val, math.Abs(gamma))
	return NewImage(this.orb, img)
}

func (this *Image) Grayscale() *Image {
	img := effect.Grayscale(this.Val)
	return NewImage(this.orb, img)
}

func (this *Image) Hue(change int) *Image {
	img := adjust.Hue(this.Val, change)
	return NewImage(this.orb, img)
}

func (this *Image) Invert() *Image {
	img := effect.Invert(this.Val)
	return NewImage(this.orb, img)
}

func (this *Image) Lighten(percentage float64) *Image {
	img := adjust.Brightness(this.Val, math.Abs(percentage)/100)
	return NewImage(this.orb, img)
}

func (this *Image) Median(radius float64) *Image {
	img := effect.Median(this.Val, radius)
	return NewImage(this.orb, img)
}

func (this *Image) Pad(x, y int) *Image {
	img := clone.Pad(this.Val, x, y, clone.NoFill)
	return NewImage(this.orb, img)
}

func (this *Image) Resize(w, h int, resampler ...imaging.ResampleFilter) *Image {
	res := imaging.Lanczos
	if len(resampler) > 0 {
		res = resampler[0]
	}
	img := imaging.Resize(this.Val, w, h, res)
	return NewImage(this.orb, img)
}

func (this *Image) Rotate(angle float64, x, y int) *Image {
	opt := &transform.RotationOptions{ResizeBounds: false, Pivot: &image.Point{X: x, Y: y}}
	img := transform.Rotate(this.Val, angle, opt)
	return NewImage(this.orb, img)
}

func (this *Image) Rotate180() *Image {
	img := imaging.Rotate180(this.Val)
	return NewImage(this.orb, img)
}

func (this *Image) Rotate270() *Image {
	img := imaging.Rotate270(this.Val)
	return NewImage(this.orb, img)
}

func (this *Image) Rotate90() *Image {
	img := imaging.Rotate90(this.Val)
	return NewImage(this.orb, img)
}

func (this *Image) Sepia() *Image {
	img := effect.Sepia(this.Val)
	return NewImage(this.orb, img)
}

func (this *Image) Saturation(percentage float64) *Image {
	img := adjust.Saturation(this.Val, percentage/100)
	return NewImage(this.orb, img)
}

func (this *Image) Sharpen(radius float64) *Image {
	img := imaging.Sharpen(this.Val, radius)
	return NewImage(this.orb, img)
}

func (this *Image) ShearH(angle float64) *Image {
	img := transform.ShearH(this.Val, angle)
	return NewImage(this.orb, img)
}

func (this *Image) ShearV(angle float64) *Image {
	img := transform.ShearV(this.Val, angle)
	return NewImage(this.orb, img)
}

func (this *Image) Sigmoid(midpoint, factor float64) *Image {
	img := imaging.AdjustSigmoid(this.Val, midpoint, factor)
	return NewImage(this.orb, img)
}

func (this *Image) Sobel() *Image {
	img := effect.Sobel(this.Val)
	return NewImage(this.orb, img)
}

func (this *Image) Threshold(level uint8) *Image {
	img := segment.Threshold(this.Val, level)
	return NewImage(this.orb, img)
}

func (this *Image) Thumbnail(w, h int, resampler ...imaging.ResampleFilter) *Image {
	res := imaging.Lanczos
	if len(resampler) > 0 {
		res = resampler[0]
	}
	img := imaging.Thumbnail(this.Val, w, h, res)
	return NewImage(this.orb, img)
}

func (this *Image) Translate(x, y int) *Image {
	img := transform.Translate(this.Val, x, y)
	return NewImage(this.orb, img)
}

func (this *Image) Transpose() *Image {
	img := imaging.Transpose(this.Val)
	return NewImage(this.orb, img)
}

func (this *Image) Transverse() *Image {
	img := imaging.Transverse(this.Val)
	return NewImage(this.orb, img)
}

func (this *Image) Trim(w, h int, anchor int) *Image {
	img := imaging.CropAnchor(this.Val, w, h, imaging.Anchor(anchor))
	return NewImage(this.orb, img)
}

func (this *Image) Unsharpen(radius, amount float64) *Image {
	img := effect.UnsharpMask(this.Val, radius, amount)
	return NewImage(this.orb, img)
}

func (this *Image) Pipe(w io.Writer, format ...int) {
	fmt := imaging.PNG
	if len(format) > 0 {
		fmt = imaging.Format(format[0])
	}
	if err := imaging.Encode(w, this.Val, imaging.Format(fmt)); err != nil {
		this.orb.Quit(err)
	}
	return
}

func (this *Image) Reader() *stream.Reader {
	b := bytes.NewBuffer(nil)
	if err := imaging.Encode(b, this.Val, imaging.PNG); err != nil {
		this.orb.Quit(err)
	}
	return stream.NewReader(this.orb, b)
}

func (this *Image) String(format ...int) string {
	f := imaging.PNG
	if len(format) > 0 {
		f = imaging.Format(format[0])
	}
	b := bytes.NewBuffer(nil)
	w := base64.NewEncoder(base64.StdEncoding, b)
	if err := imaging.Encode(w, this.Val, imaging.Format(f)); err != nil {
		this.orb.Quit(err)
	}
	return b.String()
}

func (this *Image) MarshalText() ([]byte, error) {
	b := bytes.NewBuffer(nil)
	w := base64.NewEncoder(base64.StdEncoding, b)
	if err := imaging.Encode(w, this.Val, imaging.PNG); err != nil {
		this.orb.Quit(err)
	}
	return b.Bytes(), nil
}

func (this *Image) MarshalBinary() ([]byte, error) {
	b := bytes.NewBuffer(nil)
	if err := imaging.Encode(b, this.Val, imaging.PNG); err != nil {
		this.orb.Quit(err)
	}
	return b.Bytes(), nil
}
