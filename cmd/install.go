/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-07 11:04:45
 * @LastEditTime: 2022-08-08 11:46:01
 * @LastEditors: robert zhang
 * @Description:
 */
package cmd

import (
	"fmt"
	"os"

	"github.com/robertzhangwenjie/commitizen/git"
	"github.com/spf13/cobra"
)

func NewInstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "install git-cz command",
		Long: `Install git commit command 'git-cz' to git exec-path
for automating generate specified commit message format`,
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := git.InstallSubCmd(os.Args[0], "cz")
			if err != nil {
				return err
			}
			fmt.Printf(`Git SubCommand installed successfully, run with "git cz".`)
			return nil
		},
	}
	return cmd
}
