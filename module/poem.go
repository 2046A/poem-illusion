/*
 * module
 * author:
 */
package module

import "illusion"

func show(c *illusion.Context) {
	c.View("poem/show.html", illusion.TemplateContext{})
}

func PoemBlueprint() *illusion.Blueprint {
	poem := illusion.BluePrint("/poem", "poem")
	poem.Get("/show", show)

	return poem
}
