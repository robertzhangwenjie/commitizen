/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-07 18:14:41
 * @LastEditTime: 2022-09-04 11:20:49
 * @LastEditors: robert zhang
 * @Description:
 */
package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/glycerine/goconvey/convey"
)

func TestExecPath(t *testing.T) {
	b, wantErr := exec.Command("git", "--exec-path").Output()
	got, err := ExecPath()
	want := strings.TrimSpace(string(b))
	if want != got {
		t.Errorf("got %s, want %s", got, want)
	}

	if wantErr != err {
		t.Errorf("expected err %v,got %v", wantErr, err)
	}
}

func TestIsGitRepository(t *testing.T) {
	Convey("Given a path", t, func() {
		var path string
		Convey("When it doestn't exist", func() {
			path = "testdata/xxx"
			Convey("Then should return false", func() {
				So(IsGitRepository(path), ShouldBeFalse)
			})
		})

		Convey("When it's a git repo", func() {
			path = filepath.Join("testdata", "gitrepo")
			if err := exec.Command("git", "init", path).Run(); err != nil {
				t.Errorf("git init failed: %v", err)
			}
			defer os.RemoveAll(path)

			Convey("Then should return true", func() {
				So(IsGitRepository(path), ShouldBeTrue)
			})
		})
	})
}
