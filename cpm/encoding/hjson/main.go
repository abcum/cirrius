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

package hjson

import (
	"io"
	"io/ioutil"

	"github.com/abcum/orbit"
	"github.com/hjson/hjson-go"
)

func init() {
	orbit.Add("encoding/hjson", New)
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
	return NewEncoder(this.orb, w)
}

func (this *Module) Decoder(r io.Reader) *Decoder {
	return NewDecoder(this.orb, r)
}

func (this *Module) Encode(dec interface{}) (enc string) {
	buf, err := hjson.Marshal(dec)
	if err != nil {
		this.orb.Quit(err)
	}
	return string(buf)
}

func (this *Module) Decode(enc interface{}) (dec interface{}) {
	switch v := enc.(type) {
	case string:
		if err := hjson.Unmarshal([]byte(v), &dec); err != nil {
			this.orb.Quit(err)
		}
	case io.Reader:
		buf, err := ioutil.ReadAll(v)
		if err != nil {
			this.orb.Quit(err)
		}
		if err := hjson.Unmarshal(buf, &dec); err != nil {
			this.orb.Quit(err)
		}
	}
	return dec
}
