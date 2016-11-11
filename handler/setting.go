/*
 * module
 * author:
 */
package handler

import "illusion"

func settingIndex(it *illusion.Context) {
	it.View("setting/setting.html", illusion.TemplateContext{})
}

func Setting() *illusion.Blueprint {
	setting := illusion.BluePrint("/setting", "_setting_")
	setting.Get("/", settingIndex)

	return setting
}
