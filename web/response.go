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

	"github.com/abcum/fibre"
)

type Res struct {
	*fibre.Context
}

func (r *Res) Head(key, val string) {
	r.Context.Response().Header().Set(key, val)
}

func (r *Res) Code(code int) {
	r.Context.Code(code)
}

func (r *Res) Body() io.Writer {
	return r.Context.Response()
}

func (r *Res) Xml(code int, data interface{}) {
	r.Context.XML(code, data)
}

func (r *Res) Text(code int, data interface{}) {
	r.Context.Text(code, data)
}

func (r *Res) Html(code int, data interface{}) {
	r.Context.HTML(code, data)
}

func (r *Res) Json(code int, data interface{}) {
	r.Context.JSON(code, data)
}

func (r *Res) Cbor(code int, data interface{}) {
	r.Context.CBOR(code, data)
}

func (r *Res) Pack(code int, data interface{}) {
	r.Context.PACK(code, data)
}
