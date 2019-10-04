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

package xml

import (
	"github.com/abcum/orbit"
	"github.com/beevik/etree"
)

type Element struct {
	orb *orbit.Orbit
	val *etree.Element
}

func NewElement(orb *orbit.Orbit, val *etree.Element) *Element {
	return &Element{
		orb: orb,
		val: val,
	}
}

func (this *Element) Tag() string {
	return this.val.Tag
}

func (this *Element) Set(key, val string) *Attribute {
	return NewAttribute(this.orb, this.val.CreateAttr(key, val))
}

func (this *Element) Get(key string) *Attribute {
	return NewAttribute(this.orb, this.val.SelectAttr(key))
}

func (this *Element) Del(key string) *Attribute {
	return NewAttribute(this.orb, this.val.RemoveAttr(key))
}

func (this *Element) Delete() {
	this.val.Parent().RemoveChild(this.val)
}

func (this *Element) Parent() *Element {
	return NewElement(this.orb, this.val.Parent())
}

// --------------------------------------------------

func (this *Element) Comments() []*Comment {
	var out []*Comment
	for _, c := range this.val.Child {
		if v, ok := c.(*etree.Comment); ok {
			out = append(out, NewComment(this.orb, v))
		}
	}
	return out
}

func (this *Element) Comment(value string) *Comment {
	return NewComment(this.orb, this.val.CreateComment(value))
}

func (this *Element) Directives() []*Directive {
	var out []*Directive
	for _, c := range this.val.Child {
		if v, ok := c.(*etree.Directive); ok {
			out = append(out, NewDirective(this.orb, v))
		}
	}
	return out
}

func (this *Element) Directive(value string) *Directive {
	return NewDirective(this.orb, this.val.CreateDirective(value))
}

func (this *Element) Elements() []*Element {
	var out []*Element
	for _, c := range this.val.Child {
		if v, ok := c.(*etree.Element); ok {
			out = append(out, NewElement(this.orb, v))
		}
	}
	return out
}

func (this *Element) Element(value string) *Element {
	return NewElement(this.orb, this.val.CreateElement(value))
}

func (this *Element) Instructions() []*Instruction {
	var out []*Instruction
	for _, c := range this.val.Child {
		if v, ok := c.(*etree.ProcInst); ok {
			out = append(out, NewInstruction(this.orb, v))
		}
	}
	return out
}

func (this *Element) Instruction(key, value string) *Instruction {
	return NewInstruction(this.orb, this.val.CreateProcInst(key, value))
}

// --------------------------------------------------

func (this *Element) Find(path string) *Element {
	return NewElement(this.orb, this.val.FindElement(path))
}

func (this *Element) Search(path string) []*Element {
	var out []*Element
	for _, e := range this.val.FindElements(path) {
		out = append(out, NewElement(this.orb, e))
	}
	return out
}

func (this *Element) Text(val ...string) string {
	if len(val) > 0 {
		this.val.SetText(val[0])
	}
	return this.val.Text()
}
