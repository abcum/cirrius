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
	"fmt"

	"github.com/abcum/orbit"
)

type Leaf struct {
	orb *orbit.Orbit
	lib *Lib
	doc int
	pge int
	ref int
}

func NewLeaf(orb *orbit.Orbit, lib *Lib, doc, pge int) *Leaf {
	return (&Leaf{
		orb: orb,
		lib: lib,
		doc: doc,
		pge: pge,
		ref: -1,
	}).init()
}

func (this *Leaf) init() *Leaf {

	var err error

	if this.ref, err = this.lib.val.OpenPdiPage(this.doc, this.pge, ""); err != nil {
		this.orb.Quit(err)
	}

	return this

}

func (this *Leaf) done() {

	if err := this.lib.val.ClosePdiPage(this.ref); err != nil {
		this.orb.Quit(err)
	}

}

func (this *Leaf) Width() float64 {

	val, err := this.lib.val.PcosGetNumber(this.doc, fmt.Sprintf("pages[%d]/width", this.pge-1))
	if err != nil {
		this.orb.Quit(err)
	}

	return val

}

func (this *Leaf) Height() float64 {

	val, err := this.lib.val.PcosGetNumber(this.doc, fmt.Sprintf("pages[%d]/height", this.pge-1))
	if err != nil {
		this.orb.Quit(err)
	}

	return val

}
