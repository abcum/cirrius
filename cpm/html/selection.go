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

package html

import (
	"bytes"
	"io"

	"golang.org/x/net/html"

	"github.com/abcum/orbit"

	"github.com/robertkrimen/otto"

	"github.com/PuerkitoBio/goquery"
)

type Selection struct {
	orb *orbit.Orbit
	val *goquery.Selection
}

func NewSelection(orb *orbit.Orbit, val *goquery.Selection) *Selection {
	return &Selection{
		orb: orb,
		val: val,
	}
}

func (this *Selection) AddAttr(name, value string) *Selection {
	sel := this.val.SetAttr(name, value)
	return NewSelection(this.orb, sel)
}

func (this *Selection) AddClass(class ...string) *Selection {
	sel := this.val.AddClass(class...)
	return NewSelection(this.orb, sel)
}

func (this *Selection) Append(content interface{}) *Selection {
	sel := this.val
	switch c := content.(type) {
	case string:
		if c[0] == '<' {
			sel = this.val.AppendHtml(c)
		}
		sel = this.val.Append(c)
	case *Selection:
		sel = this.val.AppendSelection(c.val)
	}
	return NewSelection(this.orb, sel)
}

func (this *Selection) Attr(name string, value ...string) interface{} {
	if len(value) > 0 {
		sel := this.val.SetAttr(name, value[0])
		return NewSelection(this.orb, sel)
	}
	val, _ := this.val.Attr(name)
	return val
}

func (this *Selection) Children(selector ...string) *Selection {
	if len(selector) > 0 {
		sel := this.val.ChildrenFiltered(selector[0])
		return NewSelection(this.orb, sel)
	}
	sel := this.val.Children()
	return NewSelection(this.orb, sel)
}

func (this *Selection) Clone() *Selection {
	sel := this.val.Clone()
	return NewSelection(this.orb, sel)
}

func (this *Selection) Closest(selector interface{}) *Selection {
	switch s := selector.(type) {
	case string:
		sel := this.val.Closest(s)
		return NewSelection(this.orb, sel)
	case *Selection:
		sel := this.val.ClosestSelection(s.val)
		return NewSelection(this.orb, sel)
	}
	return nil
}

func (this *Selection) Contents(selector ...string) *Selection {
	if len(selector) > 0 {
		sel := this.val.ContentsFiltered(selector[0])
		return NewSelection(this.orb, sel)
	}
	sel := this.val.Contents()
	return NewSelection(this.orb, sel)
}

func (this *Selection) Each(callback otto.Value) *Selection {
	sel := this.val.EachWithBreak(func(i int, s *goquery.Selection) bool {
		val, err := callback.Call(callback, i, NewSelection(this.orb, s))
		if err != nil {
			this.orb.Quit(err)
		}
		out, err := val.ToBoolean()
		if err != nil {
			this.orb.Quit(err)
		}
		return out
	})
	return NewSelection(this.orb, sel)
}

func (this *Selection) Empty() *Selection {
	sel := this.val.Empty()
	return NewSelection(this.orb, sel)
}

func (this *Selection) End() *Selection {
	sel := this.val.End()
	return NewSelection(this.orb, sel)
}

func (this *Selection) Eq(index int) *Selection {
	sel := this.val.Eq(index)
	return NewSelection(this.orb, sel)
}

func (this *Selection) Filter(selector interface{}) *Selection {
	switch s := selector.(type) {
	case string:
		sel := this.val.Filter(s)
		return NewSelection(this.orb, sel)
	case *Selection:
		sel := this.val.FilterSelection(s.val)
		return NewSelection(this.orb, sel)
	case otto.Value:
		sel := this.val.FilterFunction(func(i int, o *goquery.Selection) bool {
			val, err := s.Call(s, i, NewSelection(this.orb, o))
			if err != nil {
				this.orb.Quit(err)
			}
			out, err := val.ToBoolean()
			if err != nil {
				this.orb.Quit(err)
			}
			return out
		})
		return NewSelection(this.orb, sel)
	}
	return nil
}

