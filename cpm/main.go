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
	_ "github.com/abcum/cirrius/cpm/fmt"
	_ "github.com/abcum/cirrius/cpm/http"
	_ "github.com/abcum/cirrius/cpm/image"
	_ "github.com/abcum/cirrius/cpm/path"
	_ "github.com/abcum/cirrius/cpm/pdf"
	_ "github.com/abcum/cirrius/cpm/routes"

	// Load id modules
	_ "github.com/abcum/cirrius/cpm/id/uuid"

	// Load file modules
	_ "github.com/abcum/cirrius/cpm/file"
	_ "github.com/abcum/cirrius/cpm/file/s3"

	// Load data modules
	_ "github.com/abcum/cirrius/cpm/db/surreal"
)
