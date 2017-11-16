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
)

type Node struct {
	orb *orbit.Orbit
	wds *webdriver.Session
	val *webdriver.Element
}

func NewNode(orb *orbit.Orbit, wds *webdriver.Session, val *webdriver.Element) *Node {
	return &Node{
		orb: orb,
		wds: wds,
		val: val,
	}
}

func (this *Node) Size() interface{} {
	val, err := this.val.Size()
	if err != nil {
		this.orb.Quit(err)
	}
	return val
}

func (this *Node) Name() string {
	val, err := this.val.Name()
	if err != nil {
		this.orb.Quit(err)
	}
	return val
}

func (this *Node) Text() string {
	val, err := this.val.Text()
	if err != nil {
		this.orb.Quit(err)
	}
	return val
}

func (this *Node) Html() string {
	val, err := this.val.Html()
	if err != nil {
		this.orb.Quit(err)
	}
	return val
}

func (this *Node) Attr(name string) string {
	val, err := this.val.Attr(name)
	if err != nil {
		this.orb.Quit(err)
	}
	return val
}

func (this *Node) Css(name string) string {
	val, err := this.val.Css(name)
	if err != nil {
		this.orb.Quit(err)
	}
	return val
}

func (this *Node) Coords() interface{} {
	val, err := this.val.Location()
	if err != nil {
		this.orb.Quit(err)
	}
	return val
}

func (this *Node) Clear() *Node {
	if err := this.val.Clear(); err != nil {
		this.orb.Quit(err)
	}
	return this
}

func (this *Node) Click() *Node {
	if err := this.val.Click(); err != nil {
		this.orb.Quit(err)
	}
	return this
}

func (this *Node) Submit() *Node {
	if err := this.val.Submit(); err != nil {
		this.orb.Quit(err)
	}
	return this
}

func (this *Node) Value(keys string) *Node {
	if err := this.val.Value(keys); err != nil {
		this.orb.Quit(err)
	}
	return this
}

func (this *Node) Equals(n *Node) bool {
	val, err := this.val.Equals(n.val)
	if err != nil {
		this.orb.Quit(err)
	}
	return val
}

func (this *Node) Enabled() bool {
	val, err := this.val.Enabled()
	if err != nil {
		this.orb.Quit(err)
	}
	return val
}

func (this *Node) Displayed() bool {
	val, err := this.val.Displayed()
	if err != nil {
		this.orb.Quit(err)
	}
	return val
}

func (this *Node) Selected() bool {
	val, err := this.val.Selected()
	if err != nil {
		this.orb.Quit(err)
	}
	return val
}

func (this *Node) Find(search string) *Node {
	nde, err := this.val.Element(webdriver.FindByCss, search)
	if err != nil {
		this.orb.Quit(err)
	}
	return NewNode(this.orb, this.wds, nde)
}

func (this *Node) Search(search string) []*Node {
	nds, err := this.val.Elements(webdriver.FindByCss, search)
	if err != nil {
		this.orb.Quit(err)
	}
	var out []*Node
	for _, nde := range nds {
		out = append(out, NewNode(this.orb, this.wds, nde))
	}
	return out
}
