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

package card

import (
	"strconv"
	"strings"

	"github.com/abcum/orbit"
)

func init() {
	orbit.Add("check/card", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
	}
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) clean(number string) string {
	number = strings.Replace(number, "-", "", -1)
	number = strings.Replace(number, " ", "", -1)
	return number
}

func (this *Module) reverse(number string) string {
	digits := []rune(number)
	for i, j := 0, len(number)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return string(digits)
}

func (this *Module) length(number string, lengths ...int) bool {
	for _, length := range lengths {
		if len(number) == length {
			return true
		}
	}
	return false
}

func (this *Module) prefixed(number string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(number, prefix) {
			return true
		}
	}
	return false
}

func (this *Module) Type(number string) string {

	number = this.clean(number)

	if len(number) == 0 {
		return "Unknown"
	}

	if this.length(number, 16) && this.prefixed(number, "3") {
		return "JCB"
	}

	if this.length(number, 13, 16) && this.prefixed(number, "4") {
		return "Visa"
	}

	if this.length(number, 15) && this.prefixed(number, "34", "37") {
		return "American Express"
	}

	if this.length(number, 16) && this.prefixed(number, "6011") {
		return "Discover"
	}

	if this.length(number, 16) && this.prefixed(number, "2014", "2149") {
		return "enRoute"
	}

	if this.length(number, 15) && this.prefixed(number, "1800", "2131") {
		return "JCB"
	}

	if this.length(number, 16) && this.prefixed(number, "51", "52", "53", "54", "55") {
		return "Mastercard"
	}

	if this.length(number, 14) && this.prefixed(number, "36", "38", "300", "301", "302", "303", "304", "305") {
		return "Diners Club / Carteblanche"
	}

	if this.length(number, 16) && this.prefixed(number, "2131") {
		return "JCB"
	}

	return "Unknown"

}

func (this *Module) Check(number string) bool {

	number = this.clean(number)

	number = this.reverse(number)

	hold, even := 0, false

	for _, digit := range number {

		i, err := strconv.Atoi(string(digit))
		if err != nil {
			this.orb.Quit(err)
		}

		if even {
			i *= 2
		}

		if i > 9 {
			hold -= 9
		}

		hold += i

		even = !even

	}

	if hold%10 == 0 {
		return true
	}

	return false

}
