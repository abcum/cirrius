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

package browser

import (
	"github.com/abcum/orbit"
	"github.com/abcum/webdriver"

	"github.com/disintegration/imaging"

	img "github.com/abcum/cirrius/cpm/image"
)

type Page struct {
	orb *orbit.Orbit
	val *webdriver.Session
}

func NewPage(orb *orbit.Orbit, val *webdriver.Session) *Page {
	return (&Page{
		orb: orb,
		val: val,
	}).init()
}

func (this *Page) init() *Page {
	this.orb.Once("exit", func() {
		this.val.Delete()
	})
	return this
}

func (this *Page) Close() {
	this.val.Delete()
}

func (this *Page) URL() string {
	out, err := this.val.Url()
	if err != nil {
		this.orb.Quit(err)
	}
	return out
}

func (this *Page) Title() string {
	out, err := this.val.Title()
	if err != nil {
		this.orb.Quit(err)
	}
	return out
}

func (this *Page) Source() string {
	out, err := this.val.Source()
	if err != nil {
		this.orb.Quit(err)
	}
	return out
}

func (this *Page) Wait(ms int) {
	if err := this.val.Timeouts("implicit", ms); err != nil {
		this.orb.Quit(err)
	}
}

func (this *Page) Load(url string) {
	if err := this.val.Load(url); err != nil {
		this.orb.Quit(err)
	}
}

func (this *Page) Back() {
	if err := this.val.Back(); err != nil {
		this.orb.Quit(err)
	}
}

func (this *Page) Forward() {
	if err := this.val.Forward(); err != nil {
		this.orb.Quit(err)
	}
}

func (this *Page) Refresh() {
	if err := this.val.Refresh(); err != nil {
		this.orb.Quit(err)
	}
}

func (this *Page) Resize(w, h int) {
	if err := this.val.Window().Resize(w, h); err != nil {
		this.orb.Quit(err)
	}
}

func (this *Page) Active() *Node {
	nde, err := this.val.Active()
	if err != nil {
		this.orb.Quit(err)
	}
	return NewNode(this.orb, this.val, nde)
}

func (this *Page) Find(search string) *Node {
	nde, err := this.val.Element(webdriver.FindByCss, search)
	if err != nil {
		this.orb.Quit(err)
	}
	return NewNode(this.orb, this.val, nde)
}

func (this *Page) Search(search string) []*Node {
	nds, err := this.val.Elements(webdriver.FindByCss, search)
	if err != nil {
		this.orb.Quit(err)
	}
	var out []*Node
	for _, nde := range nds {
		out = append(out, NewNode(this.orb, this.val, nde))
	}
	return out
}

func (this *Page) Screenshot() *img.Image {
	rdr, err := this.val.Screenshot()
	if err != nil {
		this.orb.Quit(err)
	}
	out, err := imaging.Decode(rdr)
	if err != nil {
		this.orb.Quit(err)
	}
	return img.NewImage(this.orb, out)
}
