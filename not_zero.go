// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validation

import "errors"

// NotEmpty is a validation rule that checks if a value is not nil.
// NotEmpty only handles types including interface, pointer, slice, and map.
// All other types are considered valid.
var NotEmpty = &notEmptyRule{message: "is not empty"}

type notEmptyRule struct {
	message string
}

// Validate checks if the given value is valid or not.
func (r *notEmptyRule) Validate(value interface{}) error {
	value, isNil := Indirect(value)
	if isNil || IsEmpty(value) {
		return errors.New(r.message)
	}
	return nil
}

// Error sets the error message for the rule.
func (r *notEmptyRule) Error(message string) *notEmptyRule {
	return &notEmptyRule{
		message: message,
	}
}
