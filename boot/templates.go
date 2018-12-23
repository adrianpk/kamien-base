package boot

import (
	"os"
	"path"
	"path/filepath"

	"{{.Package}}/common"
	tmpl "{{.Package}}/templates"
	"github.com/adrianpk/kamien/boot"
	"github.com/alecthomas/template"
	htmlTemplate "github.com/arschles/go-bindata-html-template"
	bindata "github.com/golang-migrate/migrate/source/go_bindata"
)

func templatesFileMap(dir string) map[string]*template.Template {
	ts := make(map[string]*template.Template)
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			checkErr(err)
			if filepath.Ext(path) == ".tmpl" {
				ts[path] = template.New(path)
			}
			return nil
		})
	checkErr(err)
	return ts
}

func loadTemplates(ac *boot.AppContext) {
	log.Info("Loading assets...")
	templatesMap := ac.Templates
	assets := bindata.Resource(tmpl.AssetNames(),
		func(name string) ([]byte, error) {
			return tmpl.Asset(name)
		})
	ct := boot.ClassifyTemplates(assets)
	layout := ct["layouts"]["app"][0]
	for k, ts := range ct {
		standard := ts["standard"]
		partials := ts["partials"]
		for _, tt := range standard {
			if k != "layouts" {
				ParseAsset(tt, partials, layout, common.Routes, templatesMap)
			}
		}
	}
}

func loadExtTemplates(ac *boot.AppContext) {
	templatesMap := ac.ExtTemplates
	extTemplatesDir := path.Join(*AssetsDir, "templates")
	log.Infof("External templates dir: '%s'\n", extTemplatesDir)
	assets := templatesFileMap(extTemplatesDir)
	ct := boot.ClassifyExtTemplates(extTemplatesDir, assets)
	layoutKey := filepath.Join(extTemplatesDir, "layouts")
	layout := ct[layoutKey]["app"][0]
	for k, ts := range ct {
		standard := ts["standard"]
		partials := ts["partials"]
		for _, tt := range standard {
			if k != "layouts" {
				ParseExtAsset(tt, partials, layout, common.RoutesExt, templatesMap)
			}
		}
	}
}

// ParseAsset - Todo: complete comment.
func ParseAsset(tmplFile string, partials []string, layout string, funcs htmlTemplate.FuncMap, templatesMap map[string]*htmlTemplate.Template) {
	all := make([]string, 10)
	all = append(all, tmplFile)
	all = append(all, partials...)
	all = append(all, layout)
	boot.TrimSlice(&all)
	t, err := htmlTemplate.New(tmplFile, tmpl.Asset).Funcs(funcs).ParseFiles(all...)
	if err != nil {
		log.Warnf("Error parsing asset %s: %e", tmplFile, err)
		panic(err)
	}
	log.Infof("Processed template %s ", tmplFile)
	templatesMap[tmplFile] = t
}

// ParseExtAsset - Todo: complete comment.
func ParseExtAsset(tmplFile string, partials []string, layout string, funcs template.FuncMap, templatesMap map[string]*template.Template) {
	all := make([]string, 10)
	all = append(all, tmplFile)
	all = append(all, partials...)
	all = append(all, layout)
	boot.TrimSlice(&all)
	t, err := template.New(tmplFile).Funcs(funcs).ParseFiles(all...)
	if err != nil {
		log.Warnf("Error parsing asset %s: %e", tmplFile, err)
		panic(err)
	}
	log.Infof("Processed template %s ", tmplFile)
	templatesMap[tmplFile] = t
}
