/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-22 10:02:39
 * @LastEditTime: 2022-08-22 10:53:08
 * @LastEditors: robert zhang
 * @Description:
 */
package version

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/gosuri/uitable"
)

var (
	// GitVersion is semantic version.
	GitVersion = ""
	// BuildDate in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ').
	BuildTime = ""
	// GitCommit sha1 from git, output of $(git rev-parse HEAD).
	GitCommit = ""
	// GitTreeState state of git tree, either "clean" or "dirty".
	GitTreeState = ""
	GoVersion    = ""
)

// Info contains versioning information.
type Info struct {
	GitVersion   string `json:"GitVersion,omitempty"`
	GitCommit    string `json:"GitCommit"`
	GitTreeState string `json:"GitTreeState"`
	BuildTime    string `json:"BuildDate"`
	GoVersion    string `json:"GoVersion"`
	Compiler     string `json:"Compiler"`
	Platform     string `json:"Platform"`
}

// ToJSON returns the JSON string of version information.
func (info Info) ToJSON() string {
	s, _ := json.Marshal(info)

	return string(s)
}

func (info Info) Table() *uitable.Table {
	table := uitable.New()
	table.MaxColWidth = 80
	table.Wrap = true
	table.AddRow("GitVersion:", info.GitVersion)
	table.AddRow("GitCommit:", info.GitCommit)
	table.AddRow("GitTreeState:", info.GitTreeState)
	table.AddRow("BuildTime:", info.BuildTime)
	table.AddRow("GoVersion:", info.GoVersion)
	table.AddRow("Compiler:", info.Compiler)
	table.AddRow("Platform:", info.Platform)

	return table
}

// Get returns the overall codebase version. It's for detecting
// what code a binary was built from.
func Get() Info {
	// These variables typically come from -ldflags settings and in
	// their absence fallback to the settings in pkg/version/base.go
	return Info{
		GitVersion:   GitVersion,
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
		BuildTime:    BuildTime,
		GoVersion:    GoVersion,
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
