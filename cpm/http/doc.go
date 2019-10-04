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

The http package enables making remote HTTP requests.

	var lib = require('http');
	var res = lib.get('https://abcum.com');
	res.pipe(context.res.body);

Or create a new HTTP request and specify custom HTTP options.

	var fs = require('file');
	var lib = require('http');
	var req = lib.new('POST');
	req.head("User-Agent", "Cirrius Bot");
	req.timeout('30s');
	req.url('https://example.com/uploader');
	req.body( file.load('test.png') );
	var res = req.do();

	res.pipe(context.res.body);

*/
package http
