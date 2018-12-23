package test

import (
	"fmt"
	"strings"

	"{{.Package}}/fixtures"
	bindata "github.com/golang-migrate/migrate/source/go_bindata"
	"github.com/romanyx/polluter"
	log "github.com/siddontang/go/log"
)

// ProcessFixtures - Todo: complete comment.
func (e *Environment) processFixtures() {
	assets := bindata.Resource(fixtures.AssetNames(),
		func(name string) ([]byte, error) {
			return fixtures.Asset(name)
		})
	for _, an := range assets.Names {
		data, err := fixtures.Asset(an)
		if err != nil {
			log.Debugf("Error loading asset %s\n", an)
		} else {
			fmt.Printf("Processing %s\n", an)
			e.processAsset(data)
			log.Debugf("%s processed.\n", an)
		}
		// log.Debugf("Fixture processing completed.")
	}
}

func (e *Environment) processAsset(assetData []byte) {
	db := e.GetDBInstance()
	p := polluter.New(polluter.PostgresEngine(db), polluter.YAMLParser)
	s := string(assetData)
	err := p.Pollute(strings.NewReader(s))
	if err != nil {
		log.Infof("Cannot seed the database: %s\n", err)
	}
}
