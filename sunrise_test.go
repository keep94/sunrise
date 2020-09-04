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
	if dayLen != 24*time.Hour {
		t.Errorf("Expected light all day, but light %v", dayLen)
	}
	dayOrNight, _, _ := sunrise.DayOrNight(68, -118.25, startTime)
	if dayOrNight != sunrise.Day {
		t.Error("Expected phase to be day")
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
	dayOrNight, _, _ := sunrise.DayOrNight(68, -118.25, startTime)
	if dayOrNight != sunrise.Night {
		t.Error("Expected phase to be night")
	}
}

func TestDayOrNight(t *testing.T) {
	location, _ := time.LoadLocation("America/Chicago")
	expectedSunrise := time.Date(2014, 11, 19, 7, 0, 24, 0, location)
	expectedSunset := time.Date(2014, 11, 19, 17, 22, 57, 0, location)
	times := []time.Time{
		time.Date(2014, 11, 19, 7, 0, 24, 0, location),
		time.Date(2014, 11, 19, 8, 0, 0, 0, location),
		time.Date(2014, 11, 19, 9, 0, 0, 0, location),
		time.Date(2014, 11, 19, 10, 0, 0, 0, location),
		time.Date(2014, 11, 19, 11, 0, 0, 0, location),
		time.Date(2014, 11, 19, 12, 0, 0, 0, location),
		time.Date(2014, 11, 19, 13, 0, 0, 0, location),
		time.Date(2014, 11, 19, 14, 0, 0, 0, location),
		time.Date(2014, 11, 19, 15, 0, 0, 0, location),
		time.Date(2014, 11, 19, 16, 0, 0, 0, location),
		time.Date(2014, 11, 19, 17, 0, 0, 0, location),
		time.Date(2014, 11, 19, 17, 22, 56, 0, location),
	}
	for _, cTime := range times {
		dayOrNight, start, end := sunrise.DayOrNight(32.9, -96.2, cTime)
		if dayOrNight != sunrise.Day {
			t.Fatal("Expected day")
		}
		if start != expectedSunrise || end != expectedSunset {
			t.Errorf("For %v, got sunrise %v; sunset %v", cTime, start, end)
		}
	}
	dayOrNight, _, _ := sunrise.DayOrNight(32.9, -96.2, expectedSunset)
	if dayOrNight != sunrise.Night {
		t.Error("Expected night")
	}
	expectedSunset = time.Date(2014, 11, 18, 17, 23, 25, 0, location)
	times = []time.Time{
		time.Date(2014, 11, 18, 17, 23, 25, 0, location),
		time.Date(2014, 11, 18, 18, 0, 0, 0, location),
		time.Date(2014, 11, 18, 19, 0, 0, 0, location),
		time.Date(2014, 11, 18, 20, 0, 0, 0, location),
		time.Date(2014, 11, 18, 21, 0, 0, 0, location),
		time.Date(2014, 11, 18, 22, 0, 0, 0, location),
		time.Date(2014, 11, 18, 23, 0, 0, 0, location),
		time.Date(2014, 11, 19, 0, 0, 0, 0, location),
		time.Date(2014, 11, 19, 1, 0, 0, 0, location),
		time.Date(2014, 11, 19, 2, 0, 0, 0, location),
		time.Date(2014, 11, 19, 3, 0, 0, 0, location),
		time.Date(2014, 11, 19, 4, 0, 0, 0, location),
		time.Date(2014, 11, 19, 5, 0, 0, 0, location),
		time.Date(2014, 11, 19, 6, 0, 0, 0, location),
		time.Date(2014, 11, 19, 7, 0, 0, 0, location),
		time.Date(2014, 11, 19, 7, 0, 23, 0, location),
	}
	for _, cTime := range times {
		dayOrNight, start, end := sunrise.DayOrNight(32.9, -96.2, cTime)
		if dayOrNight != sunrise.Night {
			t.Fatal("Expected night")
		}
		if start != expectedSunset || end != expectedSunrise {
			t.Errorf("For %v, got sunset %v; sunrise %v", cTime, start, end)
		}
	}
}
