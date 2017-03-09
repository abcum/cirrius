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

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/abcum/orbit"

	timer "github.com/abcum/cirrius/cpm/time"
)

func init() {
	orbit.Add("check/domain", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
	}
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) Check(addr string) (info *Info) {

	info = this.parse(addr)

	switch info.Scheme {
	case "http":
		return this.checkHttp(info)
	case "https":
		return this.checkHttps(info)
	}

	return

}

func (this *Module) parse(addr string) *Info {

	var host, port string

	full, err := url.Parse(addr)
	if err != nil {
		this.orb.Quit(err)
	}

	if host = full.Hostname(); host == "" {
		this.orb.Quit(addrError)
	}

	switch full.Scheme {
	case "http":
		port = "80"
	case "https":
		port = "443"
	default:
		this.orb.Quit(addrError)
	}

	return &Info{
		Host:   host,
		Port:   port,
		Scheme: full.Scheme,
		Domain: fmt.Sprintf("%s:%s", host, port),
	}

}

func (this *Module) checkHttp(info *Info) *Info {

	var err error
	var con net.Conn

	con, err = net.Dial("tcp", info.Domain)
	if err != nil {
		return info
	}

	defer con.Close()

	info.Up, info.Ok = true, true

	return info

}

func (this *Module) checkHttps(info *Info) *Info {

	var err error
	var con *tls.Conn

	con, err = tls.Dial("tcp", info.Domain, config)
	if err != nil {
		return info
	}

	defer con.Close()

	info.Up, info.Ok = true, true

	now := time.Now()

	done := make(map[string]struct{})

	peer := con.ConnectionState().PeerCertificates

	opts := x509.VerifyOptions{
		CurrentTime:   now,
		Roots:         config.RootCAs,
		DNSName:       info.Host,
		Intermediates: x509.NewCertPool(),
	}

	for i, cert := range peer {
		if i > 0 {
			opts.Intermediates.AddCert(cert)
		}
	}

	if _, err = peer[0].Verify(opts); err != nil {
		info.Ok = false
	}

	if err = peer[0].VerifyHostname(info.Host); err != nil {
		info.Ok = false
	}

	for _, c := range peer {

		if _, ok := done[string(c.Signature)]; ok {
			continue
		}

		done[string(c.Signature)] = struct{}{}

		cert := &Cert{
			CA:        c.IsCA,
			Names:     c.DNSNames,
			Emails:    c.EmailAddresses,
			Issuer:    c.Issuer.CommonName,
			Subject:   c.Subject.CommonName,
			Version:   c.Version,
			NotAfter:  timer.NewTime(this.orb, c.NotAfter),
			NotBefore: timer.NewTime(this.orb, c.NotBefore),
			Algorithm: c.SignatureAlgorithm.String(),
		}

		if now.Before(c.NotBefore) || now.After(c.NotAfter) {
			cert.Expired = true
		}

		if _, ok := deprecateds[c.SignatureAlgorithm]; ok {
			cert.Deprecated = true
		}

		info.Certs = append(info.Certs, cert)

	}

	return info

}
