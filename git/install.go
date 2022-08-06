/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-06 16:45:09
 * @LastEditTime: 2022-08-08 18:35:18
 * @LastEditors: robert zhang
 * @Description:
 */
package git

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func InstallSubCmd(cmdFilePath, subCmdName string) (string, error) {

	gitExecPath, err := ExecPath()
	if err != nil {
		return "", err
	}

	dstPath := filepath.Join(gitExecPath, "git-"+subCmdName)
	if _, err = copyFile(dstPath, cmdFilePath); err != nil {
		return "", fmt.Errorf("copy file failed: %w", err)
	}
	return dstPath, nil
}

func copyFile(dstFilePath, srcFilePath string) (written int64, err error) {
	src, err := os.Open(srcFilePath)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
