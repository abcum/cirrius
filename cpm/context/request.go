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

package context

import (
	"github.com/abcum/fibre"
	"github.com/abcum/orbit"

	"github.com/abcum/cirrius/cpm/stream"
)

type Request struct {
	orb *orbit.Orbit
	fib *fibre.Context

	// Body represents the http request body of the
	// current function request. THe body can be
	// read or streamed as needed.
	Body *stream.Reader `console:"body"`

	// Head is an object containing the http
	// headers of the current request.
	Head map[string]string `console:"head"`

	// Method is the http method used when making
	// this request. It should be one of GET, PUT,
	// POST, PATCH, DELETE, TRACE, OPTIONS, HEAD.
	Method string `console:"method"`

	// User is the basic auth username used for
	// this request. If no username was specified
	// or no basic auth details were used then
	// this will ne an empty string.
	User string `console:"user"`

	// Pass is the basic auth password used for
	// this request. If no password was specified
	// or no basic auth details were used then
	// this will ne an empty string.
	Pass string `console:"pass"`

	// Host is the specific of the web host name
	// used when requesting this function request.
	Host string `console:"host"`

	// Path is the specific path used when running
	// this function. If the request was without a
	// url path, then this will be an empty string.
	Path string `console:"path"`

	// If the request url of the current request
	// contains any query paramaters, then this will
	// be populated with the full http query string.
	Query string `console:"query"`

	// IP is the parsed ipv4 or ipv6 address of the
	// client who has made the http request. It will
	// always be specified, and will never be empty.
	IP string `console:"ip"`
}

func NewRequest(orb *orbit.Orbit) *Request {

	fib := orb.Context().Value("fibre").(*fibre.Context)

	return &Request{
		orb:    orb,
		fib:    fib,
		Body:   stream.NewReader(orb, fib.Request().Body),
		Head:   fib.Head(),
		Method: fib.Request().Method,
		User:   fib.Request().URL().User,
		Pass:   fib.Request().URL().Pass,
		Host:   fib.Request().URL().Host,
		Path:   fib.Request().URL().Path,
		Query:  fib.Request().URL().Query,
		IP:     fib.IP().String(),
	}

}

func (this *Request) Form(key string) *stream.ReadCloser {
	fil, _, err := this.fib.Request().FormFile(key)
	if err != nil {
		this.orb.Quit(err)
	}
	return stream.NewReadCloser(this.orb, fil)
}
