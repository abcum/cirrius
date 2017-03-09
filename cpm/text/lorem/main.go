// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lorem

import (
	"math/rand"
	"strings"

	"github.com/abcum/orbit"
)

func init() {
	orbit.Add("text/lorem", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
	}
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) size(min, max int) int {
	if min == max {
		return this.size(min, min+1)
	}
	if min > max {
		return this.size(max, min)
	}
	return rand.Int()%(max-min) + min
}

func (this *Module) rand() int {
	f := rand.Float32() * 100
	switch {
	case f < 1.939:
		return 1
	case f < 19.01:
		return 2
	case f < 38.00:
		return 3
	case f < 50.41:
		return 4
	case f < 61.00:
		return 5
	case f < 70.09:
		return 6
	case f < 78.97:
		return 7
	case f < 85.65:
		return 8
	case f < 90.87:
		return 9
	case f < 95.05:
		return 10
	case f < 97.27:
		return 11
	case f < 98.67:
		return 12
	case f < 100.0:
		return 13
	}
	return 2
}

func (this *Module) word(size int) string {

	if size < 1 {
		size = 1
	}
	if size > 13 {
		size = 13
	}

	for n := rand.Int() % len(latin); ; n++ {
		if n >= len(latin)-1 {
			n = 0
		}
		if len(latin[n]) == size {
			return latin[n]
		}
	}

	return ""

}

func (this *Module) Word(min, max int) string {

	s := this.size(min, max)

	return this.word(s)

}

func (this *Module) Sentence(min, max int) string {

	s := this.size(min, max)

	words := []string{}

	for i, c := 0, 0; i < s; i++ {

		word := this.word(this.rand())

		if i == 0 {
			word = strings.ToUpper(word[:1]) + word[1:]
		}

		words = append(words, word)

		if c >= 2 || i <= 2 || i >= s-1 {
			continue
		}

		if rand.Int()%s == 0 {
			words[i-1] += ","
			c++
		}

	}

	return strings.Join(words, " ") + "."

}

func (this *Module) Paragraph(min, max int) string {

	s := this.size(min, max)

	sents := []string{}

	for i := 0; i < s; i++ {
		sents = append(sents, this.Sentence(5, 22))
	}

	return strings.Join(sents, " ")

}
