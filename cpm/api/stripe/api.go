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

package stripe

import (
	"github.com/abcum/orbit"

	"github.com/stripe/stripe-go/client"
)

type API struct {
	orb *orbit.Orbit
	val *client.API
}

func NewAPI(orb *orbit.Orbit, key string) *API {
	return &API{
		orb: orb,
		val: client.New(key, nil),
	}
}

func (this *API) Cards() *Cards {
	return NewCards(this.orb, this.val)
}

func (this *API) Charges() *Charges {
	return NewCharges(this.orb, this.val)
}

func (this *API) Customers() *Customers {
	return NewCustomers(this.orb, this.val)
}
