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

package xml

import (
	"io"
	"strings"

	"github.com/abcum/orbit"
	"github.com/beevik/etree"
)

func init() {
	orbit.Add("format/xml", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
	}
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) New() *Document {
	return NewDocument(this.orb, etree.NewDocument())
}

func (this *Module) Load(i interface{}) *Document {
	var r io.Reader
	switch s := i.(type) {
	case io.Reader:
		r = s
	case string:
		r = strings.NewReader(s)
	default:
		this.orb.Quit(inputErr)
	}
	return NewDocument(this.orb, etree.NewDocument()).Load(r)
}
