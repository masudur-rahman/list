/*
Copyright © 2019 Masudur Rahman

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package types

import (
	"fmt"
)

// File contains file / directory information
type File struct {
	Name  string
	IsDir bool
}

func (f File) String() string {
	return fmt.Sprintf("%s%s  ", f.Name, dirExtension(f.IsDir))
}

type version struct {
	AppName string
	Version string
	Author  string
}

// Version contains appname, version and author name
var Version = version{
	AppName: "list",
	Version: "v0.1.0",
	Author:  "Masudur Rahman <masudjuly02@gmail.com>",
}

func (v version) String() string {
	return fmt.Sprintf("%s-%s ©%s", v.AppName, v.Version, v.Author)
}
func dirExtension(isDir bool) string {
	if isDir {
		return "/"
	}
	return ""
}
