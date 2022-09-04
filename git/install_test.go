/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-06 16:52:22
 * @LastEditTime: 2022-09-04 11:37:33
 * @LastEditors: robert zhang
 * @Description:
 */
package git

import (
	"testing"

	. "github.com/glycerine/goconvey/convey"
)

func TestInstallSubCmd(t *testing.T) {
	Convey("Given a file which doesn't exist", t, func() {
		filepath := "/123/456"
		subCmdName := "123"
		Convey("When install it as git subcommand", func() {
			_, err := InstallSubCmd(filepath, subCmdName)
			Convey("Then should return err", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
