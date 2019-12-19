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
	"io"

	"github.com/abcum/orbit"

	"github.com/abcum/otto"

	"github.com/abcum/cirrius/util/args"
)

type Doc struct {
	orb *orbit.Orbit
	lib *Lib
	opt optlist
	ctx *Page
	sch otto.Value
	fnt map[string]bool
	trk struct {
		m []*Font
		p []*Page
		f []*File
		t []*Flow
	}
	Graphics *Graphics
}

func NewDoc(orb *orbit.Orbit, lib *Lib, opt optlist) *Doc {
	return (&Doc{
		orb: orb,
		lib: lib,
		opt: opt,
	}).init()
}

func (this *Doc) init() *Doc {

	this.opt["optimize"] = true
	this.opt["inmemory"] = true
	this.opt["linearize"] = true

	this.fnt = make(map[string]bool)

	this.lib.val.BeginDocument("", cull(this.opt, docOpts))

	this.Graphics = NewGraphics(this.orb, this.lib)

	return this

}

func (this *Doc) find(name string) {

	var err error
	var ret otto.Value
	var val interface{}

	if !this.sch.IsFunction() {
		return
	}

	if _, ok := this.fnt[name]; ok {
		return
	}

	if ret, err = this.sch.Call(this.sch, name); err != nil {
		this.orb.Quit(err)
	}

	if val, err = ret.Export(); err != nil {
		this.orb.Quit(err)
	}

	if rdr, ok := val.(io.Reader); ok {
		this.fnt[name] = true
		this.Font(name, rdr)
	}

}

// Find adds a callback function which is used to load fonts.
//
//	var lib = require('pdf')();
//	var doc = new lib();
// 	doc.find(function(name) {
//		console.log('Searching for font:', name);
//		return require('file').load("fonts/" + name + ".ttf");
//	});
//
func (this *Doc) Find(callback otto.Value) {

	if callback.IsFunction() {
		this.sch = callback
	}

}

// Font loads a font which can be used in a subsequent text area.
// The specified name of the font is the same name which must be
// used in subsequent text calls.
//
//	var lib = require('pdf')();
//	var doc = new lib();
//	doc.font("Arial", require('file').load("arial.ttf"));
//
func (this *Doc) Font(name string, rdr io.Reader) *Font {

	font := NewFont(this.orb, this.lib, this, name, rdr)

	this.trk.m = append(this.trk.m, font)

	return font

}

// Load loads a file for processing, enabling page manipulation,
// copying, and placement.
//
//	var lib = require('pdf')();
//	var doc = new lib();
//	var pdf = doc.Load( require('file').load("other.pdf") );
//	for (var i=1; i<=pdf.pages(); i++) {
//		console.log( pdf.page(i).width(), 'x', pdf.page(i).height() );
//	}
//
func (this *Doc) Load(rdr io.Reader) *File {

	file := NewFile(this.orb, this.lib, this, rdr)

	this.trk.f = append(this.trk.f, file)

	return file

}

func (this *Doc) Page(call otto.FunctionCall) otto.Value {

	args.Size(this.orb, call, 2, 2)

	w := args.Double(this.orb, call, 0)
	h := args.Double(this.orb, call, 1)

	for i := len(this.trk.p) - 1; i >= 0; i-- {
		this.trk.p[i].done()
		this.trk.p = this.trk.p[:len(this.trk.p)-1]
	}

	page := NewPage(this.orb, this.lib, this, w, h)

	this.trk.p = append(this.trk.p, page)

	return args.Value(this.orb, page)

}

func (this *Doc) Flow(call otto.FunctionCall) otto.Value {

	args.Size(this.orb, call, 1, 2)

	text := args.String(this.orb, call, 0)
	opts := args.Object(this.orb, call, 1)

	flow := NewFlow(this.orb, this.lib, this, text, opts)

	this.trk.t = append(this.trk.t, flow)

	return args.Value(this.orb, flow)

}

func (this *Doc) Pipe(wtr io.Writer) {

	var err error
	var out []byte

	// Close flows

	for i := len(this.trk.t) - 1; i >= 0; i-- {
		this.trk.t[i].done()
		this.trk.t = this.trk.t[:len(this.trk.t)-1]
	}

	// Close pages

	for i := len(this.trk.p) - 1; i >= 0; i-- {
		this.trk.p[i].done()
		this.trk.p = this.trk.p[:len(this.trk.p)-1]
	}

	// Close files

	for i := len(this.trk.f) - 1; i >= 0; i-- {
		this.trk.f[i].done()
		this.trk.f = this.trk.f[:len(this.trk.f)-1]
	}

	// Close fonts

	for i := len(this.trk.m) - 1; i >= 0; i-- {
		this.trk.m[i].done()
		this.trk.m = this.trk.m[:len(this.trk.m)-1]
	}

	if err = this.lib.val.EndDocument(""); err != nil {
		this.orb.Quit(err)
	}

	if out, _, err = this.lib.val.GetBuffer(); err != nil {
		this.orb.Quit(err)
	}

	this.lib.val.Delete()

	wtr.Write(out)

}
