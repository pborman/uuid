// Copyright 2014 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package uuid

import (
	"fmt"
)

func (u UUID) MarshalYAML() (interface{}, error) {
	return u.String(), nil
}

func (u *UUID) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var data string
	err := unmarshal(&data)
	if err != nil {
		return err
	}

	uu := Parse(data)
	if uu == nil {
		return fmt.Errorf("invalid UUID string: %q", data)
	}
	*u = uu
	return nil
}
