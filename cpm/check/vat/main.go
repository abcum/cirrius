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

package vat

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	"encoding/xml"

	"github.com/abcum/orbit"
	"golang.org/x/net/context/ctxhttp"
)

func init() {
	orbit.Add("check/vat", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
	}
}

type Module struct {
	orb *orbit.Orbit
}

// Check validates the format and existence of a
// European VAT number, for all European member states.
//
//	var vat = require('validate/vat');
// 	vat.Check("GB 982 1503 23");
//
func (this *Module) Check(number string) bool {

	if ok := this.checkFormat(number); !ok {
		return false
	}

	if ok := this.checkNumber(number); !ok {
		return false
	}

	return true

}

func (this *Module) clean(value string) string {
	value = strings.ToUpper(value)
	value = strings.Replace(value, "-", "", -1)
	value = strings.Replace(value, " ", "", -1)
	return value
}

func (this *Module) checkFormat(value string) bool {

	if len(value) < 3 {
		return false
	}

	value = this.clean(value)

	if regex, ok := patterns[value[0:2]]; ok {
		return regex.MatchString(value)
	}

	return false

}

func (this *Module) checkNumber(value string) bool {

	var err error
	var bdy []byte
	var ret *response
	var res *http.Response

	if len(value) < 3 {
		return false
	}

	value = this.clean(value)

	body := envelope
	body = strings.Replace(body, "{{.country}}", value[:2], 1)
	body = strings.Replace(body, "{{.vnumber}}", value[2:], 1)

	request := new(http.Client)

	content := bytes.NewBufferString(body)

	res, err = ctxhttp.Post(this.orb.Context(), request, endpoint, headtype, content)
	if err != nil {
		return false
	}

	defer res.Body.Close()

	bdy, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return false
	}

	if bytes.Contains(bdy, []byte("INVALID_INPUT")) {
		return false
	}

	err = xml.Unmarshal(bdy, &ret)
	if err != nil {
		return false
	}

	return ret.Soap.Soap.Valid

}
