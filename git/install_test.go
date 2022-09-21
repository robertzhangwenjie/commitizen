/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-06 16:52:22
 * @LastEditTime: 2022-09-22 00:34:02
 * @LastEditors: robert zhang
 * @Description:
 */
package git

import (
	"fmt"
	"math/rand"
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

func Test_copyFile(t *testing.T) {
	Convey("Given a temporary dir", t, func() {
		dir := t.TempDir()
		Convey("When copy an exist file into it", func() {
			srcPath := filepath.Join("testdata", "testcmd")
			fi, err := os.Stat(srcPath)
			if err != nil {
				t.Errorf("get file info failed: %v", err)
			}
			srcSize := fi.Size()

			dstPath := filepath.Join(dir, fmt.Sprintf("%d", rand.Int()))
			dstSize, err := copyFile(dstPath, srcPath)
			Convey("Then it should copy successfully", func() {
				So(err, ShouldBeNil)
				So(dstSize, ShouldEqual, srcSize)
			})
		})

		Convey("When you copy a file that doesn't exist into it", func() {
			srcPath := filepath.Join("testdata", "notexist")
			_, err := os.Stat(srcPath)
			if err == nil {
				t.Error("this file already exists")
			}
			dstPath := filepath.Join(dir, fmt.Sprintf("err%d", rand.Int()))
			dstSize, err := copyFile(dstPath, srcPath)

			Convey("Then it should copy failed", func() {
				So(err, ShouldNotBeNil)
				So(dstSize, ShouldBeZeroValue)
			})

		})
	})

	Convey("Given an invalid path", t, func() {
		invalidPath := "/1/2/3/4/5/5"
		Convey("When copy an exist file to it", func() {
			srcPath := filepath.Join("testdata", "testcmd")
			_, err := os.Stat(srcPath)
			if err != nil {
				t.Errorf("get file info failed: %v", err)
			}

			written, err := copyFile(invalidPath, srcPath)
			Convey("Then should copy failed", func() {
				So(err, ShouldNotBeNil)
				So(written, ShouldBeZeroValue)
			})
		})
	})

}
