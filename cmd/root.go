/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-07 11:20:22
 * @LastEditTime: 2022-08-23 16:01:05
 * @LastEditors: robert zhang
 * @Description:
 */
package cmd

import (
	"github.com/robertzhangwenjie/commitizen/commit"
	"github.com/spf13/cobra"
)

func NewRootCmd(name string) *cobra.Command {
	var dryRun bool
	cmd := &cobra.Command{
		Use:  name,
		Long: "Standardize git commit message tool",
		RunE: func(cmd *cobra.Command, args []string) error {
			return commit.Commit(dryRun)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	cmd.AddCommand(newCmdInstall(), newCmdVersion())
	cmd.Flags().BoolVar(&dryRun, "dry-run", dryRun, "preview result")
	return cmd
}
