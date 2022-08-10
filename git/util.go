/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-07 18:14:41
 * @LastEditTime: 2022-08-10 15:06:26
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

// IsGitRepository determine wheter the path is git repository
func IsGitRepository(dir string) bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	cmd.Dir = dir
	_, err := cmd.Output()
	return err == nil
}

// HasStagedFiles determine whether has changes to be committed
func HasStagedFiles(dir string) bool {
	cmd := exec.Command("git", "diff", "--quiet", "--cached")
	cmd.Dir = dir
	_, err := cmd.Output()
	return err != nil
}
