package tester_service

import (
	"github.com/nurislam03/postoffice/pkg/client"
	"github.com/spf13/viper"
	"net/http"
)

type ServiceClient struct{}

func NewServiceClient() *ServiceClient {
	return &ServiceClient{}
}

// TesterServiceClient ...
func TesterServiceClient(verb string, path string, data interface{}, headers http.Header) (map[string]interface{}, *http.Response, error) {
	host := viper.GetString("TESTER_SERVICE_HOST")
	if headers == nil {
		headers = http.Header{}
	}
	headers.Set("Content-Type", "application/json")
	return client.SendRequest(verb, host+path, data, headers)
}
