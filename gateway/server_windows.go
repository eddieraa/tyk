//go:build windows
// +build windows

package gateway

import (
	"github.com/TykTechnologies/tyk/config"
)

func setLoggerSyslog(gwConfig config.Config) {
	mainLog.Debug("Syslog is not available on windows")
}
