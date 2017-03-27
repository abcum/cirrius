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

// +build cgo,webkit

package browser

import (
	"image"

	"github.com/abcum/orbit"

	"github.com/abcum/webkit"

	img "github.com/abcum/cirrius/cpm/image"
)

type Page struct {
	orb *orbit.Orbit
	val *webkit.WebView
}

func NewPage(orb *orbit.Orbit, val *webkit.WebView) *Page {
	return &Page{
		orb: orb,
		val: val,
	}
}

func (this *Page) URL(uri string) string {
	return this.val.URI()
}

func (this *Page) Title() string {
	return this.val.Title()
}

func (this *Page) LoadUrl(url string) {
	this.val.LoadUri(url)
}

func (this *Page) LoadText(text string) {
	this.val.LoadText(text)
}

func (this *Page) LoadHtml(html, url string) {
	this.val.LoadHtml(html, url)
}

func (this *Page) Screenshot(cb func(*img.Image)) {
	this.val.GetSnapshot(func(res *image.RGBA, err error) {
		if err != nil {
			this.orb.Quit(err)
		}
		cb(img.NewImage(this.orb, res))
	})
}
