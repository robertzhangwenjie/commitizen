/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-07 18:14:41
 * @LastEditTime: 2022-08-08 10:59:29
 * @LastEditors: robert zhang
 * @Description:
 */
package git

import "testing"

func TestExecPath(t *testing.T) {
	tests := []struct {
		name     string
		wantPath string
		wantErr  bool
	}{
		{
			name:     "normal",
			wantPath: "/opt/homebrew/Cellar/git/2.36.1/libexec/git-core",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPath, err := ExecPath()
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPath != tt.wantPath {
				t.Errorf("ExecPath() = %v, want %v", gotPath, tt.wantPath)
			}
		})
	}
}

func TestIsInGitRepository(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInGitRepository(); got != tt.want {
				t.Errorf("IsInGitRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
