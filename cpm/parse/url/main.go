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

package url

import (
	"net/url"

	"github.com/abcum/orbit"
)

func init() {
	orbit.Add("parse/url", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
	}
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) Parse(addr string) *Info {

	return this.parse(addr)

}

func (this *Module) parse(addr string) *Info {

	parsed, err := url.Parse(addr)
	if err != nil {
		this.orb.Quit(err)
	}

	pass, _ := parsed.User.Password()

	return &Info{
		Scheme:   parsed.Scheme,
		User:     parsed.User.Username(),
		Pass:     pass,
		Host:     parsed.Hostname(),
		Port:     parsed.Port(),
		Domain:   parsed.Host,
		Path:     parsed.Path,
		Query:    parsed.Query(),
		Querys:   parsed.RawQuery,
		Fragment: parsed.Fragment,
	}

}