func (this *Selection) Find(selector interface{}) *Selection {
	switch s := selector.(type) {
	case string:
		sel := this.val.Find(s)
		return NewSelection(this.orb, sel)
	case *Selection:
		sel := this.val.FindSelection(s.val)
		return NewSelection(this.orb, sel)
	}
	return nil
}

func (this *Selection) First() *Selection {
	sel := this.val.First()
	return NewSelection(this.orb, sel)
}

func (this *Selection) Has(selector interface{}) *Selection {
	switch s := selector.(type) {
	case string:
		sel := this.val.Has(s)
		return NewSelection(this.orb, sel)
	case *Selection:
		sel := this.val.HasSelection(s.val)
		return NewSelection(this.orb, sel)
	}
	return nil
}

func (this *Selection) HasClass(class string) bool {
	return this.val.HasClass(class)
}

func (this *Selection) Html(value ...string) interface{} {
	if len(value) > 0 {
		sel := this.val.SetHtml(value[0])
		return NewSelection(this.orb, sel)
	}
	var buf bytes.Buffer
	for _, n := range this.val.Nodes {
		if err := html.Render(&buf, n); err != nil {
			this.orb.Quit(err)
		}
	}
	return buf.String()
}

func (this *Selection) Index(selector ...interface{}) int {
	if len(selector) > 0 {
		switch s := selector[0].(type) {
		case string:
			return this.val.IndexSelector(s)
		case *Selection:
			return this.val.IndexOfSelection(s.val)
		}
	}
	return this.val.Index()
}

func (this *Selection) Is(selector interface{}) bool {
	switch s := selector.(type) {
	case string:
		return this.val.Is(s)
	case *Selection:
		return this.val.IsSelection(s.val)
	case otto.Value:
		return this.val.IsFunction(func(i int, o *goquery.Selection) bool {
			val, err := s.Call(s, i, NewSelection(this.orb, o))
			if err != nil {
				this.orb.Quit(err)
			}
			out, err := val.ToBoolean()
			if err != nil {
				this.orb.Quit(err)
			}
			return out
		})
	}
	return false
}

func (this *Selection) Last() *Selection {
	sel := this.val.Last()
	return NewSelection(this.orb, sel)
}

func (this *Selection) Length() int {
	return this.val.Length()
}

func (this *Selection) Map(callback otto.Value) []string {
	return this.val.Map(func(i int, s *goquery.Selection) string {
		val, err := callback.Call(callback, i, NewSelection(this.orb, s))
		if err != nil {
			this.orb.Quit(err)
		}
		out, err := val.ToString()
		if err != nil {
			this.orb.Quit(err)
		}
		return out
	})
}

func (this *Selection) Next(selector ...string) *Selection {
	if len(selector) > 0 {
		sel := this.val.NextFiltered(selector[0])
		return NewSelection(this.orb, sel)
	}
	sel := this.val.Next()
	return NewSelection(this.orb, sel)
}

func (this *Selection) NextAll(selector ...string) *Selection {
	if len(selector) > 0 {
		sel := this.val.NextAllFiltered(selector[0])
		return NewSelection(this.orb, sel)
	}
	sel := this.val.NextAll()
	return NewSelection(this.orb, sel)
}

func (this *Selection) NextUntil(selectors ...interface{}) *Selection {
	if len(selectors) > 1 {
		switch s := selectors[0].(type) {
		case string:
			if f, ok := selectors[1].(string); ok {
				sel := this.val.NextFilteredUntil(f, s)
				return NewSelection(this.orb, sel)
			}
		case *Selection:
			if f, ok := selectors[1].(string); ok {
				sel := this.val.NextFilteredUntilSelection(f, s.val)
				return NewSelection(this.orb, sel)
			}
		}
	} else if len(selectors) > 0 {
		switch s := selectors[0].(type) {
		case string:
			sel := this.val.NextUntil(s)
			return NewSelection(this.orb, sel)
		case *Selection:
			sel := this.val.NextUntilSelection(s.val)
			return NewSelection(this.orb, sel)
		}
	}
	return this.NextAll()
}

