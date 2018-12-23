package boot

import (
	"fmt"

	"github.com/adrianpk/kamien/boot"
)

func initAppContext() {
	{{AppNamePascalCase}}Context = boot.MakeAppContext()
	updateAssetsPath({{AppNamePascalCase}}Context)
	loadExtTemplates({{AppNamePascalCase}}Context)
}

func updateAssetsPath(*boot.AppContext) {
	{{AppNamePascalCase}}Context.Paths["assets"] = fmt.Sprintf("%s", *AssetsDir)
	{{AppNamePascalCase}}Context.Paths["templates"] = fmt.Sprintf("%s/templates", *AssetsDir)
	log.Infof("Assets dir: '%s'", {{AppNamePascalCase}}Context.Paths["assets"])
	log.Infof("Templates dir: '%s'", {{AppNamePascalCase}}Context.Paths["templates"])
}
