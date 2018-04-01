// Copyright 2018 by festinalente-software. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"regexp"
	"time"

	"strconv"
)

func ParseDuration(str string) time.Duration {
	durationRegex := regexp.MustCompile(`(?P<years>\d+Y)?(?P<months>\d+M)?(?P<weeks>\d+W)?(?P<days>\d+D)?(?P<hours>\d+h)?(?P<minutes>\d+m)?(?P<seconds>\d+s)?`)
	matches := durationRegex.FindStringSubmatch(str)

	years := ParseInt64(matches[1])
	months := ParseInt64(matches[2])
	weeks := ParseInt64(matches[3])
	days := ParseInt64(matches[4])
	hours := ParseInt64(matches[5])
	minutes := ParseInt64(matches[6])
	seconds := ParseInt64(matches[7])

	hour := int64(time.Hour)
	minute := int64(time.Minute)
	second := int64(time.Second)
	return time.Duration(years*24*365*hour + months*30*24*hour + weeks*7*24*hour + days*24*hour + hours*hour + minutes*minute + seconds*second)
}

func ParseInt64(value string) int64 {
	if len(value) == 0 {
		return 0
	}
	parsed, err := strconv.Atoi(value[:len(value)-1])
	if err != nil {
		return 0
	}
	return int64(parsed)
}
