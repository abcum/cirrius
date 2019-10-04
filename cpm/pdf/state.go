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

type State struct {
	orb *orbit.Orbit
	lib *Lib
	opt optlist
	ref int
}

func NewState(orb *orbit.Orbit, lib *Lib, opt optlist) *State {
	return (&State{
		orb: orb,
		lib: lib,
		opt: opt,
		ref: -1,
	}).init()
}

func (this *State) init() *State {

	var err error

	if this.ref, err = this.lib.val.CreateGstate(cull(this.opt, stateOpts)); err != nil {
		this.orb.Quit(err)
	}

	return this

}
