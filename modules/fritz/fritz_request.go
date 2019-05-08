package fritz

import "net/http"

// SoapRequest is the data structure for a SOAP request including the response
type SoapRequest struct {
	Username         string
	Password         string
	URL              string
	URLPath          string
	Action           string
	Service          string
	XMLVariable      SoapRequestVariable
	soapClient       http.Client
	soapRequest      http.Request
	soapRequestBody  string
	soapResponse     http.Response
	soapResponseBody string
	soapDigestAuth   map[string]string
}

// SoapRequestVariable is the data structure for a variable that can injected in the SOAP request
type SoapRequestVariable struct {
	Name  string
	Value string
}

// NewSoapRequest creates a new FritzSoapRequest structure
func NewSoapRequest(username string, password string, hostname string, port string, urlpath string, service string, action string) SoapRequest {
	var fSR SoapRequest

	fSR.URL = "https://" + hostname + ":" + port + urlpath

	fSR.URLPath = urlpath
	fSR.Username = username
	fSR.Password = password
	fSR.Action = action
	fSR.Service = service

	return fSR
}

// NewSoapRequestVariable creates a new SoapRequestVariable
func NewSoapRequestVariable(name string, value string) *SoapRequestVariable {
	var sRV SoapRequestVariable

	sRV.Name = name
	sRV.Value = value

	return &sRV
}

// AddSoapRequestVariable adds a SoapRequestVariable to a SoapRequest
func AddSoapRequestVariable(soapRequest *SoapRequest, soapRequestVariable *SoapRequestVariable) {
	soapRequest.XMLVariable = *soapRequestVariable
}
