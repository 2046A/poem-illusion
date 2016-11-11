/*
 * module
 * author:
 */
package handler

import "illusion"

func show(c *illusion.Context) {
	c.View("poem/show.html", illusion.TemplateContext{})
}

func PoemBlueprint() *illusion.Blueprint {
	poem := illusion.BluePrint("/poem", "poem")

	poem.Get("/:name", show)

	return poem
}
