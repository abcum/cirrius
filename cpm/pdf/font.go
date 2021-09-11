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

//go:build cgo
// +build cgo

package pdf

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/abcum/orbit"
)

type Font struct {
	orb *orbit.Orbit
	lib *Lib
	doc *Doc
	fnt string
	rdr io.Reader
	ref int
}

func NewFont(orb *orbit.Orbit, lib *Lib, doc *Doc, fnt string, rdr io.Reader) *Font {
	return (&Font{
		orb: orb,
		lib: lib,
		doc: doc,
		fnt: fnt,
		rdr: rdr,
		ref: -1,
	}).init()
}

func (this *Font) init() *Font {

	var err error
	var res []byte

	id := fmt.Sprintf("%s.ttf", this.fnt)

	if res, err = ioutil.ReadAll(this.rdr); err != nil {
		this.orb.Quit(err)
	}

	if err = this.lib.val.CreatePvf(id, res, "copy"); err != nil {
		this.orb.Quit(err)
	}

	if this.ref, err = this.lib.val.LoadFont(this.fnt, "unicode", "embedding={true} subsetting={true}"); err != nil {
		this.orb.Quit(err)
	}

	if err = this.lib.val.DeletePvf(id); err != nil {
		this.orb.Quit(err)
	}

	return this

}

func (this *Font) done() {

	// Don't do anything as the font will be deleted automatically

}
