/**
 * 认证模板
 */
package handler

import (
	"illusion"
	. "poem-illusion/message"
)

/**
 * 加载登录页面
 */
func login(c *illusion.Context) {
	c.View("auth/login.html", illusion.TemplateContext{})
}

/**
 * post过来的数据验证方法
 */
func authenticate(c *illusion.Context) {
	email := c.PostForm("email")
	pwd := c.PostForm("password")
	if email == "" || pwd == "" {
		c.Json(200, ErrMessage{
			Code: Failure,
			Text: "无法获取post数据",
		})
	}
	if email == "example@illusion.com" && pwd == "example" {
		sessionId, _ := c.GetSessionId()
		session, _ := illusion.Session().StartSession(sessionId)
		session.Store("email", email)
	}
	c.Redirect("/center")
}

func Auth() *illusion.Blueprint {
	auth := illusion.BluePrint("/auth", "_auth_")

	auth.Get("/login", login)

	auth.Post("/authenticate", authenticate)

	return auth
}
