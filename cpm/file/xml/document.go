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
	"io"

	"github.com/abcum/orbit"
	"github.com/beevik/etree"
)

type Document struct {
	orb *orbit.Orbit
	val *etree.Document
}

func NewDocument(orb *orbit.Orbit, val *etree.Document) *Document {
	return &Document{
		orb: orb,
		val: val,
	}
}

func (this *Document) Root() *Element {
	return NewElement(this.orb, this.val.Root())
}

func (this *Document) Load(r io.Reader) *Document {
	if _, err := this.val.ReadFrom(r); err != nil {
		this.orb.Quit(err)
	}
	return this
}

func (this *Document) Pipe(w io.Writer) {
	this.val.IndentTabs()
	if _, err := this.val.WriteTo(w); err != nil {
		this.orb.Quit(err)
	}
	return
}

// --------------------------------------------------

func (this *Document) Comments() []*Comment {
	var out []*Comment
	for _, c := range this.val.Child {
		if v, ok := c.(*etree.Comment); ok {
			out = append(out, NewComment(this.orb, v))
		}
	}
	return out
}

func (this *Document) Comment(value string) *Comment {
	return NewComment(this.orb, this.val.CreateComment(value))
}

func (this *Document) Directives() []*Directive {
	var out []*Directive
	for _, c := range this.val.Child {
		if v, ok := c.(*etree.Directive); ok {
			out = append(out, NewDirective(this.orb, v))
		}
	}
	return out
}

func (this *Document) Directive(value string) *Directive {
	return NewDirective(this.orb, this.val.CreateDirective(value))
}

func (this *Document) Elements() []*Element {
	var out []*Element
	for _, c := range this.val.Child {
		if v, ok := c.(*etree.Element); ok {
			out = append(out, NewElement(this.orb, v))
		}
	}
	return out
}

func (this *Document) Element(value string) *Element {
	return NewElement(this.orb, this.val.CreateElement(value))
}

func (this *Document) Instructions() []*Instruction {
	var out []*Instruction
	for _, c := range this.val.Child {
		if v, ok := c.(*etree.ProcInst); ok {
			out = append(out, NewInstruction(this.orb, v))
		}
	}
	return out
}

func (this *Document) Instruction(key, value string) *Instruction {
	return NewInstruction(this.orb, this.val.CreateProcInst(key, value))
}

// --------------------------------------------------

func (this *Document) Find(path string) *Element {
	return NewElement(this.orb, this.val.FindElement(path))
}

func (this *Document) Search(path string) []*Element {
	var out []*Element
	for _, e := range this.val.FindElements(path) {
		out = append(out, NewElement(this.orb, e))
	}
	return out
}

func (this *Document) Text(val ...string) string {
	if len(val) > 0 {
		this.val.SetText(val[0])
	}
	return this.val.Text()
}
