/*
 * module
 * author: 
 */
package module

import "illusion"

func centerIndex(c *illusion.Context){
	c.View("center/center.html", illusion.TemplateContext{})
}

func CenterBlueprint()*illusion.Blueprint{
	center := illusion.BluePrint("/center", "center_blueprint")
	center.Get("/", centerIndex)
	return center
}

