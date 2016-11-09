/*
 * SanPoem
 * author: 
 */
package main
import (
	"illusion"
	"poem-illusion/module"
)

func main(){
	poem := illusion.App()
	poem.ViewPath("templates")
	poem.LogPath("log")
	poem.Resource("resource")

	//开始注册blueprint
	poem.Register(module.HomeBlueprint())
	poem.Register(module.PoemBlueprint())
	poem.Register(module.CenterBlueprint())
	poem.Register(module.Setting())

	poem.Run(":9090")

}


