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

package split

import (
	"fmt"

	"gopkg.in/neurosnap/sentences.v1"

	"github.com/abcum/cirrius/cpm/text/split/data"
)

var langs = [...]string{
	// "cs",
	// "da",
	"de",
	// "el",
	"en",
	// "es",
	// "et",
	// "fi",
	"fr",
	// "it",
	// "nl",
	// "no",
	// "pl",
	// "pt",
	// "sl",
	// "sv",
	// "tr",
}

var learn map[string]*sentences.DefaultSentenceTokenizer

func init() {

	learn = make(map[string]*sentences.DefaultSentenceTokenizer)

	for _, lang := range langs {
		jsonf := fmt.Sprintf("%s.json", lang)
		asset := data.MustAsset(jsonf)
		store, _ := sentences.LoadTraining(asset)
		learn[lang] = sentences.NewSentenceTokenizer(store)
	}

}
