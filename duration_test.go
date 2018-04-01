// Copyright 2018 by festinalente-software. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDurationParse(t *testing.T) {
	var duration time.Duration

	const Day = 24 * time.Hour
	const Week = 7 * Day
	const Month = 30 * Day
	const Year = 365 * Day

	duration = ParseDuration("1s")
	assert.Equal(t, time.Duration(1*time.Second), duration)
	duration = ParseDuration("1m")
	assert.Equal(t, time.Duration(1*time.Minute), duration)
	duration = ParseDuration("1h")
	assert.Equal(t, time.Duration(1*time.Hour), duration)
	duration = ParseDuration("1D")
	assert.Equal(t, time.Duration(1*Day), duration)
	duration = ParseDuration("1W")
	assert.Equal(t, time.Duration(1*Week), duration)
	duration = ParseDuration("1Y")
	assert.Equal(t, time.Duration(1*Year), duration)

	duration = ParseDuration("2h1s")
	assert.Equal(t, time.Duration(2*time.Hour+1*time.Second), duration, fmt.Sprintf("wrong value=%s", duration))

	duration = ParseDuration("3D2h1s")
	assert.Equal(t, time.Duration(3*Day+2*time.Hour+1*time.Second), duration, fmt.Sprintf("wrong value=%s", duration))

	duration = ParseDuration("4W3D2h1s")
	assert.Equal(t, time.Duration(4*Week+3*Day+2*time.Hour+1*time.Second), duration, fmt.Sprintf("wrong value=%s", duration))

	duration = ParseDuration("5M3D2h1s")
	assert.Equal(t, time.Duration(5*Month+3*Day+2*time.Hour+1*time.Second), duration, fmt.Sprintf("wrong value=%s", duration))

	duration = ParseDuration("6Y5M3D2h1s")
	assert.Equal(t, time.Duration(6*Year+5*Month+3*Day+2*time.Hour+1*time.Second), duration, fmt.Sprintf("wrong value=%s", duration))

}
