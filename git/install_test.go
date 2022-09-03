/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-06 16:52:22
 * @LastEditTime: 2022-09-03 17:24:07
 * @LastEditors: robert zhang
 * @Description:
 */
package git

import (
	"os"
	"path/filepath"
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
		wantErr bool
	}{
		{
			name: "noExistFile",
			args: args{
				filePath:   "/123/456",
				subCmdName: "123",
			},
			wantErr: true,
		},
		{
			name: "normalFile",
			args: args{
				filePath:   filepath.Join("testdata", "testcmd"),
				subCmdName: "testcmd",
			},
			wantErr: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			dstFilePath, err := InstallSubCmd(tc.args.filePath, tc.args.subCmdName)
			if err == nil {
				os.Remove(dstFilePath)
			}
			if (err != nil) != tc.wantErr {
				t.Errorf("InstallSubCmd() error = %v,wantErr %v ", err, tc.wantErr)
			}

		})
	}
}
