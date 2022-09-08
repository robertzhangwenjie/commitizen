/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-09-07 22:49:55
 * @LastEditTime: 2022-09-08 09:50:37
 * @LastEditors: robert zhang
 * @Description:
 */
package commit

import (
	"path/filepath"
	"testing"

	. "github.com/glycerine/goconvey/convey"
	"github.com/google/go-cmp/cmp"
)

func Test_getConfigFrom(t *testing.T) {
	Convey("Given a path", t, func() {
		var path string
		Convey("When the path contains a valid config named .git-czrc", func() {
			path = filepath.Join("testdata")
			Convey("Then should return a valid messageConfig", func() {
				got, err := getConfigFrom(path)
				want := &messageConfig{
					Template: "{{.type}}{{with .scope}}({{.}}){{end}}: {{.subject}}{{with .body}}\n\n{{.}}{{end}}{{with .footer}}\n\n{{.}}{{end}}",
					Items: []Form{
						{
							"type",
							"Select the type of change that you're committing:",
							"select",
							[]SelectOption{
								{"feat", "A new feature"},
								{"fix", "A bug fix"},
							},
							true,
						},
						{
							"scope",
							"Scope. Could be anything specifying place of the commit change (users, db, poll):",
							"input",
							nil,
							false,
						},
						{
							"body",
							"Body. Motivation for the change and contrast this with previous behavior:",
							"multiline",
							nil,
							false,
						},
					},
				}

				So(err, ShouldBeNil)
				So(cmp.Equal(got, want), ShouldBeTrue)
			})
		})
	})
}
