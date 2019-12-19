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

type Flow struct {
	orb *orbit.Orbit
	lib *Lib
	doc *Doc
	txt string
	opt optlist
	ref int
}

func NewFlow(orb *orbit.Orbit, lib *Lib, doc *Doc, txt string, opt optlist) *Flow {
	return (&Flow{
		orb: orb,
		lib: lib,
		doc: doc,
		txt: txt,
		opt: opt,
		ref: -1,
	}).init()
}

func (this *Flow) init() *Flow {

	var err error

	if val, ok := this.opt["fontname"].(string); ok {
		this.doc.find(val)
	}

	if this.ref, err = this.lib.val.CreateTextflow(this.txt, cull(this.opt, textOpts)); err != nil {
		this.orb.Quit(err)
	}

	return this

}

func (this *Flow) done() {

	if err := this.lib.val.DeleteTextflow(this.ref); err != nil {
		this.orb.Quit(err)
	}

}

func (this *Flow) Box(call otto.FunctionCall) otto.Value {

	args.Size(this.orb, call, 4, 5)

	x := args.Double(this.orb, call, 0)
	y := args.Double(this.orb, call, 1)
	r := args.Double(this.orb, call, 2)
	t := args.Double(this.orb, call, 3)
	o := args.Object(this.orb, call, 4)

	if val, ok := o["backgroundcolor"]; ok {
		if col, ok := val.(*colour.Colour); ok {
			k, v1, v2, v3, v4 := col.Output()
			this.lib.val.SetColor("fill", k, v1, v2, v3, v4)
			this.lib.val.Rect(x, y, r-x, t-y)
			this.lib.val.Fill()
		}
	}

	if _, err := this.lib.val.FitTextflow(this.ref, x, y, r, t, cull(o, flowOpts)); err != nil {
		this.orb.Quit(err)
	}

	return args.Value(this.orb, this)

}
