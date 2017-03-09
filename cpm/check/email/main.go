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

package email

import (
	"strings"

	"net/mail"

	"github.com/abcum/orbit"
)

func init() {
	orbit.Add("check/email", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
	}
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) Check(value string) (info *Info) {

	info = new(Info)

	if !eregex.MatchString(value) {
		return
	}

	val, err := mail.ParseAddress(value)
	if err != nil {
		return
	}

	pos := strings.LastIndexByte(val.Address, '@')

	info.Valid = true
	info.User = val.Address[:pos]
	info.Host = val.Address[pos+1:]
	info.Address = val.Address

	return

}
