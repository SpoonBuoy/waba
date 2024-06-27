package middleware

import (
	"os"
	"sync"

	gl "github.com/ahsmha/gashtools/logger"
)

var (
	_loggerOnce sync.Once
)

func InitLogger() gl.IGashaLogger {
	var logger gl.IGashaLogger
	_loggerOnce.Do(func() {
		os.Setenv("LOGLEVEL", "trace")
		logger = gl.NewGLogger(gl.WithAccessLogOptions())
	})
	return logger
}
