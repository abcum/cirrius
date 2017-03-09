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

package json

import (
	"bytes"
	"io"
	"strings"

	"github.com/abcum/orbit"
	"github.com/ugorji/go/codec"
)

func init() {
	orbit.Add("encoding/json", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
	}
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) Encoder(w io.Writer) *Encoder {
	cdc := codec.NewEncoder(w, &opt)
	return NewEncoder(this.orb, cdc)
}

func (this *Module) Decoder(r io.Reader) *Decoder {
	cdc := codec.NewDecoder(r, &opt)
	return NewDecoder(this.orb, cdc)
}

func (this *Module) Encode(dec interface{}) (enc string) {
	var buf bytes.Buffer
	if err := codec.NewEncoder(&buf, &opt).Encode(dec); err != nil {
		this.orb.Quit(err)
	}
	return buf.String()
}

func (this *Module) Decode(enc interface{}) (dec interface{}) {
	var r io.Reader
	switch v := enc.(type) {
	case io.Reader:
		r = v
	case string:
		r = strings.NewReader(v)
	}
	if err := codec.NewDecoder(r, &opt).Decode(&dec); err != nil {
		this.orb.Quit(err)
	}
	return dec
}
