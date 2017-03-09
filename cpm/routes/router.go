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

package routes

import (
	"sort"

	"github.com/robertkrimen/otto"
)

type ranker []*route

func (s ranker) Len() int {
	return len(s)
}
func (s ranker) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ranker) Less(i, j int) bool {
	return s[i].rank < s[j].rank
}

type router struct {
	routes map[string][]*route
}

func (r *router) add(meth, path string, hand otto.Value) {

	route := &route{
		path:    path,
		method:  meth,
		handler: hand,
	}

	// Rank the route
	route.rank = route.sort()

	// Add the route to the tree
	r.routes[meth] = append(r.routes[meth], route)

	// Sort the routes according to rank
	sort.Sort(ranker(r.routes[meth]))

}

type route struct {
	rank    int
	path    string
	method  string
	handler otto.Value
}

func (r *route) sort() (rank int) {

	for _, c := range r.path {
		switch c {
		default:
			rank++
		case ':':
			rank += 100
		case '*':
			rank += 10000
		}
	}

	if rank > 20000 {
		panic("Path must only have 1 match all (*)")
	}

	if rank > 10000 && r.path[len(r.path)-1] != '*' {
		panic("Match all (*) must be at end of path")
	}

	return

}

func (r *route) test(path string) (map[string]string, bool) {

	var i int
	var j int
	var k string
	var v string

	param := make(map[string]string)

	for i < len(r.path) {

		switch {

		case r.path[i] == '*':
			i++
			v, j = consumeCatch(j, path)
			param["*"] = v

		case r.path[i] == ':':
			i++
			if k, i = consumeIdent(i, r.path); len(k) == 0 {
				return nil, false
			}
			if v, j = consumeIdent(j, path); len(v) == 0 {
				return nil, false
			}
			param[k] = v

		default:
			k, i = consumeChars(i, r.path)
			v, j = consumeCount(j, path, len(k))
			if k != v {
				return nil, false
			}

		}

	}

	if hasRemaining(j, path) {
		return nil, false
	}

	return param, true

}

func hasRemaining(pos int, path string) bool {
	return len(path) > pos+1
}

func consumeCatch(pos int, path string) (s string, i int) {
	return path[pos:], len(path)
}

func consumeIdent(pos int, path string) (s string, i int) {
	for i = pos; i < len(path); i++ {
		if path[i] == '/' || path[i] == ':' || path[i] == '*' {
			break
		}
	}
	return path[pos:i], i
}

func consumeChars(pos int, path string) (s string, i int) {
	for i = pos; i < len(path); i++ {
		if path[i] == ':' || path[i] == '*' {
			break
		}
	}
	return path[pos:i], i
}

func consumeCount(pos int, path string, c int) (s string, i int) {
	if len(path) < pos+c {
		i = len(path)
	} else {
		i = pos + c
	}
	return path[pos:i], i
}
