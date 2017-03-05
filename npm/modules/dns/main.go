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
	"github.com/robertkrimen/otto"
)

var client *dns.Client

type task struct {
	cb func()
}

func (t *task) Startup(orb *orbit.Orbit) {

}

func (t *task) Cleanup(orb *orbit.Orbit) {
}

func (t *task) Execute(orb *orbit.Orbit) (err error) {
	t.cb()
	return
}

type signature func(string, callback2) otto.Value
type callback2 func(otto.Value, otto.Value)
type callback3 func(otto.Value, otto.Value, otto.Value)

func process(name string, kind uint16) (res []dns.RR, err error) {

	var m *dns.Msg
	var x *dns.Msg

	m = new(dns.Msg)
	m.RecursionDesired = true
	m.SetQuestion(dns.Fqdn(name), kind)

	if x, _, _ = client.Exchange(m, "8.8.8.8:53"); x == nil {
		return nil, fmt.Errorf("Unable to query DNS server.")
	}

	if x.Rcode != dns.RcodeSuccess {
		return nil, fmt.Errorf("Invalid answer from DNS server.")
	}

	return x.Answer, nil

}

func init() {

	client = new(dns.Client)

	orbit.Add("dns", func(orb *orbit.Orbit) (otto.Value, error) {

		dns := map[string]interface{}{

			"getServers": func() (val otto.Value) {
				val, _ = otto.ToValue([]string{"8.8.8.8", "8.8.4.4"})
				return
			},

			"setServers": func(servers []string) (val otto.Value) {
				return
			},

			"lookupService": func(addr string, port int, cb callback3) (val otto.Value) {

				query := &task{}
				orb.Push(query)

				go func() {
					var e, r, s otto.Value
					defer orb.Pull(query)
					defer orb.Next(&task{func() { cb(e, r, s) }})
					if res, err := net.LookupAddr(addr); err != nil {
						e, _ = orb.ToValue(orb.MakeCustomError("Error", err.Error()))
					} else {
						switch len(res) {
						case 0:
							r, _ = orb.ToValue(addr)
							s, _ = orb.ToValue(port)
						default:
							r, _ = orb.ToValue(res[0])
							s, _ = orb.ToValue(port)
						}
						if serv, ok := ports[port]; ok {
							s, _ = orb.ToValue(serv)
						}
					}
				}()

				return

			},

			"resolve4": func(name string, cb callback2) (val otto.Value) {

				query := &task{}
				orb.Push(query)

				go func() {
					var e, r otto.Value
					var a []string
					defer orb.Pull(query)
					defer orb.Next(&task{func() { cb(e, r) }})
					if res, err := process(name, dns.TypeA); err != nil {
						e, _ = orb.ToValue(orb.MakeCustomError("Error", err.Error()))
					} else {
						for _, i := range res {
							a = append(a, i.(*dns.A).A.String())
						}
						r, _ = orb.ToValue(a)
					}
				}()

				return

			},

			"resolve6": func(name string, cb callback2) (val otto.Value) {

				query := &task{}
				orb.Push(query)

				go func() {
					var e, r otto.Value
					var a []string
					defer orb.Pull(query)
					defer orb.Next(&task{func() { cb(e, r) }})
					if res, err := process(name, dns.TypeAAAA); err != nil {
						e, _ = orb.ToValue(orb.MakeCustomError("Error", err.Error()))
					} else {
						for _, i := range res {
							a = append(a, i.(*dns.AAAA).AAAA.String())
						}
						r, _ = orb.ToValue(a)
					}
				}()

				return

			},

			"resolveCname": func(name string, cb callback2) (val otto.Value) {

				query := &task{}
				orb.Push(query)

				go func() {
					var e, r otto.Value
					var a []string
					defer orb.Pull(query)
					defer orb.Next(&task{func() { cb(e, r) }})
					if res, err := process(name, dns.TypeCNAME); err != nil {
						e, _ = orb.ToValue(orb.MakeCustomError("Error", err.Error()))
					} else {
						for _, i := range res {
							a = append(a, i.(*dns.CNAME).Target)
						}
						r, _ = orb.ToValue(a)
					}
				}()

				return

			},

			"resolveMx": func(name string, cb callback2) (val otto.Value) {

				query := &task{}
				orb.Push(query)

				go func() {
					var e, r otto.Value
					var a []map[string]interface{}
					defer orb.Pull(query)
					defer orb.Next(&task{func() { cb(e, r) }})
					if res, err := process(name, dns.TypeMX); err != nil {
						e, _ = orb.ToValue(orb.MakeCustomError("Error", err.Error()))
					} else {
						for _, i := range res {
							a = append(a, map[string]interface{}{
								"priority": i.(*dns.MX).Preference,
								"exchange": i.(*dns.MX).Mx,
							})
						}
						r, _ = orb.ToValue(a)
					}
				}()

				return

			},

			"resolveNaptr": func(name string, cb callback2) (val otto.Value) {

				query := &task{}
				orb.Push(query)

				go func() {
					var e, r otto.Value
					var a []map[string]interface{}
					defer orb.Pull(query)
					defer orb.Next(&task{func() { cb(e, r) }})
					if res, err := process(name, dns.TypeNAPTR); err != nil {
						e, _ = orb.ToValue(orb.MakeCustomError("Error", err.Error()))
					} else {
						for _, i := range res {
							a = append(a, map[string]interface{}{
								"flags":       i.(*dns.NAPTR).Flags,
								"service":     i.(*dns.NAPTR).Service,
								"regexp":      i.(*dns.NAPTR).Regexp,
								"replacement": i.(*dns.NAPTR).Replacement,
								"order":       i.(*dns.NAPTR).Order,
								"preference":  i.(*dns.NAPTR).Preference,
							})
						}
						r, _ = orb.ToValue(a)
					}
				}()

				return

			},

			"resolveNs": func(name string, cb callback2) (val otto.Value) {

				query := &task{}
				orb.Push(query)

				go func() {
					var e, r otto.Value
					var a []string
					defer orb.Pull(query)
					defer orb.Next(&task{func() { cb(e, r) }})
					if res, err := process(name, dns.TypeNS); err != nil {
						e, _ = orb.ToValue(orb.MakeCustomError("Error", err.Error()))
					} else {
						for _, i := range res {
							a = append(a, i.(*dns.NS).Ns)
						}
						r, _ = orb.ToValue(a)
					}
				}()

				return

			},

			"resolveSoa": func(name string, cb callback2) (val otto.Value) {

				query := &task{}
				orb.Push(query)

				go func() {
					var e, r otto.Value
					var a map[string]interface{}
					defer orb.Pull(query)
					defer orb.Next(&task{func() { cb(e, r) }})
					if res, err := process(name, dns.TypeSOA); err != nil {
						e, _ = orb.ToValue(orb.MakeCustomError("Error", err.Error()))
					} else {
						for _, i := range res {
							a = map[string]interface{}{
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
						r, _ = orb.ToValue(a)
					}
				}()

				return

			},

			"resolveSrv": func(name string, cb callback2) (val otto.Value) {

				query := &task{}
				orb.Push(query)

				go func() {
					var e, r otto.Value
					var a []map[string]interface{}
					defer orb.Pull(query)
					defer orb.Next(&task{func() { cb(e, r) }})
					if res, err := process(name, dns.TypeSRV); err != nil {
						e, _ = orb.ToValue(orb.MakeCustomError("Error", err.Error()))
					} else {
						for _, i := range res {
							a = append(a, map[string]interface{}{
								"priority": i.(*dns.SRV).Priority,
								"weight":   i.(*dns.SRV).Weight,
								"port":     i.(*dns.SRV).Port,
								"name":     i.(*dns.SRV).Target,
							})
						}
						r, _ = orb.ToValue(a)
					}
				}()

				return

			},

			"resolvePtr": func(name string, cb callback2) (val otto.Value) {

				query := &task{}
				orb.Push(query)

				go func() {
					var e, r otto.Value
					var a []string
					defer orb.Pull(query)
					defer orb.Next(&task{func() { cb(e, r) }})
					if res, err := process(name, dns.TypePTR); err != nil {
						e, _ = orb.ToValue(orb.MakeCustomError("Error", err.Error()))
					} else {
						for _, i := range res {
							a = append(a, i.(*dns.PTR).Ptr)
						}
						r, _ = orb.ToValue(a)
					}
				}()

				return

			},

			"resolveTxt": func(name string, cb callback2) (val otto.Value) {

				query := &task{}
				orb.Push(query)

				go func() {
					var e, r otto.Value
					var a [][]string
					defer orb.Pull(query)
					defer orb.Next(&task{func() { cb(e, r) }})
					if res, err := process(name, dns.TypeTXT); err != nil {
						e, _ = orb.ToValue(orb.MakeCustomError("Error", err.Error()))
					} else {
						for _, i := range res {
							var m []string
							for _, s := range i.(*dns.TXT).Txt {
								m = append(m, s)
							}
							a = append(a, m)
						}
						r, _ = orb.ToValue(a)
					}
				}()

				return

			},

			"resolveSpf": func(name string, cb callback2) (val otto.Value) {

				query := &task{}
				orb.Push(query)

				go func() {
					var e, r otto.Value
					var a [][]string
					defer orb.Pull(query)
					defer orb.Next(&task{func() { cb(e, r) }})
					if res, err := process(name, dns.TypeSPF); err != nil {
						e, _ = orb.ToValue(orb.MakeCustomError("Error", err.Error()))
					} else {
						for _, i := range res {
							var m []string
							for _, s := range i.(*dns.SPF).Txt {
								m = append(m, s)
							}
							a = append(a, m)
						}
						r, _ = orb.ToValue(a)
					}
				}()

				return

			},

			"reverse": func(addr string, cb callback2) (val otto.Value) {

				query := &task{}
				orb.Push(query)

				go func() {
					var e, r otto.Value
					defer orb.Pull(query)
					defer orb.Next(&task{func() { cb(e, r) }})
					if res, err := net.LookupAddr(addr); err != nil {
						e, _ = orb.ToValue(orb.MakeCustomError("Error", err.Error()))
					} else {
						r, _ = orb.ToValue(res)
					}
				}()

				return

			},
		}

		dns["lookup"] = func(name string, kind string, cb callback2) (val otto.Value) {
			switch kind {
			case "A":
				return dns["resolve4"].(signature)(name, cb)
			case "AAAA":
				return dns["resolve6"].(signature)(name, cb)
			case "MX":
				return dns["resolveMx"].(signature)(name, cb)
			case "TXT":
				return dns["resolveTxt"].(signature)(name, cb)
			case "SRV":
				return dns["resolveSrv"].(signature)(name, cb)
			case "PTR":
				return dns["resolvePtr"].(signature)(name, cb)
			case "NS":
				return dns["resolveNs"].(signature)(name, cb)
			case "CNAME":
				return dns["resolveCname"].(signature)(name, cb)
			case "SOA":
				return dns["resolveSoa"].(signature)(name, cb)
			case "NAPTR":
				return dns["resolveNaptr"].(signature)(name, cb)
			case "SPF":
				return dns["resolveSpf"].(signature)(name, cb)
			}
			return otto.UndefinedValue()
		}

		return orb.ToValue(dns)

	})

}
