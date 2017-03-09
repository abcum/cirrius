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
	"bytes"
	"io"

	"github.com/abcum/orbit"
	"github.com/hjson/hjson-go"
)

type Encoder struct {
	orb *orbit.Orbit
	wtr io.Writer
}

func NewEncoder(orb *orbit.Orbit, wtr io.Writer) *Encoder {
	return &Encoder{
		orb: orb,
		wtr: wtr,
	}
}

func (this *Encoder) Encode(dec interface{}) {
	buf, err := hjson.Marshal(dec)
	if err != nil {
		this.orb.Quit(err)
	}
	if _, err := bytes.NewBuffer(buf).WriteTo(this.wtr); err != nil {
		this.orb.Quit(err)
	}
	return
}
