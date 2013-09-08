// Copyright 2013 Travis Keep. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or
// at http://opensource.org/licenses/BSD-3-Clause.

// Package sunrise computes sunrises and sunsets using wikipedia article
// http://en.wikipedia.org/wiki/Sunrise_equation. Testing at my
// latitude and longitude in California shows that computed sunrises and
// sunsets can vary by as much as 2 minutes from those that NOAA reports
// at http://www.esrl.noaa.gov/gmd/grad/solcalc/sunrise.html.
package sunrise

import (
  "math"
  "time"
)

const (
  jepoch = float64(2451545.0)
  uepoch = int64(946728000.0)
)

// Sunrise gives sunrise and sunset times.
type Sunrise struct {
  location *time.Location
  sinLat float64
  cosLat float64
  jstar float64
  solarNoon float64
  hourAngleInDays float64
}

// Around computes the sunrise and sunset times for latitude and longitude
// around currentTime. The computed sunrise will be no earlier than 24
// hours before currentTime and the computed sunset will be no later than
// 24 hours after currentTime. The computed sunset is no more than 24 hours
// after the current sunrise.
func (s *Sunrise) Around(latitude, longitude float64, currentTime time.Time) {
  s.location = currentTime.Location()
  s.sinLat = sin(latitude)
  s.cosLat = cos(latitude)
  s.jstar = math.Floor(
      julianDay(currentTime.Unix()) - 0.0009 + longitude / 360.0 + 0.5) + 0.0009 - longitude / 360.0
  s.computeSolarNoonHourAngle()
}

// Skip computes the sunrise and sunset a set number of days after
// (or before if argument is negative) the current sunrise and sunset at the
// same latitude and longitude.
func (s *Sunrise) Skip(days int) {
  s.jstar += float64(days)
  s.computeSolarNoonHourAngle()
}

// Sunrise returns the current computed sunrise.
func (s *Sunrise) Sunrise() time.Time {
  return goTime(s.solarNoon - s.hourAngleInDays, s.location)
}

// Sunset returns the current computed sunset.
func (s *Sunrise) Sunset() time.Time {
  return goTime(s.solarNoon + s.hourAngleInDays, s.location)
}

func (s *Sunrise) computeSolarNoonHourAngle() {
  ma := mod360(357.5291 + 0.98560028 * (s.jstar - jepoch))
  center := 1.9148 * sin(ma) + 0.02 * sin(2.0 * ma) + 0.0003 * sin(3.0 * ma)
  el := mod360(ma + 102.9372 + center + 180.0)
  s.solarNoon = s.jstar + 0.0053 * sin(ma) - 0.0069 * sin(2.0 * el)
  declination := asin(sin(el) * sin(23.45))
  s.hourAngleInDays = acos((sin(-0.83) - s.sinLat * sin(declination)) / (s.cosLat * cos(declination))) / 360.0
}

func julianDay(unix int64) float64 {
  return float64(unix - uepoch) / 86400.0 + jepoch
}

func goTime(julianDay float64, loc *time.Location) time.Time {
  unix := uepoch + int64((julianDay - jepoch) * 86400.0)
  result := time.Unix(unix, 0)
  if result.Location() == loc {
    return result
  }
  return result.In(loc)
}

func sin(degrees float64) float64 {
  return math.Sin(degrees * math.Pi / 180.0)
}

func cos(degrees float64) float64 {
  return math.Cos(degrees * math.Pi / 180.0)
}

func asin(x float64) float64 {
  return math.Asin(x) * 180.0 / math.Pi
}

func acos(x float64) float64 {
  return math.Acos(x) * 180.0 / math.Pi
}

func mod360(x float64) float64 {
  return x - 360.0 * math.Floor(x / 360.0)
}
