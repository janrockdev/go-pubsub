package utils

import (
	lr "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

var Logr = &lr.Logger{
	Out:   os.Stdout,
	Level: lr.TraceLevel,
	Formatter: &prefixed.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
		ForceFormatting: true,
	},
}
