/*
 * 网站首页等无关用户相关的页面加载
 */
package handler

import (
	"illusion"
)

func index(b *illusion.Context) {
	b.View("index.html", illusion.TemplateContext{})
}

func Index() *illusion.Blueprint {
	home := illusion.BluePrint("/", "homeModule")

	home.Get("/", index)

	return home
}
