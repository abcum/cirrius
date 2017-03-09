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

package domain

import "github.com/abcum/cirrius/cpm/time"

type Info struct {
	Up     bool    `console:"up" codec:"up"`
	Ok     bool    `console:"ok" codec:"ok"`
	Host   string  `console:"host" codec:"host"`
	Port   string  `console:"port" codec:"port"`
	Scheme string  `console:"scheme" codec:"scheme"`
	Domain string  `console:"domain" codec:"domain"`
	Certs  []*Cert `console:"certs" codec:"certs"`
}

type Cert struct {
	CA         bool       `console:"ca" codec:"ca"`
	IPs        []string   `console:"ips" codec:"ips"`
	Names      []string   `console:"names" codec:"names"`
	Emails     []string   `console:"emails" codec:"emails"`
	Issuer     string     `console:"issuer" codec:"issuer"`
	Subject    string     `console:"subject" codec:"subject"`
	Version    int        `console:"version" codec:"version"`
	NotAfter   *time.Time `console:"notafter" codec:"notafter"`
	NotBefore  *time.Time `console:"notbefore" codec:"notbefore"`
	Algorithm  string     `console:"algorithm" codec:"algorithm"`
	Expired    bool       `console:"expired" codec:"expired"`
	Deprecated bool       `console:"deprecated" codec:"deprecated"`
}
