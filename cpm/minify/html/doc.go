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

/*

The minify/html package minifies html content.

	var lib = require('minify/html');
	var out = lib.minify(data)
	console.log(out);

It is also possible to minify html data directly from a reader.

	var lib = require('minify/html');
	var dec = lib.decoder(context.req.body);
	var out = dec.string();
	console.log(out);

It is also possible to minify html data directly to a writer.

	var lib = require('minify/html');
	var enc = lib.encoder(context.res.body);
	enc.write(data);

*/
package html
