/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-22 10:09:33
 * @LastEditTime: 2022-08-23 15:03:44
 * @LastEditors: robert zhang
 * @Description:
 */
package cmd

import (
	"fmt"

	"github.com/robertzhangwenjie/commitizen/pkg/version"
	"github.com/spf13/cobra"
)

func newCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "print tool version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(version.Get().ToJSON())
			return nil
		},
	}
	return cmd
}
