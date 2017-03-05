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

package web

import (
	// Load general
	_ "github.com/abcum/cirrius/npm"

	// Load globals
	_ "github.com/abcum/cirrius/npm/globals/console"
	_ "github.com/abcum/cirrius/npm/globals/context"
	_ "github.com/abcum/cirrius/npm/globals/process"

	// Load modules
	_ "github.com/abcum/cirrius/npm/modules/dns"
	_ "github.com/abcum/cirrius/npm/modules/http"
	_ "github.com/abcum/cirrius/npm/modules/https"
	_ "github.com/abcum/cirrius/npm/modules/os"
	_ "github.com/abcum/cirrius/npm/modules/path"
	_ "github.com/abcum/cirrius/npm/modules/routes"
	_ "github.com/abcum/cirrius/npm/modules/shortid"
	_ "github.com/abcum/cirrius/npm/modules/uuid"
	_ "github.com/abcum/cirrius/npm/modules/ws"

	// Load imaging modules
	_ "github.com/abcum/cirrius/npm/modules/imaging/bmp"
	_ "github.com/abcum/cirrius/npm/modules/imaging/colour"
	_ "github.com/abcum/cirrius/npm/modules/imaging/gif"
	_ "github.com/abcum/cirrius/npm/modules/imaging/imaging"
	_ "github.com/abcum/cirrius/npm/modules/imaging/jpeg"
	_ "github.com/abcum/cirrius/npm/modules/imaging/png"
	_ "github.com/abcum/cirrius/npm/modules/imaging/tiff"

	// Load encoder modules
	_ "github.com/abcum/cirrius/npm/modules/encoding/binc"
	_ "github.com/abcum/cirrius/npm/modules/encoding/cbor"
	_ "github.com/abcum/cirrius/npm/modules/encoding/json"
	_ "github.com/abcum/cirrius/npm/modules/encoding/msgpack"

	// Load barcode modules
	_ "github.com/abcum/cirrius/npm/modules/barcode/codabar"
	_ "github.com/abcum/cirrius/npm/modules/barcode/code128"
	_ "github.com/abcum/cirrius/npm/modules/barcode/code39"
	_ "github.com/abcum/cirrius/npm/modules/barcode/datamatrix"
	_ "github.com/abcum/cirrius/npm/modules/barcode/ean"
	_ "github.com/abcum/cirrius/npm/modules/barcode/qr"
	_ "github.com/abcum/cirrius/npm/modules/barcode/twooffive"
)
