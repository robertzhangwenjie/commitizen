/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-06 16:52:22
 * @LastEditTime: 2023-03-04 12:11:28
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
	Convey("Given a cmdFilePath", t, func() {
		cmdFilePath := filepath.Join("testdata", "testcmd")
		Convey("When install as git subCmd", func() {
			path, err := InstallSubCmd(cmdFilePath, "test")
			Convey("Then it shoud install successfully", func() {
				So(err, ShouldBeNil)
				So(path, ShouldNotBeEmpty)
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
