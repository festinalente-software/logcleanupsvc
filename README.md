## Motivation

Logfile cleanup tool for Windows which should be kept only for a specific timeframe

Tool can run as a service and can be installed without additional tools

## Config

Configure logcleanup via config.yml file (see example "logcleanup.yml"). 

Define config-file via "--config config.yml" when installing the service

## Install, Start, Stop, Remove

#### run as administrator to install service
> logcleanupsvc.exe install --config C:\full\path\to\config\logcleanup.yml

#### start service
> logcleanupsvc.exe start

#### stop service
> logcleanupsvc.exe stop

#### remove service
> logcleanupsvc.exe remove

## Test

> logcleanupsvc.exe install --config C:\Users\me\go\src\github.com\festinalente-software\logcleanupsvc\logcleanup_test.yml

## Development

Windows service handling is based on: 
https://github.com/golang/sys/tree/master/windows/svc/example

spf13\viper makes config file handling easy