func (this *Selection) Not(selector interface{}) *Selection {
	switch s := selector.(type) {
	case string:
		sel := this.val.Not(s)
		return NewSelection(this.orb, sel)
	case *Selection:
		sel := this.val.NotSelection(s.val)
		return NewSelection(this.orb, sel)
	case otto.Value:
		sel := this.val.NotFunction(func(i int, o *goquery.Selection) bool {
			val, err := s.Call(s, i, NewSelection(this.orb, o))
			if err != nil {
				this.orb.Quit(err)
			}
			out, err := val.ToBoolean()
			if err != nil {
				this.orb.Quit(err)
			}
			return out
		})
		return NewSelection(this.orb, sel)
	}
	return nil
}

func (this *Selection) Parent(selector ...string) *Selection {
	if len(selector) > 0 {
		sel := this.val.ParentFiltered(selector[0])
		return NewSelection(this.orb, sel)
	}
	sel := this.val.Parent()
	return NewSelection(this.orb, sel)
}

func (this *Selection) Parents(selector ...string) *Selection {
	if len(selector) > 0 {
		sel := this.val.ParentsFiltered(selector[0])
		return NewSelection(this.orb, sel)
	}
	sel := this.val.Parents()
	return NewSelection(this.orb, sel)
}

func (this *Selection) ParentsUntil(selectors ...interface{}) *Selection {
	if len(selectors) > 1 {
		switch s := selectors[0].(type) {
		case string:
			if f, ok := selectors[1].(string); ok {
				sel := this.val.ParentsFilteredUntil(f, s)
				return NewSelection(this.orb, sel)
			}
		case *Selection:
			if f, ok := selectors[1].(string); ok {
				sel := this.val.ParentsFilteredUntilSelection(f, s.val)
				return NewSelection(this.orb, sel)
			}
		}
	} else if len(selectors) > 0 {
		switch s := selectors[0].(type) {
		case string:
			sel := this.val.ParentsUntil(s)
			return NewSelection(this.orb, sel)
		case *Selection:
			sel := this.val.ParentsUntilSelection(s.val)
			return NewSelection(this.orb, sel)
		}
	}
	return this.Parents()
}

func (this *Selection) Prepend(content interface{}) *Selection {
	sel := this.val
	switch c := content.(type) {
	case string:
		if c[0] == '<' {
			sel = this.val.PrependHtml(c)
		}
		sel = this.val.Prepend(c)
	case *Selection:
		sel = this.val.PrependSelection(c.val)
	}
	return NewSelection(this.orb, sel)
}

func (this *Selection) Prev(selector ...string) *Selection {
	if len(selector) > 0 {
		sel := this.val.NextFiltered(selector[0])
		return NewSelection(this.orb, sel)
	}
	sel := this.val.Next()
	return NewSelection(this.orb, sel)
}

func (this *Selection) PrevAll(selector ...string) *Selection {
	if len(selector) > 0 {
		sel := this.val.PrevAllFiltered(selector[0])
		return NewSelection(this.orb, sel)
	}
	sel := this.val.PrevAll()
	return NewSelection(this.orb, sel)
}

func (this *Selection) PrevUntil(selectors ...interface{}) *Selection {
	if len(selectors) > 1 {
		switch s := selectors[0].(type) {
		case string:
			if f, ok := selectors[1].(string); ok {
				sel := this.val.PrevFilteredUntil(f, s)
				return NewSelection(this.orb, sel)
			}
		case *Selection:
			if f, ok := selectors[1].(string); ok {
				sel := this.val.PrevFilteredUntilSelection(f, s.val)
				return NewSelection(this.orb, sel)
			}
		}
	} else if len(selectors) > 0 {
		switch s := selectors[0].(type) {
		case string:
			sel := this.val.PrevUntil(s)
			return NewSelection(this.orb, sel)
		case *Selection:
			sel := this.val.PrevUntilSelection(s.val)
			return NewSelection(this.orb, sel)
		}
	}
	return this.PrevAll()
}

