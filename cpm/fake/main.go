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

package fake

import (
	"time"

	"github.com/abcum/orbit"
	"github.com/brianvoe/gofakeit"
)

func init() {
	orbit.Add("fake", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
	}
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) Seed(seed int64) {
	gofakeit.Seed(seed)
}

func (this *Module) Address() interface{} {
	return gofakeit.Address()
}

func (this *Module) Bool() bool {
	return gofakeit.Bool()
}

func (this *Module) Buzzword() string {
	return gofakeit.BuzzWord()
}

func (this *Module) City() string {
	return gofakeit.City()
}

func (this *Module) Color() string {
	return gofakeit.Color()
}

func (this *Module) Company() string {
	return gofakeit.Company()
}

func (this *Module) CompanySuffix() string {
	return gofakeit.CompanySuffix()
}

func (this *Module) Contact() interface{} {
	return gofakeit.Contact()
}

func (this *Module) Country() string {
	return gofakeit.Country()
}

func (this *Module) CreditCard() interface{} {
	return gofakeit.CreditCard()
}

func (this *Module) CreditCardCvv() string {
	return gofakeit.CreditCardCvv()
}

func (this *Module) CreditCardExp() string {
	return gofakeit.CreditCardExp()
}

func (this *Module) CreditCardNumber() int {
	return gofakeit.CreditCardNumber()
}

func (this *Module) CreditCardType() string {
	return gofakeit.CreditCardType()
}

func (this *Module) Date() time.Time {
	return gofakeit.Date()
}

func (this *Module) DateRange(start, end time.Time) time.Time {
	return gofakeit.DateRange(start, end)
}

func (this *Module) Day() int {
	return gofakeit.Day()
}

func (this *Module) DomainName() string {
	return gofakeit.DomainName()
}

func (this *Module) DomainSuffix() string {
	return gofakeit.DomainSuffix()
}

func (this *Module) Email() string {
	return gofakeit.Email()
}

func (this *Module) Extension() string {
	return gofakeit.Extension()
}
func (this *Module) FirefoxUserAgent() string {
	return gofakeit.FirefoxUserAgent()
}

func (this *Module) FirstName() string {
	return gofakeit.FirstName()
}

func (this *Module) Gender() string {
	return gofakeit.Gender()
}

func (this *Module) Generate(data string) string {
	return gofakeit.Generate(data)
}

func (this *Module) HexColor() string {
	return gofakeit.HexColor()
}

func (this *Module) Hour() int {
	return gofakeit.Hour()
}

func (this *Module) IPv4Address() string {
	return gofakeit.IPv4Address()
}

func (this *Module) IPv6Address() string {
	return gofakeit.IPv6Address()
}

func (this *Module) ImageURL(width, height int) string {
	return gofakeit.ImageURL(width, height)
}

func (this *Module) Job() interface{} {
	return gofakeit.Job()
}

func (this *Module) JobDescriptor() string {
	return gofakeit.JobDescriptor()
}

func (this *Module) JobLevel() string {
	return gofakeit.JobLevel()
}

func (this *Module) JobTitle() string {
	return gofakeit.JobTitle()
}

func (this *Module) LastName() string {
	return gofakeit.LastName()
}

func (this *Module) Latitude() float64 {
	return gofakeit.Latitude()
}

func (this *Module) Letter() string {
	return gofakeit.Letter()
}

func (this *Module) Lexify(data string) string {
	return gofakeit.Lexify(data)
}

func (this *Module) Longitude() float64 {
	return gofakeit.Longitude()
}

func (this *Module) Minute() int {
	return gofakeit.Minute()
}

func (this *Module) Month() string {
	return gofakeit.Month()
}

func (this *Module) Name() string {
	return gofakeit.Name()
}

func (this *Module) NamePrefix() string {
	return gofakeit.NamePrefix()
}

func (this *Module) NameSuffix() string {
	return gofakeit.NameSuffix()
}

func (this *Module) NanoSecond() int {
	return gofakeit.NanoSecond()
}

func (this *Module) Number(min, max int) int {
	return gofakeit.Number(min, max)
}

func (this *Module) Numerify(data string) string {
	return gofakeit.Numerify(data)
}

func (this *Module) OperaUserAgent() string {
	return gofakeit.OperaUserAgent()
}

func (this *Module) Paragraph(paragraphs, sentences, words int, separator string) string {
	return gofakeit.Paragraph(paragraphs, sentences, words, separator)
}

func (this *Module) Password(lower, upper, numeric, special, space bool, length int) string {
	return gofakeit.Password(lower, upper, numeric, special, space, length)
}

func (this *Module) Person() interface{} {
	return gofakeit.Person()
}

func (this *Module) RGBColor() []int {
	return gofakeit.RGBColor()
}

func (this *Module) SSN() string {
	return gofakeit.SSN()
}

func (this *Module) SafariUserAgent() string {
	return gofakeit.SafariUserAgent()
}

func (this *Module) SafeColor() string {
	return gofakeit.SafeColor()
}

func (this *Module) Second() int {
	return gofakeit.Second()
}

func (this *Module) Sentence(words int) string {
	return gofakeit.Sentence(words)
}

func (this *Module) ShuffleInts(ints []int) []int {
	return gofakeit.ShuffleInts(ints)
}

func (this *Module) ShuffleStrings(strings []string) []string {
	return gofakeit.ShuffleStrings(strings)
}

func (this *Module) State() string {
	return gofakeit.State()
}

func (this *Module) StateAbr() string {
	return gofakeit.StateAbr()
}

func (this *Module) Street() string {
	return gofakeit.Street()
}

func (this *Module) StreetName() string {
	return gofakeit.StreetName()
}

func (this *Module) StreetNumber() string {
	return gofakeit.StreetNumber()
}

func (this *Module) StreetPrefix() string {
	return gofakeit.StreetPrefix()
}

func (this *Module) StreetSuffix() string {
	return gofakeit.StreetSuffix()
}

func (this *Module) URL() string {
	return gofakeit.URL()
}

func (this *Module) UniqueID() string {
	return gofakeit.UniqueID()
}

func (this *Module) UserAgent() string {
	return gofakeit.UserAgent()
}

func (this *Module) Username() string {
	return gofakeit.Username()
}

func (this *Module) Word() string {
	return gofakeit.Word()
}

func (this *Module) Year() int {
	return gofakeit.Year()
}

func (this *Module) Zip() string {
	return gofakeit.Zip()
}
