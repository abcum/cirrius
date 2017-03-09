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

package spam

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"encoding/json"

	"github.com/abcum/orbit"
	"golang.org/x/net/context/ctxhttp"
)

func init() {
	orbit.Add("check/spam", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
	}
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) Check(value string) *Info {

	var err error
	var bdy []byte
	var ret *Info
	var res *http.Response

	body, _ := json.Marshal(map[string]interface{}{
		"email":   value,
		"options": "short",
	})

	request := new(http.Client)

	content := bytes.NewBuffer(body)

	res, err = ctxhttp.Post(this.orb.Context(), request, endpoint, headtype, content)
	if err != nil {
		return nil
	}

	defer res.Body.Close()

	bdy, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(bdy, &ret)
	if err != nil {
		return nil
	}

	return ret

}
