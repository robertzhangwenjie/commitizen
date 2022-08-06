/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-07 18:14:41
 * @LastEditTime: 2022-08-08 11:27:44
 * @LastEditors: robert zhang
 * @Description:
 */
package git

import (
	"os/exec"
	"strings"
)

// ExecPath get git exec-path
func ExecPath() (path string, err error) {
	cmd := exec.Command("git", "--exec-path")
	output, err := cmd.Output()
	if err != nil {
		return
	}
	return strings.TrimSpace(string((output))), nil
}

// IsInGitRepository determine wheter the path is git repository
func IsInGitRepository() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	_, err := cmd.Output()
	return err == nil
}
