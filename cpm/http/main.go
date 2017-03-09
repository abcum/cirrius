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
)

func init() {
	orbit.Add("http", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
	}
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) New(method string) *Request {
	return NewRequest(this.orb, method)
}

func (this *Module) Get(url string) *stream.ReadCloser {
	cli := new(http.Client)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		this.orb.Quit(err)
	}
	res, err := ctxhttp.Do(this.orb.Context(), cli, req)
	if err != nil {
		this.orb.Quit(err)
	}
	return stream.NewReadCloser(this.orb, res.Body)
}

func (this *Module) Head(url string) *stream.ReadCloser {
	cli := new(http.Client)
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		this.orb.Quit(err)
	}
	res, err := ctxhttp.Do(this.orb.Context(), cli, req)
	if err != nil {
		this.orb.Quit(err)
	}
	return stream.NewReadCloser(this.orb, res.Body)
}

func (this *Module) Put(url, head string, body io.Reader) *stream.ReadCloser {
	cli := new(http.Client)
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		this.orb.Quit(err)
	}
	req.Header.Set("Content-Type", head)
	res, err := ctxhttp.Do(this.orb.Context(), cli, req)
	if err != nil {
		this.orb.Quit(err)
	}
	return stream.NewReadCloser(this.orb, res.Body)
}

func (this *Module) Post(url, head string, body io.Reader) *stream.ReadCloser {
	cli := new(http.Client)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		this.orb.Quit(err)
	}
	req.Header.Set("Content-Type", head)
	res, err := ctxhttp.Do(this.orb.Context(), cli, req)
	if err != nil {
		this.orb.Quit(err)
	}
	return stream.NewReadCloser(this.orb, res.Body)
}

func (this *Module) Patch(url, head string, body io.Reader) *stream.ReadCloser {
	cli := new(http.Client)
	req, err := http.NewRequest("PATCH", url, body)
	if err != nil {
		this.orb.Quit(err)
	}
	req.Header.Set("Content-Type", head)
	res, err := ctxhttp.Do(this.orb.Context(), cli, req)
	if err != nil {
		this.orb.Quit(err)
	}
	return stream.NewReadCloser(this.orb, res.Body)
}

func (this *Module) Delete(url, head string) *stream.ReadCloser {
	cli := new(http.Client)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		this.orb.Quit(err)
	}
	req.Header.Set("Content-Type", head)
	res, err := ctxhttp.Do(this.orb.Context(), cli, req)
	if err != nil {
		this.orb.Quit(err)
	}
	return stream.NewReadCloser(this.orb, res.Body)
}
