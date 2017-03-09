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

package dns

import (
	"fmt"
	"net"

	"github.com/miekg/dns"

	"github.com/abcum/orbit"
)

type Client struct {
	orb *orbit.Orbit
	Val *dns.Client `structs:"-"`
}

func NewClient(orb *orbit.Orbit, val *dns.Client) *Client {
	return &Client{
		orb: orb,
		Val: val,
	}
}

func (this *Client) A(host string) (v []string) {
	for _, i := range this.run(host, dns.TypeA) {
		v = append(v, i.(*dns.A).A.String())
	}
	return
}

func (this *Client) Aaaa(host string) (v []string) {
	for _, i := range this.run(host, dns.TypeAAAA) {
		v = append(v, i.(*dns.AAAA).AAAA.String())
	}
	return
}

func (this *Client) Cname(host string) (v []string) {
	for _, i := range this.run(host, dns.TypeCNAME) {
		v = append(v, i.(*dns.CNAME).Target)
	}
	return
}

func (this *Client) Mx(host string) (v []map[string]interface{}) {
	for _, i := range this.run(host, dns.TypeMX) {
		v = append(v, map[string]interface{}{
			"priority": i.(*dns.MX).Preference,
			"exchange": i.(*dns.MX).Mx,
		})
	}
	return
}

func (this *Client) Naptr(host string) (v []map[string]interface{}) {
	for _, i := range this.run(host, dns.TypeNAPTR) {
		v = append(v, map[string]interface{}{
			"flags":       i.(*dns.NAPTR).Flags,
			"service":     i.(*dns.NAPTR).Service,
			"regexp":      i.(*dns.NAPTR).Regexp,
			"replacement": i.(*dns.NAPTR).Replacement,
			"order":       i.(*dns.NAPTR).Order,
			"preference":  i.(*dns.NAPTR).Preference,
		})
	}
	return
}

func (this *Client) Ns(host string) (v []string) {
	for _, i := range this.run(host, dns.TypeNS) {
		v = append(v, i.(*dns.NS).Ns)
	}
	return
}

func (this *Client) Ptr(host string) (v []string) {
	for _, i := range this.run(host, dns.TypePTR) {
		v = append(v, i.(*dns.PTR).Ptr)
	}
	return
}

func (this *Client) Rr(addr string) (v []string) {
	v, err := net.LookupAddr(addr)
	if err != nil {
		this.orb.Quit(err)
	}
	return v
}

func (this *Client) Soa(host string) (v map[string]interface{}) {
	for _, i := range this.run(host, dns.TypeSOA) {
		v = map[string]interface{}{
			"nsname":     i.(*dns.SOA).Ns,
			"hostmaster": i.(*dns.SOA).Mbox,
			"serial":     i.(*dns.SOA).Serial,
			"refresh":    i.(*dns.SOA).Refresh,
			"retry":      i.(*dns.SOA).Retry,
			"expire":     i.(*dns.SOA).Expire,
			"minttl":     i.(*dns.SOA).Minttl,
		}
		break
	}
	return
}

func (this *Client) Spf(host string) (v [][]string) {
	for _, i := range this.run(host, dns.TypeSPF) {
		var m []string
		for _, s := range i.(*dns.SPF).Txt {
			m = append(m, s)
		}
		v = append(v, m)
	}
	return
}

func (this *Client) Srv(host string) (v []map[string]interface{}) {
	for _, i := range this.run(host, dns.TypeSRV) {
		v = append(v, map[string]interface{}{
			"priority": i.(*dns.SRV).Priority,
			"weight":   i.(*dns.SRV).Weight,
			"port":     i.(*dns.SRV).Port,
			"name":     i.(*dns.SRV).Target,
		})
	}
	return
}

func (this *Client) Txt(host string) (v [][]string) {
	for _, i := range this.run(host, dns.TypeTXT) {
		var m []string
		for _, s := range i.(*dns.TXT).Txt {
			m = append(m, s)
		}
		v = append(v, m)
	}
	return
}

func (this *Client) run(name string, kind uint16) (res []dns.RR) {

	var e error
	var q *dns.Msg
	var x *dns.Msg

	q = new(dns.Msg)
	q.RecursionDesired = true
	q.SetQuestion(dns.Fqdn(name), kind)

	if x, _, e = this.Val.Exchange(q, "8.8.8.8:53"); x == nil {
		this.orb.Quit(fmt.Errorf("Unable to query DNS server. %s", e))
	}

	if x.Rcode != dns.RcodeSuccess {
		this.orb.Quit(fmt.Errorf("Invalid answer from DNS server for %s.", name))
	}

	return x.Answer

}
