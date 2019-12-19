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

package args

import (
	"fmt"
	"strconv"

	"github.com/abcum/orbit"
	"github.com/abcum/otto"

	"github.com/abcum/cirrius/cpm/image"
)

func Size(orb *orbit.Orbit, fnc otto.FunctionCall, min, max int) {

	if len(fnc.ArgumentList) < min {
		if min == max {
			orb.Quit(fmt.Errorf("Expected %d arguments.", min))
		} else {
			orb.Quit(fmt.Errorf("Expected between %d and %d arguments.", min, max))
		}
	}

	if len(fnc.ArgumentList) > max {
		if min == max {
			orb.Quit(fmt.Errorf("Expected %d arguments.", min))
		} else {
			orb.Quit(fmt.Errorf("Expected between %d and %d arguments.", min, max))
		}
	}

}

func Any(orb *orbit.Orbit, fnc otto.FunctionCall, idx int) interface{} {
	exp, _ := fnc.Argument(idx).Export()
	return exp
}

func String(orb *orbit.Orbit, fnc otto.FunctionCall, idx int) string {
	exp, _ := fnc.Argument(idx).Export()
	val, _ := exp.(string)
	return val
}

func Image(orb *orbit.Orbit, fnc otto.FunctionCall, idx int) *image.Image {
	exp, _ := fnc.Argument(idx).Export()
	val, _ := exp.(*image.Image)
	return val
}

func Number(orb *orbit.Orbit, fnc otto.FunctionCall, idx int) int {
	exp, _ := fnc.Argument(idx).Export()
	switch val := exp.(type) {
	case int:
		return int(val)
	case int64:
		return int(val)
	case float64:
		return int(val)
	case string:
		if n, err := strconv.ParseInt(val, 10, 64); err == nil {
			return int(n)
		}
	}
	return 0
}

func Double(orb *orbit.Orbit, fnc otto.FunctionCall, idx int) float64 {
	exp, _ := fnc.Argument(idx).Export()
	switch val := exp.(type) {
	case int:
		return float64(val)
	case int64:
		return float64(val)
	case float64:
		return float64(val)
	case string:
		if n, err := strconv.ParseFloat(val, 64); err == nil {
			return float64(n)
		}
	}
	return 0
}

func Array(orb *orbit.Orbit, fnc otto.FunctionCall, idx int) []interface{} {
	exp, _ := fnc.Argument(idx).Export()
	val, _ := exp.([]interface{})
	if val != nil {
		return val
	}
	return []interface{}{}
}

func Object(orb *orbit.Orbit, fnc otto.FunctionCall, idx int) map[string]interface{} {
	exp, _ := fnc.Argument(idx).Export()
	val, _ := exp.(map[string]interface{})
	if val != nil {
		return val
	}
	return map[string]interface{}{}
}

func Value(orb *orbit.Orbit, exp interface{}) otto.Value {
	val, err := orb.ToValue(exp)
	if err != nil {
		orb.Quit(err)
	}
	return val
}
