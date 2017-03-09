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
	"github.com/abcum/orbit"
	"github.com/beevik/etree"
)

type Attribute struct {
	orb *orbit.Orbit
	atr *etree.Attr
}

func NewAttribute(orb *orbit.Orbit, atr *etree.Attr) *Attribute {
	return &Attribute{
		orb: orb,
		atr: atr,
	}
}

func (this *Attribute) Key(key ...string) string {
	if len(key) > 0 {
		this.atr.Key = key[0]
	}
	return this.atr.Key
}

func (this *Attribute) Val(val ...string) string {
	if len(val) > 0 {
		this.atr.Value = val[0]
	}
	return this.atr.Value
}
