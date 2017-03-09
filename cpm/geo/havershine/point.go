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

package havershine

import (
	"math"

	"github.com/abcum/orbit"
)

type Point struct {
	orb *orbit.Orbit
	Lat float64 `console:"lat" codec:"lat"`
	Lng float64 `console:"lng" codec:"lng"`
}

func NewPoint(orb *orbit.Orbit, lat, lng float64) *Point {
	return &Point{
		orb: orb,
		Lat: lat,
		Lng: lng,
	}
}

func (this *Point) Distance(to *Point) float64 {

	lata, lnga := this.Lat*radian, this.Lng*radian
	latb, lngb := to.Lat*radian, to.Lng*radian
	latc, lngc := lata-latb, lnga-lngb

	dis := math.Pow(math.Sin(latc/2), 2) + math.Cos(lata)*math.Cos(latb)*math.Pow(math.Sin(lngc/2), 2)

	c := 2 * math.Atan2(math.Sqrt(dis), math.Sqrt(1-dis))

	return (earthr * c)

}
