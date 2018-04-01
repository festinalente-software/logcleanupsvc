// Copyright 2018 by festinalente-software. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"golang.org/x/sys/windows/svc/debug"
	"os"
	"time"
)

const svcName = "logcleanup"
const svcDisplayName = "Logfile cleaner"
const svcDesc = "Logcleanup: cleans all logfiles specified as pattern which have not been modified for the specified time"

type Cleanupconfig struct {
	Intervall      time.Duration
	ReloadOnChange bool
	// Parameters that can be changed by reloading following:
	Keeplimit   time.Duration
	Paused      bool
	Filepattern []string
}

var Config = Cleanupconfig{
	Intervall:      ParseDuration("10m"), // 10 Minutes
	ReloadOnChange: false,
	Keeplimit:      ParseDuration("5D"), // 5 Days
	Paused:         false,
	Filepattern:    []string{"c:\\logs\\*.log", os.Getenv("HOME") + "\\logs\\*.log"},
}

func LogConfig(elog debug.Log) {
	elog.Info(1, fmt.Sprintf("Config"))
	elog.Info(1, fmt.Sprintf("  Intervall = %v", Config.Intervall))
	elog.Info(1, fmt.Sprintf("  Keeplimit = %v", Config.Keeplimit))
	elog.Info(1, fmt.Sprintf("  Filepattern = %v", Config.Filepattern))
}

func initConfig() {

	var cfgFile = ""

	if len(os.Args) > 2 && os.Args[1] == "--config" {
		cfgFile = os.Args[2]
	} else if len(os.Args) > 3 && os.Args[2] == "--config" {
		cfgFile = os.Args[3]
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			SetConfigFromViper()
		}

		if Config.ReloadOnChange {
			viper.WatchConfig()
			viper.OnConfigChange(func(e fsnotify.Event) {
				if err := viper.ReadInConfig(); err == nil {
					SetConfigFromViper()
					elog.Info(1, fmt.Sprintf("read config file %s", viper.ConfigFileUsed()))
					LogConfig(elog)
				}
			})
		}
	}
}

func SetConfigFromViper() {
	intervall := ParseDuration(viper.GetString("Intervall"))
	if intervall > 0 {
		Config.Intervall = intervall
	}

	reloadOnChange := viper.GetBool("ReloadOnChange")
	Config.ReloadOnChange = reloadOnChange

	keeplimit := ParseDuration(viper.GetString("Keeplimit"))
	if keeplimit > 0 {
		Config.Keeplimit = keeplimit
	}

	filepattern := viper.GetStringSlice("Filepattern")
	if len(filepattern) > 0 {
		Config.Filepattern = append([]string{}, filepattern...)
	}
}
