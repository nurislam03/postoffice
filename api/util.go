package api

import (
	"encoding/base64"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func b64decode(str string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// SetLogLevel sets Logger level ...
func SetLogLevel() {
	logLevel := viper.GetString("LOGRUS_LOGLEVEL")

	switch logLevel {
	case "Panic":
		logger.SetLevel(logrus.DebugLevel)
	case "Fatal":
		logger.SetLevel(logrus.FatalLevel)
	case "Error":
		logger.SetLevel(logrus.ErrorLevel)
	case "Warn":
		logger.SetLevel(logrus.WarnLevel)
	case "Info":
		logger.SetLevel(logrus.InfoLevel)
	case "Debug":
		logger.SetLevel(logrus.DebugLevel)
	case "Trace":
		logger.SetLevel(logrus.TraceLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}
}
