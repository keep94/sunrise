// Copyright 2014 Travis Keep. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or
// at http://opensource.org/licenses/BSD-3-Clause.

package sunrise_test

import (
  "github.com/keep94/sunrise"
  "testing"
  "time"
)

func TestLightAllDay(t *testing.T) {
  var s sunrise.Sunrise

  location, _ := time.LoadLocation("America/Los_Angeles")
  startTime := time.Date(2014, 6, 20, 0, 0, 0, 0, location)

  // Above tropic of cancer should be light all day
  s.Around(68, -118.25, startTime)
  dayLen := s.Sunset().Sub(s.Sunrise())
  if dayLen != 24 * time.Hour {
    t.Errorf("Expected light all day, but light %v", dayLen)
  }
}

func TestDarkAllDay(t *testing.T) {
  var s sunrise.Sunrise
  location, _ := time.LoadLocation("America/Los_Angeles")
  startTime := time.Date(2014, 12, 20, 0, 0, 0, 0, location)

  // Above tropic of cancer should be dark all day
  s.Around(68, -118.25, startTime)
  dayLen := s.Sunset().Sub(s.Sunrise())
  if dayLen != 0 {
    t.Errorf("Expected dark all day, but light %v", dayLen)
  }
}
