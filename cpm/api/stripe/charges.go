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

type Charges struct {
	orb *orbit.Orbit
	api *client.API
}

type Charge struct {
	*stripe.Charge
}

func NewCharges(orb *orbit.Orbit, api *client.API) *Charges {
	return &Charges{
		orb: orb,
		api: api,
	}
}

func (this *Charges) Make(amount float64, currency, user, card string) *Charge {

	params := &stripe.ChargeParams{
		Amount:       stripe.Int64(int64(amount * 100)),
		Currency:     stripe.String(string(stripe.Currency(currency))),
		ReceiptEmail: stripe.String(user),
	}

	params.SetSource(card)

	ret, err := this.api.Charges.New(params)
	if err != nil {
		this.orb.Quit(err)
	}

	return &Charge{ret}

}
