/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-06 16:52:22
 * @LastEditTime: 2022-09-21 20:24:20
 * @LastEditors: robert zhang
 * @Description:
 */
package git

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInstallSubCmd(t *testing.T) {
	Convey("Given a file which doesn't exist", t, func() {
		path := "/123/456"
		subCmdName := "123"
		Convey("When install it as git subcommand", func() {
			_, err := InstallSubCmd(path, subCmdName)
			Convey("Then should return err", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given a exists cmd file", t, func() {
		path := filepath.Join("testdata", "testcmd")
		subCmdName := "testcmd"
		Convey("When install it as git subcommand", func() {
			cmdPath, err := InstallSubCmd(path, subCmdName)
			if err == nil {
				defer os.Remove(cmdPath)
			}
			Convey("Then it should return nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
