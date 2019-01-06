package fritz

import "net/http"

type SoapRequest struct {
	Username         string
	Password         string
	URL              string
	Action           string
	Service          string
	soapClient       http.Client
	soapRequest      http.Request
	soapRequestBody  string
	soapResponse     http.Response
	soapResponseBody string
	soapDigestAuth   map[string]string
}

// NewSoapRequest creates a new FritzSoapRequest structure
func NewSoapRequest(url string, username string, password string, service string, action string) SoapRequest {
	var fSR SoapRequest

	fSR.URL = url
	fSR.Username = username
	fSR.Password = password
	fSR.Action = action
	fSR.Service = service

	return fSR
}
