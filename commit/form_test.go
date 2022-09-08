/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-09-08 13:49:05
 * @LastEditTime: 2022-09-08 14:27:22
 * @LastEditors: robert zhang
 * @Description:
 */
package commit

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_loadForm(t *testing.T) {
	Convey("Given a message config", t, func() {
		config := &messageConfig{
			Items: []Form{
				{
					"type",
					"commit type",
					"select",
					[]SelectOption{{
						"feat",
						"new function",
					}, {
						"fix",
						"fix bug",
					}},
					true,
				},
			},
			Template: "template",
		}
		Convey("When load commit form by it", func() {
			_, template, err := loadForm(config)
			Convey("Then should get corresponding survey questions and templates", func() {
				So(err, ShouldBeNil)
				So(template, ShouldEqual, "template")
			})
		})
	})
}
