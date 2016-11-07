package module

import (
	"illusion"
)

func index(b *illusion.Context){
	b.View("home/home.html", illusion.TemplateContext{})
}

func poem(b *illusion.Context){
	b.Json(200, "iii")
}

func HomeBlueprint() *illusion.Blueprint{
	home := illusion.BluePrint("/home", "homeModule")

	home.Get("/", index)

	//home.Get("/poem", poem)
	return home
}
