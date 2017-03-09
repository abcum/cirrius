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

package allot

import (
	"github.com/abcum/orbit"
	"github.com/sbstjn/allot"
)

type Match struct {
	orb *orbit.Orbit
	val allot.MatchInterface
}

func NewMatch(orb *orbit.Orbit, val allot.MatchInterface) *Match {
	return &Match{
		orb: orb,
		val: val,
	}
}

func (this *Match) Match(pos int) string {

	str, err := this.val.Match(pos)
	if err != nil {
		this.orb.Quit(err)
	}

	return str

}

func (this *Match) Integer(name string) int {

	num, err := this.val.Integer(name)
	if err != nil {
		this.orb.Quit(err)
	}

	return num

}

func (this *Match) String(name string) string {

	str, err := this.val.String(name)
	if err != nil {
		this.orb.Quit(err)
	}

	return str

}

func (this *Match) Param(name, kind string) string {

	par := allot.NewParameterWithType(name, kind)

	str, err := this.val.Parameter(par)
	if err != nil {
		this.orb.Quit(err)
	}

	return str

}
