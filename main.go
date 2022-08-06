/*
 * @Author: robert zhang <robertzhangwenjie@gmail.com>
 * @Date: 2022-08-04 11:05:08
 * @LastEditTime: 2022-08-07 17:57:49
 * @LastEditors: robert zhang
 * @Description:
 */
package main

import (
	"fmt"

	"github.com/robertzhangwenjie/commitizen/cmd"
)

func main() {
	c := cmd.NewRootCmd("commitizen")
	if err := c.Execute(); err != nil {
		fmt.Println(err)
	}

}
