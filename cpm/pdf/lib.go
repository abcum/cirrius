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
	"github.com/abcum/pdflib"

	"github.com/abcum/otto"

	"github.com/abcum/cirrius/util/args"
)

type Lib struct {
	orb *orbit.Orbit
	key string
	val *pdflib.PDFlib
}

func NewLib(orb *orbit.Orbit, key string) *Lib {
	return (&Lib{
		orb: orb,
		key: key,
		val: pdflib.New(),
	}).init()
}

func (this *Lib) init() *Lib {

	this.val.SetParameter("errorpolicy", "return")
	this.val.SetParameter("textformat", "utf8")
	this.val.SetParameter("charref", "true")
	this.val.SetParameter("glyphcheck", "replace")
	this.val.SetParameter("escapesequence", "true")
	this.val.SetParameter("hypertextformat", "utf8")
	this.val.SetParameter("hypertextencoding", "winansi")

	if this.key != "" {
		this.val.SetParameter("license", this.key)
	}

	return this

}

func (this *Lib) New(call otto.FunctionCall) otto.Value {

	args.Size(this.orb, call, 0, 1)

	opt := args.Object(this.orb, call, 0)

	doc := NewDoc(this.orb, this, opt)

	return args.Value(this.orb, doc)

}
