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

package getaddress

import (
	"fmt"
	"net/http"
	"time"

	"encoding/json"

	"github.com/abcum/orbit"
	"golang.org/x/net/context/ctxhttp"
)

type API struct {
	orb *orbit.Orbit
	key string
}

func NewAPI(orb *orbit.Orbit, key string) *API {
	return &API{
		orb: orb,
		key: key,
	}
}

func (this *API) Postcode(code string) (out *Info) {

	cli := &http.Client{Timeout: 3 * time.Second}

	end := "https://api.getaddress.io/v2/uk/%s?format=true&api-key=%s"

	url := fmt.Sprintf(end, code, this.key)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		this.orb.Quit(err)
	}

	res, err := ctxhttp.Do(this.orb.Context(), cli, req)
	if err != nil {
		this.orb.Quit(err)
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		this.orb.Quit(err)
	}

	return

}

func (this *API) House(code, house string) (out *Info) {

	cli := &http.Client{Timeout: 3 * time.Second}

	end := "https://api.getAddress.io/v2/uk/%s/%s?format=true&api-key=%s"

	url := fmt.Sprintf(end, code, house, this.key)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		this.orb.Quit(err)
	}

	res, err := ctxhttp.Do(this.orb.Context(), cli, req)
	if err != nil {
		this.orb.Quit(err)
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		this.orb.Quit(err)
	}

	return

}
