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

The text/allot enables matching and parsing of predefined text commands.

	var lib = require('text/allot');

	lib.add("revert <commits:integer> commits on <project:string> at (live|test|dev)", function(match) {
		console.log('Commits:', match.integer('commits') );
		console.log('Project:', match.string('project') );
		console.log('Branch:', match.match(2) );
	});

	lib.add("commit all changes in <project:string> to (live|test|dev)", function(match) {
		console.log('Project:', match.string('project') );
		console.log('Branch:', match.match(1) );
	});

	lib.run(context.req.body);

*/
package allot
