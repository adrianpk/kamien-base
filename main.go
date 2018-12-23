/**
 * Copyright (c) 2018 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package main

//go:generate rm -rf templates/templates.go
//go:generate rm -rf migrations/migrations.go
//go:generate go-bindata -prefix "resources/templates/" -pkg templates -o templates/templates.go resources/templates/...
//go:generate go-bindata -prefix "resources/migrations/" -pkg migrations -o migrations/migrations.go resources/migrations/...
//go:generate go-bindata -prefix "resources/seeds/" -pkg seeds -o seeds/seeds.go resources/seeds/...
//go:generate go-bindata -prefix "resources/fixtures/" -pkg fixtures -o fixtures/fixtures.go resources/fixtures/...

import (
	"github.com/adrianpk/kamien"
	"{{.Package}}/app"
	"{{.Package}}/boot"
)

var (
	log = kamien.Log
)

func main() {
	boot.Init()
	errs := app.InitServer()
	select {
	case err := <-errs:
		log.Printf("Server init error: %s", err)
	}
}
