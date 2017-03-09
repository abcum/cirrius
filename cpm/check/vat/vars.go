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

package vat

import (
	"encoding/xml"
	"regexp"
)

type response struct {
	XMLName xml.Name `xml:"Envelope"`
	Soap    struct {
		XMLName xml.Name `xml:"Body"`
		Soap    struct {
			XMLName     xml.Name `xml:"checkVatResponse"`
			CountryCode string   `xml:"countryCode"`
			VATNumber   string   `xml:"vatNumber"`
			RequestDate string   `xml:"requestDate"` // 2015-03-06+01:00
			Valid       bool     `xml:"valid"`
			Name        string   `xml:"name"`
			Address     string   `xml:"address"`
		}
	}
}

var headtype = "text/xml;charset=UTF-8"

var endpoint = "http://ec.europa.eu/taxation_customs/vies/services/checkVatService"

var envelope = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
	<soapenv:Header/>
	<soapenv:Body>
  		<checkVat xmlns="urn:ec.europa.eu:taxud:vies:services:checkVat:types">
			<countryCode>{{.country}}</countryCode>
    		<vatNumber>{{.vnumber}}</vatNumber>
		</checkVat>
	</soapenv:Body>
</soapenv:Envelope>
`

var patterns = map[string]*regexp.Regexp{
	"AT": regexp.MustCompile("U[A-Z\\d]{8}"),
	"BE": regexp.MustCompile("(0\\d{9}|\\d{10})"),
	"BG": regexp.MustCompile("\\d{9,10}"),
	"CY": regexp.MustCompile("\\d{8}[A-Z]"),
	"CZ": regexp.MustCompile("\\d{8,10}"),
	"DE": regexp.MustCompile("\\d{9}"),
	"DK": regexp.MustCompile("(\\d{2} ?){3}\\d{2}"),
	"EE": regexp.MustCompile("\\d{9}"),
	"EL": regexp.MustCompile("\\d{9}"),
	"ES": regexp.MustCompile("[A-Z]\\d{7}[A-Z]|\\d{8}[A-Z]|[A-Z]\\d{8}"),
	"FI": regexp.MustCompile("\\d{8}"),
	"FR": regexp.MustCompile("([A-Z]{2}|\\d{2})\\d{9}"),
	"GB": regexp.MustCompile("\\d{9}|\\d{12}|(GD|HA)\\d{3}"),
	"HR": regexp.MustCompile("\\d{11}"),
	"HU": regexp.MustCompile("\\d{8}"),
	"IE": regexp.MustCompile("[A-Z\\d]{8}|[A-Z\\d]{9}"),
	"IT": regexp.MustCompile("\\d{11}"),
	"LT": regexp.MustCompile("(\\d{9}|\\d{12})"),
	"LU": regexp.MustCompile("\\d{8}"),
	"LV": regexp.MustCompile("\\d{11}"),
	"MT": regexp.MustCompile("\\d{8}"),
	"NL": regexp.MustCompile("\\d{9}B\\d{2}"),
	"PL": regexp.MustCompile("\\d{10}"),
	"PT": regexp.MustCompile("\\d{9}"),
	"RO": regexp.MustCompile("\\d{2,10}"),
	"SE": regexp.MustCompile("\\d{12}"),
	"SI": regexp.MustCompile("\\d{8}"),
	"SK": regexp.MustCompile("\\d{10}"),
}
