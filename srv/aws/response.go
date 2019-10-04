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

package aws

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"io"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ugorji/go/codec"
)

type Res struct {
	res *events.APIGatewayProxyResponse
	buf *bytes.Buffer
	enc io.WriteCloser
}

func (r *Res) Code(code int) {
	r.res.StatusCode = code
}

func (r *Res) Head(key, val string) {
	if r.res.Headers == nil {
		r.res.Headers = make(map[string]string)
	}
	r.res.Headers[key] = val
}

func (r *Res) Body() io.Writer {
	if r.enc == nil {
		r.buf = bytes.NewBuffer(nil)
		r.enc = base64.NewEncoder(base64.StdEncoding, r.buf)
	}
	return r.enc
}

func (r *Res) Xml(code int, data interface{}) {
	r.Head("Content-Type", "application/xml; charset=utf-8")
	r.Code(code)
	r.Body().Write([]byte(xml.Header))
	if data != nil {
		xml.NewEncoder(r.Body()).Encode(data)
	}
}

func (r *Res) Text(code int, data interface{}) {
	r.Head("Content-Type", "text/plain; charset=utf-8")
	r.Code(code)
	switch conv := data.(type) {
	case []byte:
		r.Body().Write(conv)
	case string:
		r.Body().Write([]byte(conv))
	}
}

func (r *Res) Html(code int, data interface{}) {
	r.Head("Content-Type", "text/html; charset=utf-8")
	r.Code(code)
	switch conv := data.(type) {
	case []byte:
		r.Body().Write(conv)
	case string:
		r.Body().Write([]byte(conv))
	}
}

func (r *Res) Json(code int, data interface{}) {
	r.Head("Content-Type", "application/json; charset=utf-8")
	r.Code(code)
	if data != nil {
		codec.NewEncoder(r.Body(), &jh).Encode(data)
	}
}

func (r *Res) Cbor(code int, data interface{}) {
	r.Head("Content-Type", "application/cbor; charset=utf-8")
	r.Code(code)
	if data != nil {
		codec.NewEncoder(r.Body(), &ch).Encode(data)
	}
}

func (r *Res) Pack(code int, data interface{}) {
	r.Head("Content-Type", "application/msgpack; charset=utf-8")
	r.Code(code)
	if data != nil {
		codec.NewEncoder(r.Body(), &mh).Encode(data)
	}
}

func (r *Res) done() *events.APIGatewayProxyResponse {

	// Specify base64 output

	r.res.IsBase64Encoded = true

	// Ensure that status is set

	if r.res.StatusCode == 0 {
		r.res.StatusCode = 200
	}

	// Encode the body as base64

	if r.enc != nil {
		r.enc.Close()
		r.res.Body = r.buf.String()
	}

	return r.res

}
