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
)

type Table struct {
	orb *orbit.Orbit
	lib *Lib
	doc *Doc
	ref int
}

func NewTable(orb *orbit.Orbit, lib *Lib, doc *Doc) *Table {
	return (&Table{
		orb: orb,
		lib: lib,
		doc: doc,
		ref: -1,
	}).init()
}

func (this *Table) init() *Table {

	return this

}

func (this *Table) done() {

	// Don't do anything as the table will be deleted automatically

}

func (this *Table) Cell(col, row int, text string, opts optlist) *Table {

	var err error

	if this.ref, err = this.lib.val.AddTableCell(this.ref, col, row, text, cull(opts, tcellOpts)); err != nil {
		this.orb.Quit(err)
	}

	return this

}

func (this *Flow) Place(llx, lly, urx, ury float64, opts optlist) {

	if _, err := this.lib.val.FitTable(this.ref, llx, lly, urx, ury, cull(opts, tableOpts)); err != nil {
		this.orb.Quit(err)
	}

}
