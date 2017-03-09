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

package mailspike

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/abcum/orbit"

	"github.com/abcum/cirrius/cpm/dns"
)

func init() {
	orbit.Add("check/mailspike", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
	}
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) label(num int) string {
	switch num {
	default:
		return "-"
	case 10:
		return "Worst possible reputation"
	case 11:
		return "Very bad reputation"
	case 12:
		return "Bad reputation"
	case 13:
		return "Suspicious behaviour"
	case 14:
		return "Neutral - probably spam"
	case 15:
		return "Neutral"
	case 16:
		return "Neutral - probably legit"
	case 17:
		return "Possible legit sender"
	case 18:
		return "Good reputation"
	case 19:
		return "Very good reputation"
	case 20:
		return "Excellent reputation"
	}
}

func (this *Module) parse(addr string) string {

	ip := net.ParseIP(addr)

	if ip.To4() == nil {
		this.orb.Quit(addrError)
	}

	return ip.String()

}

func (this *Module) reverse(addr string) string {

	parts := strings.Split(this.parse(addr), ".")

	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}

	return strings.Join(parts, ".")

}

func (this *Module) Check(addr string) (score *Score) {

	add := this.reverse(addr)

	req := fmt.Sprintf("%s.query.mailspike.net", add)

	res := dns.New(this.orb).(*dns.Module).A(req)

	if len(res) > 0 {
		ip := this.parse(res[0])
		pc := strings.Split(ip, ".")[3]
		sc, _ := strconv.Atoi(pc)
		return &Score{Score: sc, Label: this.label(sc)}
	}

	return &Score{Score: 0, Label: "-"}

}
