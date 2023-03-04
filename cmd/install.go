/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-07 11:04:45
 * @LastEditTime: 2023-03-04 12:05:02
 * @LastEditors: robert zhang
 * @Description:
 */
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/robertzhangwenjie/commitizen/git"
	"github.com/spf13/cobra"
)

func newCmdInstall() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "install git-cz command",
		Long: `Install git commit command 'git-cz' to git exec-path
for automating generate specified commit message format`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmdFilePath := os.Args[0]

			// find real path of cmdFile
			if _, err := os.Stat(cmdFilePath); err != nil {
				if cmdFilePath, err = exec.LookPath(cmdFilePath); err != nil {
					return err
				}
			}

			_, err := git.InstallSubCmd(cmdFilePath, "cz")
			if err != nil {
				return err
			}
			fmt.Printf(`Git SubCommand installed successfully, run with "git cz".`)
			return nil
		},
	}
	return cmd
}
