package repo

import (
	"encoding/json"

	"github.com/adrianpk/kamien"
	"{{.Package}}/boot"
	"github.com/markbates/pop/nulls"
)

var (
	log *kamien.Logger
)

func init() {
	initLogger()
}

func initLogger() {
	logLevel := boot.Configuration.GetLogLevel()
	log = kamien.GetLogger(logLevel, boot.Env)
}

func isJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func getDb(err error) {
	if err != nil {
		log.Error(err)
	}
}

func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ToNullsString - Gets the nullable version of a string.
func ToNullsString(str string) nulls.String {
	return nulls.NewString(str)
}
