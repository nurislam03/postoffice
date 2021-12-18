package services

import (
	"github.com/nurislam03/postoffice/pkg/testerclient"
	"github.com/spf13/viper"
	"sync"
	"time"
)

var testerClient *testerclient.Client
var testerOnce sync.Once

// TesterService ...
func TesterService() *testerclient.Client {
	testerOnce.Do(func() {
		testerClient = testerclient.NewClient(testerclient.Options{
			Host:     viper.GetString("TESTER_SERVICE_HOST"),
			Port:     viper.GetInt("TESTER_SERVICE_PORT"),
			Secure:   viper.GetBool("TESTER_SERVICE_SECURE"),
			BasePath: viper.GetString("TESTER_SERVICE_BASEPATH"),
			TimeOut:  time.Second*10,
		})
	})
	return testerClient
}