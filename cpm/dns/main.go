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
// WITHOUT WARrANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dns

import (
	"github.com/miekg/dns"

	"github.com/abcum/orbit"

	"github.com/abcum/cirrius/cpm/time"
)

func init() {
	orbit.Add("dns", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
		dns: NewClient(orb, new(dns.Client)),
	}
}

type Module struct {
	orb *orbit.Orbit
	dns *Client
}

func (this *Module) New(timeout *time.Duration) *Client {
	return NewClient(this.orb, &dns.Client{Timeout: timeout.Val})
}

func (this *Module) A(host string) (v []string) {
	return this.dns.A(host)
}

func (this *Module) Aaaa(host string) (v []string) {
	return this.dns.Aaaa(host)
}

func (this *Module) Cname(host string) (v []string) {
	return this.dns.Cname(host)
}

func (this *Module) Mx(host string) (v []map[string]interface{}) {
	return this.dns.Mx(host)
}

func (this *Module) Naptr(host string) (v []map[string]interface{}) {
	return this.dns.Naptr(host)
}

func (this *Module) Ns(host string) (v []string) {
	return this.dns.Ns(host)
}

func (this *Module) Ptr(host string) (v []string) {
	return this.dns.Ptr(host)
}

func (this *Module) Rr(addr string) (v []string) {
	return this.dns.Rr(addr)
}

func (this *Module) Soa(host string) (v map[string]interface{}) {
	return this.dns.Soa(host)
}

func (this *Module) Spf(host string) (v [][]string) {
	return this.dns.Spf(host)
}

func (this *Module) Srv(host string) (v []map[string]interface{}) {
	return this.dns.Srv(host)
}

func (this *Module) Txt(host string) (v [][]string) {
	return this.dns.Txt(host)
}
