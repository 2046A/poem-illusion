/*
 * module
 * author: 
 */
package module

import "illusion"

func indexSetting(it *illusion.Context){
	it.View("setting/setting.html", illusion.TemplateContext{})
}

func Setting()*illusion.Blueprint{
	setting := illusion.BluePrint("/setting", "_setting_")
	setting.Get("/", indexSetting)

	return setting
}
