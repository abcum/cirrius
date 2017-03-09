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

package http

import (
	"io"
	"net/http"

	"github.com/abcum/orbit"
	"golang.org/x/net/context/ctxhttp"

	"github.com/abcum/cirrius/cpm/stream"
	"github.com/abcum/cirrius/cpm/time"
)

type Request struct {
	orb *orbit.Orbit
	met string
	dur string
	url string
	usr string
	pss string
	cli *http.Client
	req io.Reader
	hdr map[string]string
}

func NewRequest(orb *orbit.Orbit, met string) *Request {
	return &Request{
		orb: orb,
		met: met,
		cli: new(http.Client),
		hdr: make(map[string]string),
	}
}

func (this *Request) Auth(user, pass string) *Request {
	this.usr, this.pss = user, pass
	return this
}

func (this *Request) Body(req io.Reader) *Request {
	this.req = req
	return this
}

func (this *Request) Method(met string) *Request {
	this.met = met
	return this
}

func (this *Request) Head(key, val string) *Request {
	this.hdr[key] = val
	return this
}

func (this *Request) Timeout(dur *time.Duration) *Request {
	this.cli.Timeout = dur.Val
	return this
}

func (this *Request) URL(url string) *Request {
	this.url = url
	return this
}

func (this *Request) Do() *stream.ReadCloser {
	req, err := http.NewRequest(this.met, this.url, this.req)
	if err != nil {
		this.orb.Quit(err)
	}
	if this.usr != "" || this.pss != "" {
		req.SetBasicAuth(this.usr, this.pss)
	}
	for k, v := range this.hdr {
		req.Header.Set(k, v)
	}
	res, err := ctxhttp.Do(this.orb.Context(), this.cli, req)
	if err != nil {
		this.orb.Quit(err)
	}
	return stream.NewReadCloser(this.orb, res.Body)
}
