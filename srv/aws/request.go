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
	"encoding/base64"
	"io"
	"net"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type Req struct {
	req *events.APIGatewayProxyRequest
}

func (r *Req) IP() net.IP {
	ip := ""
	if xff, ok := r.req.Headers["X-Forwarded-For"]; !ok {
		ip = r.req.RequestContext.Identity.SourceIP
	} else {
		ip = xff
	}
	ips := strings.Split(ip, ", ")
	ipa := net.ParseIP(ips[0])
	return ipa
}

func (r *Req) Head() map[string]string {
	return r.req.Headers
}

func (r *Req) Body() io.Reader {
	rdr := strings.NewReader(r.req.Body)
	dec := base64.NewDecoder(base64.StdEncoding, rdr)
	return dec
}

func (r *Req) Meth() string {
	return r.req.HTTPMethod
}

func (r *Req) User() string {
	return ""
}

func (r *Req) Pass() string {
	return ""
}

func (r *Req) Host() string {
	if val, ok := r.req.Headers["Host"]; ok {
		return val
	}
	return ""
}

func (r *Req) Path() string {
	if val, err := url.PathUnescape(r.req.Path); err == nil {
		return val
	}
	return r.req.Path
}

func (r *Req) Query() string {
	return ""
}
