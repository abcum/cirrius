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

type Decoder struct {
	orb *orbit.Orbit
	rdr io.Reader
}

func NewDecoder(orb *orbit.Orbit, rdr io.Reader) *Decoder {
	return &Decoder{
		orb: orb,
		rdr: rdr,
	}
}

func (this *Decoder) Decode() (dec interface{}) {
	buf, err := ioutil.ReadAll(this.rdr)
	if err != nil {
		this.orb.Quit(err)
	}
	if err := hjson.Unmarshal(buf, &dec); err != nil {
		this.orb.Quit(err)
	}
	return dec
}
