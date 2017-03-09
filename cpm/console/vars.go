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

package console

const LOG = "CREATE log SET `project`=$p, `request`=@request:$r, `kind`=$kind, `fold`=$fold, `file`=$file, `line`=$line, `char`=$char, `args`=$args, `time`=$time"

const ERR = "CREATE error SET `project`=$p, `request`=@request:$r, `kind`=$kind, `fold`=$fold, `file`=$file, `line`=$line, `char`=$char, `info`=$info, `time`=$time"

const REQ = "CREATE @request:$id SET `ip`=$ip, `time`=$time, `method`=$method, `duration`=$duration, `requestSize`=$requestSize, `responseSize`=$responseSize"
