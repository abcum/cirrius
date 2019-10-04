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

The image/patch package applies an image diff to another image.

	var fs = require('file');
	var img = require('image');
	var dif = require('image/diff');
	var lib = require('image/patch');
	var one = img.load( fs.load('diff1.png') );
	var two = img.load( fs.load('diff2.png') );
	var tre = dif(one, two);
	var out = lib(one, tre);
	out.pipe(context.res.body);

*/
package patch
