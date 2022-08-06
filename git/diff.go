/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-07 19:19:42
 * @LastEditTime: 2022-08-07 19:23:40
 * @LastEditors: robert zhang
 * @Description:
 */
package git

import "os/exec"

// HasStagedFiles determine whether has changes to be committed
func HasStagedFiles() bool {
	cmd := exec.Command("git", "diff", "--quiet", "--cached")
	_, err := cmd.Output()
	return err != nil
}
