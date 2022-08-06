/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-06 16:52:22
 * @LastEditTime: 2022-08-07 10:55:53
 * @LastEditors: robert zhang
 * @Description:
 */
package git

import (
	"testing"
)

func TestInstallSubCmd(t *testing.T) {
	type args struct {
		filePath   string
		subCmdName string
	}

	testcases := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "noExistFile",
			args: args{
				filePath:   "/123/456",
				subCmdName: "123",
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			dstFilePath, err := InstallSubCmd(tc.args.filePath, tc.args.subCmdName)
			if (err != nil) != tc.wantErr && dstFilePath != tc.want {
				t.Errorf("InstallSubCmd() error = %v,wantErr %v", err, tc.wantErr)
			}
		})
	}

}
