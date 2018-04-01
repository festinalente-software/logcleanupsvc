// Copyright 2018 by festinalente-software. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func cleanupLogs() {
	if Config.Paused {
		return
	}
	for _, pattern := range Config.Filepattern {
		// elog.Info(1, fmt.Sprintf("cleaning logfiles with pattern '%s' after %d minutes", pattern, Config.Keeplimit/time.Minute))
		matches, err := filepath.Glob(pattern)
		if err != nil {
			elog.Warning(1, fmt.Sprintf("Filepath matching failed for %s with %v", pattern, err))
		} else {
			for _, filePath := range matches {
				// elog.Info(1, fmt.Sprintf("found file '%s'", filepath))
				info, err := os.Stat(filePath)
				if err != nil {
					elog.Error(1, fmt.Sprintf("Error get os.Stat for '%s': %v", filePath, err))
					break
				}
				if info.IsDir() {
					// ignore directories
					break
				}
				umodifiedDuration := time.Since(info.ModTime())
				if umodifiedDuration > Config.Keeplimit {
					elog.Info(1, fmt.Sprintf("Purge file '%s': has not been modified for %v", filePath, umodifiedDuration))
					if err := os.Remove(filePath); err != nil {
						elog.Error(1, fmt.Sprintf("Error deleting file '%s': %v", filePath, err))
					}
				}
			}
		}
	}
}
