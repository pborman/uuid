// Copyright 2014 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"testing"
	"errors"
	"reflect"
)

func TestMarshalYAML(t *testing.T) {
	str := "f47ac10b-58cc-0372-8567-0e02b2c3d479"
	u := Parse(str)

	actual, err := u.MarshalYAML()

	if err != nil {
		t.Fatal(err)
	}

	if actual != str {
		t.Errorf("got %#v, want %#v", actual, str)
	}
}

func TestUnmarshalYAML(t *testing.T) {
	str := "f47ac10b-58cc-0372-8567-0e02b2c3d479"
	unmarshal := func(val interface{}) error {
		out := val.(*string)
		*out = str
		return nil
	}

	var actual UUID
	actual.UnmarshalYAML(unmarshal)

	if actual.String() != str {
		t.Errorf("got %#v, want %#v", actual.String(), str)
	}
}

func TestUnmarshalYAMLWithUnmarshalError(t *testing.T) {
	expected := errors.New("woops")
	unmarshal := func(val interface{}) error {
		return expected
	}

	var actual UUID
	err := actual.UnmarshalYAML(unmarshal)

	if !reflect.DeepEqual(&expected, &err) {
		t.Errorf("got %#v, want %#v", err, expected)
	}
}

func TestUnmarshalYAMLWithParseError(t *testing.T) {
	expected := errors.New("invalid UUID string: \"f\"")
	str := "f"
	unmarshal := func(val interface{}) error {
		out := val.(*string)
		*out = str
		return nil
	}

	var actual UUID
	err := actual.UnmarshalYAML(unmarshal)

	if !reflect.DeepEqual(&expected, &err) {
		t.Errorf("got %#v, want %#v", err, expected)
	}
}
