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
	"io/ioutil"

	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"

	"github.com/abcum/cirrius/util/args"
)

type File struct {
	orb *orbit.Orbit
	lib *Lib
	doc *Doc
	rdr io.Reader
	ref int
	trk struct {
		l []*Leaf
	}
}

func NewFile(orb *orbit.Orbit, lib *Lib, doc *Doc, rdr io.Reader) *File {
	return (&File{
		orb: orb,
		lib: lib,
		doc: doc,
		rdr: rdr,
		ref: -1,
	}).init()
}

func (this *File) init() *File {

	var err error
	var res []byte

	id := uniq()

	if res, err = ioutil.ReadAll(this.rdr); err != nil {
		this.orb.Quit(err)
	}

	if err = this.lib.val.CreatePvf(id, res, ""); err != nil {
		this.orb.Quit(err)
	}

	if this.ref, err = this.lib.val.OpenPdiDocument(id, ""); err != nil {
		this.orb.Quit(err)
	}

	if err = this.lib.val.DeletePvf(id); err != nil {
		this.orb.Quit(err)
	}

	return this

}

func (this *File) done() {

	for i := len(this.trk.l) - 1; i >= 0; i-- {
		this.trk.l[i].done()
		this.trk.l = this.trk.l[:len(this.trk.l)-1]
	}

	if err := this.lib.val.ClosePdiDocument(this.ref); err != nil {
		this.orb.Quit(err)
	}

}

func (this *File) Page(call otto.FunctionCall) otto.Value {

	args.Size(this.orb, call, 1, 1)

	page := args.Number(this.orb, call, 0)

	leaf := NewLeaf(this.orb, this.lib, this.ref, page)

	this.trk.l = append(this.trk.l, leaf)

	return args.Value(this.orb, leaf)

}

func (this *File) Pages(call otto.FunctionCall) otto.Value {

	args.Size(this.orb, call, 0, 0)

	val, err := this.lib.val.PcosGetNumber(this.ref, "length:pages")
	if err != nil {
		this.orb.Quit(err)
	}

	return args.Value(this.orb, val)

}
