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

type Comment struct {
	orb *orbit.Orbit
	val *etree.Comment
}

func NewComment(orb *orbit.Orbit, val *etree.Comment) *Comment {
	return &Comment{
		orb: orb,
		val: val,
	}
}

func (this *Comment) Delete() {
	this.val.Parent().RemoveChild(this.val)
}

func (this *Comment) Parent() *Element {
	return NewElement(this.orb, this.val.Parent())
}

func (this *Comment) Val(val ...string) string {
	if len(val) > 0 {
		this.val.Data = val[0]
	}
	return this.val.Data
}