func (this *Selection) Remove(selector ...string) *Selection {
	if len(selector) > 0 {
		sel := this.val.RemoveFiltered(selector[0])
		return NewSelection(this.orb, sel)
	}
	sel := this.val.Remove()
	return NewSelection(this.orb, sel)
}

func (this *Selection) ReplaceWith(selector interface{}) *Selection {
	sel := this.val
	switch s := selector.(type) {
	case string:
		if s[0] == '<' {
			sel = this.val.ReplaceWithHtml(s)
		}
		sel = this.val.ReplaceWith(s)
	case *Selection:
		sel = this.val.ReplaceWithSelection(s.val)
	}
	return NewSelection(this.orb, sel)
}

func (this *Selection) RemoveAttr(name string) *Selection {
	sel := this.val.RemoveAttr(name)
	return NewSelection(this.orb, sel)
}

func (this *Selection) RemoveClass(class ...string) *Selection {
	sel := this.val.RemoveClass(class...)
	return NewSelection(this.orb, sel)
}

func (this *Selection) Siblings(selector ...string) *Selection {
	if len(selector) > 0 {
		sel := this.val.SiblingsFiltered(selector[0])
		return NewSelection(this.orb, sel)
	}
	sel := this.val.Siblings()
	return NewSelection(this.orb, sel)
}

func (this *Selection) Slice(start, end int) *Selection {
	sel := this.val.Slice(start, end)
	return NewSelection(this.orb, sel)
}

func (this *Selection) Text(value ...string) interface{} {
	if len(value) > 0 {
		sel := this.val.SetText(value[0])
		return NewSelection(this.orb, sel)
	}
	return this.val.Text()
}

func (this *Selection) Type() (data string) {
	for _, n := range this.val.Nodes {
		return n.Data
	}
	return
}

func (this *Selection) ToggleClass(class ...string) *Selection {
	sel := this.val.ToggleClass(class...)
	return NewSelection(this.orb, sel)
}

func (this *Selection) Union(selection *Selection) *Selection {
	sel := this.val.Union(selection.val)
	return NewSelection(this.orb, sel)
}

func (this *Selection) Unwrap() *Selection {
	sel := this.val.Unwrap()
	return NewSelection(this.orb, sel)
}

func (this *Selection) Wrap(selector interface{}) *Selection {
	sel := this.val
	switch s := selector.(type) {
	case string:
		if s[0] == '<' {
			sel = this.val.WrapHtml(s)
		}
		sel = this.val.Wrap(s)
	case *Selection:
		sel = this.val.WrapSelection(s.val)
	}
	return NewSelection(this.orb, sel)
}

func (this *Selection) WrapAll(selector interface{}) *Selection {
	sel := this.val
	switch s := selector.(type) {
	case string:
		if s[0] == '<' {
			sel = this.val.WrapAllHtml(s)
		}
		sel = this.val.WrapAll(s)
	case *Selection:
		sel = this.val.WrapAllSelection(s.val)
	}
	return NewSelection(this.orb, sel)
}

func (this *Selection) WrapInner(selector interface{}) *Selection {
	sel := this.val
	switch s := selector.(type) {
	case string:
		if s[0] == '<' {
			sel = this.val.WrapInnerHtml(s)
		}
		sel = this.val.WrapInner(s)
	case *Selection:
		sel = this.val.WrapInnerSelection(s.val)
	}
	return NewSelection(this.orb, sel)
}

func (this *Selection) Pipe(w io.Writer) {
	var buf bytes.Buffer
	for _, n := range this.val.Nodes {
		if err := html.Render(&buf, n); err != nil {
			this.orb.Quit(err)
		}
	}
	if _, err := buf.WriteTo(w); err != nil {
		this.orb.Quit(err)
	}
}
