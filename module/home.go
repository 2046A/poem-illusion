package module

import (
	"illusion"
)

func index(b *illusion.Context){
	b.View("home/home.html", illusion.TemplateContext{})
}

func HomeBlueprint() *illusion.Blueprint{
	home := illusion.BluePrint("/home", "homeModule")

	home.Get("/", index)

	return home
}
