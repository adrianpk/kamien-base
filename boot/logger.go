/**
 * Copyright (c) 2018 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */
package boot

import (
	"github.com/adrianpk/kamien"
)

func initLogger() {
	logLevel := Configuration.GetLogLevel()
	if logLevel == "" {
		logLevel = kamien.ErrorLogLevel
	}
	log = kamien.GetLogger(logLevel, Env)
}
