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

package shortid

import (
	"github.com/abcum/orbit"

	"github.com/SKAhack/go-shortid"
)

func init() {
	orbit.Add("id/shortid", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
		gen: shortid.Generator(),
	}
}

type Module struct {
	orb *orbit.Orbit
	gen *shortid.Gen
}

func (this *Module) Seed(seed float64) {
	this.gen.SetSeed(seed)
}

func (this *Module) Generate() (v string) {
	return this.gen.Generate()
}

func (this *Module) Characters(chars string) {
	if err := this.gen.SetCharacters(chars); err != nil {
		this.orb.Quit(err)
	}
}
