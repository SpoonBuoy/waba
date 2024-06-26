package middleware

import (
	"os"

	gl "github.com/ahsmha/gashtools/logger"
)

var Logger gl.IGashaLogger

func InitLogger() {
	os.Setenv("LOGLEVEL", "trace")
	Logger = gl.NewGLogger(gl.WithAccessLogOptions())
}
