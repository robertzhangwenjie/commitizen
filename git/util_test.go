/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-07 18:14:41
 * @LastEditTime: 2022-08-10 15:09:22
 * @LastEditors: robert zhang
 * @Description:
 */
package git

import (
	"os/exec"
	"strings"
	"testing"
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
	testcases := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "invalid path",
			path: "ttttt/",
			want: false,
		},
		{
			name: "current path",
			path: ".",
			want: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsGitRepository(tc.path)
			if got != tc.want {
				t.Errorf("expected %v, got %v", tc.want, got)
			}
		})
	}
}
