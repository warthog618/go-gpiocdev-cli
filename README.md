<!--
SPDX-FileCopyrightText: 2024 Kent Gibson <warthog618@gmail.com>

SPDX-License-Identifier: MIT
-->

# gpiocdev-cli

[![Build Status](https://img.shields.io/github/actions/workflow/status/warthog618/go-gpiocdev-cli/go.yml?logo=github&branch=master)](https://github.com/warthog618/go-gpiocdev-cli/actions/workflows/go.yml)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/warthog618/go-gpiocdev-cli)](https://pkg.go.dev/github.com/warthog618/go-gpiocdev-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/warthog618/go-gpiocdev-cli)](https://goreportcard.com/report/github.com/warthog618/go-gpiocdev-cli)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/warthog618/go-gpiocdev-cli/blob/master/LICENSE)

A command line utility to allow manual or scripted manipulation of GPIO lines
on Linux using the GPIO character device (**/dev/gpiochipN**).

This utility combines the Go equivalent of all the **libgpiod** command line
tools into a single tool.

> [!NOTE]
> This is a WIP.
>
> The existing version is split off from my **gpiod** library (now renamed **[gpiocdev](https://github.com/warthog618/go-gpiocdev)**) and mirrors the **libgpiod v1** tools, as at the time there was no **libgpiod v2** and so no tools using the latest kernel uAPI.
>
> The plan is to update these tools to mirror the **libgpiod v2** tools, but that isn't urgent given those tools now exist.
>
> The purpose of the split was to reduce the dependencies in **gpiocdev**, and to decouple updates to the cli from the core library.

## Installation

On Linux:

```shell
go install github.com/warthog618/go-gpiocdev-cli
```

## Usage

```shell
gpiocdev is a utility to control GPIO lines on Linux GPIO character devices

Usage:
  gpiocdev [flags]
  gpiocdev [command]

Available Commands:
  chip        Detect available GPIO chips
  edges       Monitor the state of a line or lines
  get         Get the state of a line or lines
  help        Help about any command
  line        Info about chip lines
  platform    Provide info about the platform iGPIO uAPI support.
  set         Set the state of a line or lines
  version     Display the version
  watch       Watch lines for changes to the line info

Flags:
  -h, --help   help for gpiocdev

Use "gpiocdev [command] --help" for more information about a command.

```
