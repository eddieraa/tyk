//go:build linux
// +build linux

package gateway

import (
	"log/syslog"

	"github.com/TykTechnologies/tyk/config"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
)

func setLoggerSyslog(gwConfig config.Config) {
	mainLog.Debug("Enabling Syslog support")
	hook, err := logrus_syslog.NewSyslogHook(gwConfig.SyslogTransport,
		gwConfig.SyslogNetworkAddr,
		syslog.LOG_INFO, "")

	if err == nil {
		log.Hooks.Add(hook)
		rawLog.Hooks.Add(hook)
	}
	mainLog.Debug("Syslog hook active")
}
