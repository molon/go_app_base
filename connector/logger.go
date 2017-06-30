package connector

import (
	"github.com/molon/go_app_base/internal/logger"
)

// Debugf logs a debug statement
func (a *Connector) Debugf(format string, v ...interface{}) {
	a.logger.Logf(logger.DEBUG, format, v...)
}

// Infof logs a info statement
func (a *Connector) Infof(format string, v ...interface{}) {
	a.logger.Logf(logger.INFO, format, v...)
}

// Warnf logs a warn statement
func (a *Connector) Warnf(format string, v ...interface{}) {
	a.logger.Logf(logger.WARN, format, v...)
}

// Errorf logs a error statement
func (a *Connector) Errorf(format string, v ...interface{}) {
	a.logger.Logf(logger.ERROR, format, v...)
}

// Fatalf logs a fatal statement
func (a *Connector) Fatalf(format string, v ...interface{}) {
	a.logger.Logf(logger.FATAL, format, v...)
}
