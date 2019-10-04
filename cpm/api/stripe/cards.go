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

type Cards struct {
	orb *orbit.Orbit
	api *client.API
}

type Card struct {
	*stripe.Card
}

func NewCards(orb *orbit.Orbit, api *client.API) *Cards {
	return &Cards{
		orb: orb,
		api: api,
	}
}

func (this *Cards) Create(customer, token string) *Card {

	params := &stripe.CardParams{
		Customer: stripe.String(customer),
		Token:    stripe.String(token),
	}

	ret, err := this.api.Cards.New(params)
	if err != nil {
		this.orb.Quit(err)
	}

	return &Card{ret}

}

func (this *Cards) Update(id string, name string) *Card {

	params := &stripe.CardParams{
		Name: stripe.String(name),
	}

	ret, err := this.api.Cards.Update(id, params)
	if err != nil {
		this.orb.Quit(err)
	}

	return &Card{ret}

}

func (this *Cards) Delete(customer, id string) {

	params := &stripe.CardParams{
		Customer: stripe.String(customer),
	}

	_, err := this.api.Cards.Del(id, params)
	if err != nil {
		this.orb.Quit(err)
	}

}
