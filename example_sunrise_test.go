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

  for s.Sunrise().Before(startTime) {
    s.AddDays(1)
  }

  formatStr := "Jan 2 15:04:05"
  for i := 0; i < 5; i++ {
    fmt.Printf("Sunrise: %s Sunset: %s\n", s.Sunrise().Format(formatStr), s.Sunset().Format(formatStr))
    s.AddDays(1)
  }
  // Output:
  // Sunrise: Jun 1 05:44:01 Sunset: Jun 1 20:00:42
  // Sunrise: Jun 2 05:43:45 Sunset: Jun 2 20:01:17
  // Sunrise: Jun 3 05:43:30 Sunset: Jun 3 20:01:52
  // Sunrise: Jun 4 05:43:17 Sunset: Jun 4 20:02:25
  // Sunrise: Jun 5 05:43:05 Sunset: Jun 5 20:02:57
}
