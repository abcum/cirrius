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

package blacklists

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/miekg/dns"

	"github.com/abcum/orbit"
)

func init() {
	orbit.Add("check/blacklists", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
		cli: &dns.Client{Timeout: 3 * time.Second},
	}
}

type Module struct {
	orb *orbit.Orbit
	cli *dns.Client
}

func (this *Module) parse(addr string) string {

	ip := net.ParseIP(addr)

	if ip.To4() == nil {
		this.orb.Quit(addrError)
	}

	return ip.String()

}

func (this *Module) reverse(addr string) string {

	parts := strings.Split(this.parse(addr), ".")

	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}

	return strings.Join(parts, ".")

}

func (this *Module) lookup(ip, bl string, chn chan *Check) {

	var q, r *dns.Msg

	q = new(dns.Msg)
	q.RecursionDesired = true
	q.SetQuestion(dns.Fqdn(fmt.Sprintf("%s.%s", ip, bl)), dns.TypeA)

	if r, _, _ = this.cli.Exchange(q, "8.8.8.8:53"); r == nil {
		chn <- &Check{Blacklist: bl, Listed: false, Timeout: true}
	} else {
		chn <- &Check{Blacklist: bl, Listed: len(r.Answer) > 0}
	}

}

func (this *Module) Check(addr string) (checks []*Check) {

	chn := make(chan *Check)

	add := this.reverse(addr)

	cnt := 0

	for _, bl := range ipv4 {
		go this.lookup(add, bl, chn)
	}

	for res := range chn {
		checks = append(checks, res)
		if cnt++; cnt >= len(ipv4) {
			break
		}
	}

	return

}
