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

package senderbase

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/abcum/orbit"

	"github.com/abcum/cirrius/cpm/dns"
)

func init() {
	orbit.Add("check/senderbase", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
	}
}

type Module struct {
	orb *orbit.Orbit
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

	req := fmt.Sprintf("%s.query.senderbase.org", add)

	res := dns.New(this.orb).(*dns.Module).Txt(req)

	if len(res) > 0 {
		return this.info(strings.Join(res[0], " "))
	}

	return &Score{Version: 0}

}

func (this *Module) info(txt string) (score *Score) {

	score = new(Score)

	parts := strings.Split(txt, "|")

	for _, part := range parts {

		kv := strings.Split(part, "=")

		switch kv[0] {
		case "0":
			score.Version = this.infoInt(kv[1])
		case "1":
			score.Organisation.Name = this.infoStr(kv[1])
		case "2":
			score.Statistics.Organisation.Daily = this.infoDec(kv[1])
		case "3":
			score.Statistics.Organisation.Monthly = this.infoDec(kv[1])
		case "4":
			score.Organisation.ID = this.infoStr(kv[1])
		case "5":
			score.Organisation.Category = this.infoStr(kv[1])
		case "6":
			score.Statistics.Organisation.First = this.infoInt(kv[1])
		case "7":
			score.Organisation.TotalDomains = this.infoInt(kv[1])
		case "8":
			score.Organisation.TotalIpsHave = this.infoInt(kv[1])
		case "9":
			score.Organisation.TotalIpsUsed = this.infoInt(kv[1])
		case "10":
			score.Organisation.Fortune1000 = this.infoVal(kv[1])
		case "20":
			score.Hostname = this.infoStr(kv[1])
		case "21":
			score.Domain = this.infoStr(kv[1])
		case "22":
			score.Matches = this.infoVal(kv[1])
		case "23":
			score.Statistics.Domain.Daily = this.infoDec(kv[1])
		case "24":
			score.Statistics.Domain.Monthly = this.infoDec(kv[1])
		case "25":
			score.Statistics.Domain.First = this.infoInt(kv[1])
		case "26":
			score.Rating = this.infoDec(kv[1])
		case "40":
			score.Statistics.IP.Daily = this.infoDec(kv[1])
		case "41":
			score.Statistics.IP.Monthly = this.infoDec(kv[1])
		case "43":
			score.Statistics.IP.Average = this.infoDec(kv[1])
		case "44":
			score.Statistics.IP.Percent = this.infoDec(kv[1])
		case "47":
			score.Statistics.IP.Blacklist = this.infoDec(kv[1])
		case "50":
			score.Geolocation.City = this.infoStr(kv[1])
		case "51":
			score.Geolocation.State = this.infoStr(kv[1])
		case "52":
			score.Geolocation.Postcode = this.infoStr(kv[1])
		case "53":
			score.Geolocation.Country = this.infoStr(kv[1])
		case "54":
			score.Geolocation.Longitude = this.infoDec(kv[1])
		case "55":
			score.Geolocation.Latitude = this.infoDec(kv[1])
		}

	}

	return

}

func (this *Module) infoStr(txt string) string {
	return txt
}

func (this *Module) infoInt(txt string) int64 {
	out, _ := strconv.ParseInt(txt, 10, 64)
	return out
}

func (this *Module) infoDec(txt string) float64 {
	out, _ := strconv.ParseFloat(txt, 64)
	return out
}

func (this *Module) infoVal(txt string) bool {
	if txt == "Y" {
		return true
	}
	return false
}
