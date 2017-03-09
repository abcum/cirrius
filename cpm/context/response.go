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

type Response struct {
	orb *orbit.Orbit
	fib *fibre.Context

	// Body represents the http response body of the
	// current function request. The body can be
	// can be written or streamed to.
	Body *stream.Writer `console:"body"`
}

func NewResponse(orb *orbit.Orbit) *Response {

	fib := orb.Context().Value("fibre").(*fibre.Context)

	return &Response{
		orb:  orb,
		fib:  fib,
		Body: stream.NewWriter(orb, fib.Response()),
	}

}

func (this *Response) Head(key, val string) {
	this.fib.Response().Header().Set(key, val)
}
