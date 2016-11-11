/*
 * module
 * author:
 */
package handler

import "illusion"

func centerIndex(c *illusion.Context) {
	c.View("center/index.html", illusion.TemplateContext{})
}

func Center() *illusion.Blueprint {
	center := illusion.BluePrint("/center", "center_blueprint")
	center.Get("/", centerIndex)
	return center
}
