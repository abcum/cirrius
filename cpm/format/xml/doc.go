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

The format/xml package enables creating or manipulating xml data.

	var lib = require('format/xml');
	var doc = lib.new();
	doc.Instruction('xml', 'version="1.0" encoding="UTF-8"');

	var res = doc.Element("Response");
	var elm = res.Element("Say");
	elm.Set("voice", "woman");
	elm.Text("Please leave a message after the tone.");
	var elm = res.Element("Record");
	elm.Set("maxLength", "20");

	doc.pipe(context.res.body);

Or load a previously generated xml document, and manipulate its contents.

	var lib = require('format/xml');
	var doc = lib.load(context.req.body);
	doc.Instruction('xml', 'version="1.0" encoding="UTF-8"');

	var elm = doc.Find("People/Person[@id='sally']");
	elm.Element("Firstname").Set("value", "Sally");
	elm.Element("Lastname").Set("value", "Morgan");
	elm.AddComment("Sally is a great person");
	elm.AddAttribute("gender", "femals");
	elm.Text("Mrs. Sally Morgan");

	doc.pipe(context.res.body)

*/
package xml
