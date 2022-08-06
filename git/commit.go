/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-06 10:02:03
 * @LastEditTime: 2022-08-06 10:20:25
 * @LastEditors: robert zhang
 * @Description:
 */
package git

import (
	"os"
	"os/exec"
)

func CommitMessage(msg []byte) error {
	cmd := exec.Command("git", "commit", "-m", string(msg))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
