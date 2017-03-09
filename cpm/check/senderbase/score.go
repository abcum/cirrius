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

var keys = map[string]string{
	"0":  "version_number",
	"1":  "org_name",
	"2":  "org_daily_magnitude",
	"3":  "org_monthly_magnitude",
	"4":  "org_id",
	"5":  "org_category",
	"6":  "org_first_message",
	"7":  "org_domains_count",
	"8":  "org_ip_controlled_count",
	"9":  "org_ip_used_count",
	"10": "org_fortune_1000",

	"20": "hostname",
	"21": "domain_name",
	"22": "hostname_matches_ip",
	"23": "domain_daily_magnitude",
	"24": "domain_monthly_magnitude",
	"25": "domain_first_message",
	"26": "domain_rating",

	"40": "ip_daily_magnitude",
	"41": "ip_monthly_magnitude",
	"43": "ip_average_magnitude",
	"44": "ip_30_day_volume_percent",
	"45": "ip_in_bonded_sender",
	"46": "ip_cidr_range",
	"47": "ip_blacklist_score",

	"50": "ip_city",
	"51": "ip_state",
	"52": "ip_postal_code",
	"53": "ip_country",
	"54": "ip_longitude",
	"55": "ip_latitude",
}

type Score struct {
	Version int64 `console:"version" codec:"version"`

	Rating   float64 `console:"rating" codec:"rating"`
	Domain   string  `console:"domain" codec:"domain"`
	Hostname string  `console:"hostname" codec:"hostname"`
	Matches  bool    `console:"matches" codec:"matches"`

	Geolocation struct {
		City      string  `console:"city" codec:"city"`
		State     string  `console:"state" codec:"state"`
		Postcode  string  `console:"postcode" codec:"postcode"`
		Country   string  `console:"country" codec:"country"`
		Longitude float64 `console:"longitude" codec:"longitude"`
		Latitude  float64 `console:"latitude" codec:"latitude"`
	} `console:"geolocation" codec:"geolocation"`

	Organisation struct {
		ID           string `console:"id" codec:"id"`
		Name         string `console:"name" codec:"name"`
		Category     string `console:"category" codec:"category"`
		TotalDomains int64  `console:"totaldomains" codec:"totaldomains"`
		TotalIpsHave int64  `console:"totalipshave" codec:"totalipshave"`
		TotalIpsUsed int64  `console:"totalipsused" codec:"totalipsused"`
		Fortune1000  bool   `console:"fortune1000" codec:"fortune1000"`
	} `console:"organisation" codec:"organisation"`

	Statistics struct {
		IP struct {
			First     int64   `console:"first" codec:"first"`
			Daily     float64 `console:"daily" codec:"daily"`
			Monthly   float64 `console:"monthly" codec:"monthly"`
			Average   float64 `console:"average" codec:"average"`
			Percent   float64 `console:"percent" codec:"percent"`
			Blacklist float64 `console:"blacklist" codec:"blacklist"`
		} `console:"ip" codec:"ip"`
		Domain struct {
			First   int64   `console:"first" codec:"first"`
			Daily   float64 `console:"daily" codec:"daily"`
			Monthly float64 `console:"monthly" codec:"monthly"`
		} `console:"domain" codec:"domain"`
		Organisation struct {
			First   int64   `console:"first" codec:"first"`
			Daily   float64 `console:"daily" codec:"daily"`
			Monthly float64 `console:"monthly" codec:"monthly"`
		} `console:"organisation" codec:"organisation"`
	} `console:"statistics" codec:"statistics"`
}
