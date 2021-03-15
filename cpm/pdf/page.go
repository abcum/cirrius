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

// +build cgo

package pdf

import (
	"github.com/abcum/orbit"
	"github.com/abcum/otto"

	"github.com/abcum/cirrius/util/args"

	"github.com/abcum/cirrius/cpm/colour"
)

type Page struct {
	orb *orbit.Orbit
	lib *Lib
	doc *Doc
	w   float64
	h   float64
	trk struct {
		f []*Flow
	}
}

func NewPage(orb *orbit.Orbit, lib *Lib, doc *Doc, w, h float64) *Page {
	return (&Page{
		orb: orb,
		lib: lib,
		doc: doc,
		w:   w,
		h:   h,
	}).init()
}

func (this *Page) init() *Page {

	if err := this.lib.val.BeginPageExt(this.w, this.h, ""); err != nil {
		this.orb.Quit(err)
	}

	return this

}

func (this *Page) done() *Page {

	if err := this.lib.val.EndPageExt(""); err != nil {
		this.orb.Quit(err)
	}

	return this

}

func (this *Page) Text(call otto.FunctionCall) otto.Value {

	var err error

	args.Size(this.orb, call, 3, 4)

	t := args.String(this.orb, call, 0)
	x := args.Double(this.orb, call, 1)
	y := args.Double(this.orb, call, 2)
	o := args.Object(this.orb, call, 3)

	if val, ok := o["fontname"].(string); ok {
		this.doc.find(val)
	}

	if err = this.lib.val.FitTextline(t, x, y, cull(o, textOpts)); err != nil {
		this.orb.Quit(err)
	}

	return args.Value(this.orb, this)

}

func (this *Page) Image(call otto.FunctionCall) otto.Value {

	var ref int
	var err error

	args.Size(this.orb, call, 3, 5)

	i := args.Image(this.orb, call, 0)
	x := args.Double(this.orb, call, 1)
	y := args.Double(this.orb, call, 2)
	o := args.Object(this.orb, call, 3)
	n := args.Number(this.orb, call, 4)

	if len(call.ArgumentList) == 5  {

		if err = this.lib.val.FitImage(n, x, y, cull(o, imageOpts)); err != nil {
			this.orb.Quit(err)
		}

	} else {

		id := uniq()

		if err = this.lib.val.CreatePvf(id, i.Bytes(), ""); err != nil {
			this.orb.Quit(err)
		}

		if ref, err = this.lib.val.LoadImage("auto", id, cull(o, loadOpts)); err != nil {
			this.orb.Quit(err)
		}

		if err = this.lib.val.FitImage(ref, x, y, cull(o, imageOpts)); err != nil {
			this.orb.Quit(err)
		}

		if this.lib.val.CloseImage(ref); err != nil {
			this.orb.Quit(err)
		}

		if err = this.lib.val.DeletePvf(id); err != nil {
			this.orb.Quit(err)
		}

	}

	return args.Value(this.orb, this)

}

func (this *Page) Block(call otto.FunctionCall) otto.Value {

	var ref int
	var err error

	args.Size(this.orb, call, 5, 6)

	i := args.Image(this.orb, call, 0)
	x := args.Double(this.orb, call, 1)
	y := args.Double(this.orb, call, 2)
	w := args.Double(this.orb, call, 3)
	h := args.Double(this.orb, call, 4)
	o := args.Object(this.orb, call, 5)

	id := uniq()

	if val, ok := o["backgroundcolor"]; ok {
		if col, ok := val.(*colour.Colour); ok {
			k, v1, v2, v3, v4 := col.Output()
			this.lib.val.SetColor("fill", k, v1, v2, v3, v4)
			this.lib.val.Rect(x, y, w, h)
			this.lib.val.Fill()
		}
	}

	if err = this.lib.val.CreatePvf(id, i.Bytes(), ""); err != nil {
		this.orb.Quit(err)
	}

	if ref, err = this.lib.val.LoadImage("auto", id, cull(o, loadOpts)); err != nil {
		this.orb.Quit(err)
	}

	if err = this.lib.val.FitImage(ref, x, y, cull(o, imageOpts)); err != nil {
		this.orb.Quit(err)
	}

	if this.lib.val.CloseImage(ref); err != nil {
		this.orb.Quit(err)
	}

	if err = this.lib.val.DeletePvf(id); err != nil {
		this.orb.Quit(err)
	}

	return args.Value(this.orb, this)

}

func (this *Page) Place(call otto.FunctionCall) otto.Value {

	args.Size(this.orb, call, 3, 4)

	l, _ := args.Any(this.orb, call, 0).(*Leaf)
	x := args.Double(this.orb, call, 1)
	y := args.Double(this.orb, call, 2)
	o := args.Object(this.orb, call, 3)

	if err := this.lib.val.FitPdiPage(l.ref, x, y, cull(o, pageOpts)); err != nil {
		this.orb.Quit(err)
	}

	return args.Value(this.orb, this)

}
