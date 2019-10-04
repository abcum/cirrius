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

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
)

type Customers struct {
	orb *orbit.Orbit
	api *client.API
}

type Customer struct {
	*stripe.Customer
}

func NewCustomers(orb *orbit.Orbit, api *client.API) *Customers {
	return &Customers{
		orb: orb,
		api: api,
	}
}

func (this *Customers) Create(email string) *Customer {

	params := &stripe.CustomerParams{
		Email: stripe.String(email),
	}

	ret, err := this.api.Customers.New(params)
	if err != nil {
		this.orb.Quit(err)
	}

	return &Customer{ret}

}

func (this *Customers) Update(id string, email string) *Customer {

	params := &stripe.CustomerParams{
		Email: stripe.String(email),
	}

	ret, err := this.api.Customers.Update(id, params)
	if err != nil {
		this.orb.Quit(err)
	}

	return &Customer{ret}

}

func (this *Customers) Delete(id string) {

	params := &stripe.CustomerParams{}

	_, err := this.api.Customers.Del(id, params)
	if err != nil {
		this.orb.Quit(err)
	}

}
