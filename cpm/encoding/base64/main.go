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

package base64

import (
	"io"
	"io/ioutil"

	"encoding/base64"

	"github.com/abcum/orbit"
)

func init() {
	orbit.Add("encoding/base64", New)
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
	cdc := base64.NewEncoder(base64.StdEncoding, w)
	return NewEncoder(this.orb, cdc)
}

func (this *Module) Decoder(r io.Reader) *Decoder {
	cdc := base64.NewDecoder(base64.StdEncoding, r)
	return NewDecoder(this.orb, cdc)
}

func (this *Module) Encode(dec string) (enc string) {
	return base64.StdEncoding.EncodeToString([]byte(dec))
}

func (this *Module) Decode(enc interface{}) (dec string) {
	switch v := enc.(type) {
	case io.Reader:
		decoder := base64.NewDecoder(base64.StdEncoding, v)
		bit, err := ioutil.ReadAll(decoder)
		if err != nil {
			this.orb.Quit(err)
		}
		return string(bit)
	case string:
		bit, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			this.orb.Quit(err)
		}
		return string(bit)
	}
	return dec
}
