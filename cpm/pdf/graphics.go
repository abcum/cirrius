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
	"github.com/robertkrimen/otto"

	"github.com/abcum/cirrius/util/args"
)

type Graphics struct {
	orb *orbit.Orbit
	lib *Lib
}

func NewGraphics(orb *orbit.Orbit, lib *Lib) *Graphics {
	return (&Graphics{
		orb: orb,
		lib: lib,
	}).init()
}

func (this *Graphics) init() *Graphics {

	return this

}

func (this *Graphics) Save() {

	if err := this.lib.val.Save(); err != nil {
		this.orb.Quit(err)
	}

}

func (this *Graphics) Restore() {

	if err := this.lib.val.Restore(); err != nil {
		this.orb.Quit(err)
	}

}

func (this *Graphics) Set(state *State) {

	if err := this.lib.val.SetGstate(state.ref); err != nil {
		this.orb.Quit(err)
	}

}

func (this *Graphics) New(call otto.FunctionCall) otto.Value {

	args.Size(this.orb, call, 1, 1)

	opts := args.Object(this.orb, call, 0)

	stat := NewState(this.orb, this.lib, opts)

	return args.Value(this.orb, stat)

}

func (this *Graphics) Use(call otto.FunctionCall) otto.Value {

	args.Size(this.orb, call, 1, 1)

	opts := args.Object(this.orb, call, 0)

	stat := NewState(this.orb, this.lib, opts)

	if err := this.lib.val.SetGstate(stat.ref); err != nil {
		this.orb.Quit(err)
	}

	return args.Value(this.orb, stat)

}
