// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"
	"math"
)

const accuracy = 0.0000001

// Length returns a validation rule that checks if a value's length is within the specified range.
// If max is 0, it means there is no upper bound for the length.
// This rule should only be used for validating strings, slices, maps, and arrays.
// An empty value is considered valid. Use the Required rule to make sure a value is not empty.
func Range(min, max float64) *RangeRule {
	message := "the value must be empty"
	if min == 0 && max > 0 {
		message = fmt.Sprintf("the value must be no more than %v", max)
	} else if min > 0 && max == 0 {
		message = fmt.Sprintf("the value must be no less than %v", min)
	} else if min > 0 && max > 0 {
		message = fmt.Sprintf("the value must be between %v and %v", min, max)
	}
	return &RangeRule{
		min:     min,
		max:     max,
		message: message,
	}
}

type RangeRule struct {
	min, max float64
	message  string
	rune     bool
}

// Validate checks if the given value is valid or not.
func (v *RangeRule) Validate(value interface{}) error {
	value, isNil := Indirect(value)
	if isNil || IsEmpty(value) {
		return nil
	}
	var f64 float64
	switch value.(type) {
	case int:
		f64 = float64(value.(int))
	case int8:
		f64 = float64(value.(int))
	case int16:
		f64 = float64(value.(int))
	case int32:
		f64 = float64(value.(int))
	case int64:
		f64 = float64(value.(int))
	case uint:
		f64 = float64(value.(uint))
	case uint8:
		f64 = float64(value.(uint))
	case uint16:
		f64 = float64(value.(uint))
	case uint32:
		f64 = float64(value.(uint))
	case uint64:
		f64 = float64(value.(uint))
	case float32:
		f64 = float64(value.(float32))
	case float64:
		f64 = value.(float64)
	default:
		return fmt.Errorf("invalid value type: %v", value)
	}
	if math.Min(f64, v.min) == f64 && math.Abs(f64-v.min) > accuracy {
		return fmt.Errorf("smaller than %f: %v", v.min, value)
	} else if math.Max(f64, v.max) == f64 && math.Abs(f64-v.max) > accuracy {
		return fmt.Errorf("bigger than %f: %v", v.max, value)
	}
	return nil
}

// Error sets the error message for the rule.
func (v *RangeRule) Error(message string) *RangeRule {
	v.message = message
	return v
}
