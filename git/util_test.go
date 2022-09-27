/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-07 18:14:41
 * @LastEditTime: 2022-09-27 23:18:29
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

	. "github.com/smartystreets/goconvey/convey"
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
	Convey("Given a path which doesn't exist", t, func() {
		path := "testdata/xxx"
		Convey("When determining whether it is a git repo", func() {
			result := IsGitRepository(path)
			Convey("Then should return false", func() {
				So(result, ShouldBeFalse)
			})
		})

	})

	Convey("Given a path which is a git repo", t, func() {
		path := filepath.Join("testdata", "gitrepo")
		if err := exec.Command("git", "init", path).Run(); err != nil {
			t.Errorf("git init failed: %v", err)
		}
		defer os.RemoveAll(path)
		Convey("When determining whether it is a git repo", func() {
			result := IsGitRepository(path)
			Convey("Then should return true", func() {
				So(result, ShouldBeTrue)
			})
		})
	})
}
