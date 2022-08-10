/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-07 18:49:50
 * @LastEditTime: 2022-08-10 15:07:23
 * @LastEditors: robert zhang
 * @Description:
 */
package commit

import (
	"fmt"

	"github.com/robertzhangwenjie/commitizen/git"
)

func Commit(dryRun bool) error {
	currentDir := "."
	if !git.IsGitRepository(currentDir) {
		return fmt.Errorf("not in git repository")
	}

	if !git.HasStagedFiles(currentDir) {
		return fmt.Errorf("no staged files to commit")
	}

	commitMsg, err := fillOutForm()
	if err != nil {
		return err
	}

	if dryRun {
		fmt.Printf("Commit Message:\n%v", string(commitMsg))
		return nil
	}

	return git.CommitMessage(commitMsg)
}
