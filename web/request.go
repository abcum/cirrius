// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"io"
	"net"

	"github.com/abcum/fibre"
)

type Req struct {
	*fibre.Context
}

func (r *Req) IP() net.IP {
	return r.Context.IP()
}

func (r *Req) Head() map[string]string {
	return r.Context.Head()
}

func (r *Req) Body() io.Reader {
	return r.Context.Request().Body
}

func (r *Req) Meth() string {
	return r.Context.Request().Method
}

func (r *Req) User() string {
	return r.Context.Request().URL().User
}

func (r *Req) Pass() string {
	return r.Context.Request().URL().Pass
}

func (r *Req) Host() string {
	return r.Context.Request().URL().Host
}

func (r *Req) Path() string {
	return r.Context.Request().URL().Path
}

func (r *Req) Query() string {
	return r.Context.Request().URL().Query
}

func (r *Req) Fibre() *fibre.Context {
	return r.Context
}
