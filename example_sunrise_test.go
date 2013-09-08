// Copyright 2013 Travis Keep. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or
// at http://opensource.org/licenses/BSD-3-Clause.

package sunrise_test

import (
  "fmt"
  "github.com/keep94/sunrise"
  "time"
)

// Show sunrise and sunset for first 5 days of June in LA
func ExampleSunrise() {
  var s sunrise.Sunrise

  // Start time is June 1, 2013 PST
  location, _ := time.LoadLocation("America/Los_Angeles")
  startTime := time.Date(2013, 6, 1, 0, 0, 0, 0, location)

  // Coordinates of LA are 34.05N 118.25W
  s.Around(34.05, -118.25, startTime)

  if (s.Sunrise().Before(startTime)) {
    s.Skip(1)
  }

  formatStr := "Jan 2 15:04:05"
  for i := 0; i < 5; i++ {
    fmt.Printf("Sunrise: %s Sunset: %s\n", s.Sunrise().Format(formatStr), s.Sunset().Format(formatStr))
    s.Skip(1)
  }
  // Output:
  // Jun 1 5:35 Jun 1 17:35
}
