# list
[![Go Report Card](https://goreportcard.com/badge/github.com/masudur-rahman/list)](https://goreportcard.com/report/github.com/masudur-rahman/list)
[![Build Status](https://travis-ci.org/masudur-rahman/list.svg?branch=master)](https://travis-ci.org/masudur-rahman/list)

`list` is a CLI tool which lists files and directories at the provided directory.

## Installation

#### `list` tool can be installed in the following ways:
- Run `go get github.com/masudur-rahman/list`
- Download appropriate version from the [release](https://github.com/masudur-rahman/list/releases) page & put in your `bin` folder

## Usage

- `list --help`
- `list [flags]`
    - Available flags : `file`, `dir`, `show-hidden` & `version`

- a dir argument can be provided with `list` command
- By default `list` shows all the visible files and directories
- If you provide `show-hidden` flag, hidden files and directories will also be shown
- `file` flag shows only files
- `dir` flag shows only directories
- `file` and `dir` flag at the same time shows all files and directories just like `list`.

