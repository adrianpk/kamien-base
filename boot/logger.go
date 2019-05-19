/**
 * Copyright (c) 2018 Adrian K <adrian.git@kuguar.dev>
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
