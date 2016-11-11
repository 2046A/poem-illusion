/*
 * SanPoem
 * author:
 */
package main

import (
	"illusion"
	. "poem-illusion/handler"
)

func main() {
	PoemApp := illusion.App()
	PoemApp.ViewPath("templates")
	PoemApp.LogPath("log")
	PoemApp.Resource("resource")

	//认证模板注册
	PoemApp.Register(Auth())

	//首页注册
	PoemApp.Register(Index())

	//用户中心注册
	PoemApp.Register(Center())

	//设置注册
	PoemApp.Register(Setting())

	PoemApp.Run(":9090")

}
