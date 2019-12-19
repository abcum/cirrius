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

package cpm

import (
	// Load globals
	_ "github.com/abcum/cirrius/cpm/console"
	_ "github.com/abcum/cirrius/cpm/context"
	_ "github.com/abcum/cirrius/cpm/process"

	// Load modules
	_ "github.com/abcum/cirrius/cpm/colour"
	_ "github.com/abcum/cirrius/cpm/dns"
	_ "github.com/abcum/cirrius/cpm/fake"
	_ "github.com/abcum/cirrius/cpm/fmt"
	_ "github.com/abcum/cirrius/cpm/go"
	_ "github.com/abcum/cirrius/cpm/html"
	_ "github.com/abcum/cirrius/cpm/http"
	_ "github.com/abcum/cirrius/cpm/inflect"
	_ "github.com/abcum/cirrius/cpm/path"
	_ "github.com/abcum/cirrius/cpm/pdf"
	_ "github.com/abcum/cirrius/cpm/routes"
	_ "github.com/abcum/cirrius/cpm/time"
	_ "github.com/abcum/cirrius/cpm/ws"

	// Load id modules
	_ "github.com/abcum/cirrius/cpm/id/guid"
	_ "github.com/abcum/cirrius/cpm/id/hashids"
	_ "github.com/abcum/cirrius/cpm/id/shortid"
	_ "github.com/abcum/cirrius/cpm/id/uuid"

	// Load file modules
	_ "github.com/abcum/cirrius/cpm/file"
	_ "github.com/abcum/cirrius/cpm/file/s3"

	// Load data modules
	_ "github.com/abcum/cirrius/cpm/db/surreal"

	// Load geo modules
	_ "github.com/abcum/cirrius/cpm/geo/haversine"

	// Load file modules
	_ "github.com/abcum/cirrius/cpm/text"
	_ "github.com/abcum/cirrius/cpm/text/allot"
	_ "github.com/abcum/cirrius/cpm/text/ascii"
	_ "github.com/abcum/cirrius/cpm/text/diff"
	_ "github.com/abcum/cirrius/cpm/text/levenshtein"
	_ "github.com/abcum/cirrius/cpm/text/lorem"
	_ "github.com/abcum/cirrius/cpm/text/patch"
	_ "github.com/abcum/cirrius/cpm/text/plain"
	_ "github.com/abcum/cirrius/cpm/text/rake"
	_ "github.com/abcum/cirrius/cpm/text/slug"
	_ "github.com/abcum/cirrius/cpm/text/urls"

	// Load image modules
	_ "github.com/abcum/cirrius/cpm/image"
	_ "github.com/abcum/cirrius/cpm/image/change"
	_ "github.com/abcum/cirrius/cpm/image/codabar"
	_ "github.com/abcum/cirrius/cpm/image/code"
	_ "github.com/abcum/cirrius/cpm/image/code128"
	_ "github.com/abcum/cirrius/cpm/image/code39"
	_ "github.com/abcum/cirrius/cpm/image/datamatrix"
	_ "github.com/abcum/cirrius/cpm/image/diff"
	_ "github.com/abcum/cirrius/cpm/image/ean"
	_ "github.com/abcum/cirrius/cpm/image/histogram"
	_ "github.com/abcum/cirrius/cpm/image/noise"
	_ "github.com/abcum/cirrius/cpm/image/patch"
	_ "github.com/abcum/cirrius/cpm/image/placeholder"
	_ "github.com/abcum/cirrius/cpm/image/qr"
	_ "github.com/abcum/cirrius/cpm/image/twooffive"

	// Load format modules
	_ "github.com/abcum/cirrius/cpm/format/csv"
	_ "github.com/abcum/cirrius/cpm/format/xml"

	// Load minify modules
	_ "github.com/abcum/cirrius/cpm/minify"
	_ "github.com/abcum/cirrius/cpm/minify/css"
	_ "github.com/abcum/cirrius/cpm/minify/html"
	_ "github.com/abcum/cirrius/cpm/minify/js"
	_ "github.com/abcum/cirrius/cpm/minify/json"
	_ "github.com/abcum/cirrius/cpm/minify/svg"
	_ "github.com/abcum/cirrius/cpm/minify/xml"

	// Load encoding modules
	_ "github.com/abcum/cirrius/cpm/encoding/base32"
	_ "github.com/abcum/cirrius/cpm/encoding/base64"
	_ "github.com/abcum/cirrius/cpm/encoding/binc"
	_ "github.com/abcum/cirrius/cpm/encoding/cbor"
	_ "github.com/abcum/cirrius/cpm/encoding/hjson"
	_ "github.com/abcum/cirrius/cpm/encoding/json"
	_ "github.com/abcum/cirrius/cpm/encoding/msgpack"

	// Load parse modules
	_ "github.com/abcum/cirrius/cpm/parse/url"

	// Load check modules
	_ "github.com/abcum/cirrius/cpm/check/abuse"
	_ "github.com/abcum/cirrius/cpm/check/blacklists"
	_ "github.com/abcum/cirrius/cpm/check/card"
	_ "github.com/abcum/cirrius/cpm/check/domain"
	_ "github.com/abcum/cirrius/cpm/check/email"
	_ "github.com/abcum/cirrius/cpm/check/mailgun"
	_ "github.com/abcum/cirrius/cpm/check/mailspike"
	_ "github.com/abcum/cirrius/cpm/check/ping"
	_ "github.com/abcum/cirrius/cpm/check/senderbase"
	_ "github.com/abcum/cirrius/cpm/check/senderscore"
	_ "github.com/abcum/cirrius/cpm/check/spam"
	_ "github.com/abcum/cirrius/cpm/check/vat"
	_ "github.com/abcum/cirrius/cpm/check/whois"
)
