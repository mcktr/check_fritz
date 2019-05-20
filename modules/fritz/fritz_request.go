package fritz

// SoapData is the data structure for a SOAP request including the response
type SoapData struct {
	Username    []byte
	Password    []byte
	URL         string
	URLPath     string
	Action      string
	Service     string
	XMLVariable SoapDataVariable
}

// SoapDataVariable is the data structure for a variable that can injected in the SOAP request
type SoapDataVariable struct {
	Name  string
	Value string
}

// AddSoapDataVariable adds a SoapRequestVariable to a SoapRequest
func (soapData *SoapData) AddSoapDataVariable(soapDataVariable SoapDataVariable) {
	soapData.XMLVariable = soapDataVariable
}

// CreateNewSoapData creates a new FritzSoapRequest structure
func CreateNewSoapData(username string, password string, hostname string, port string, urlpath string, service string, action string) SoapData {
	var fSR SoapData

	fSR.URL = "https://" + hostname + ":" + port + urlpath

	fSR.URLPath = urlpath
	fSR.Username = []byte(username)
	fSR.Password = []byte(password)
	fSR.Action = action
	fSR.Service = service

	return fSR
}

// CreateNewSoapVariable creates a new SoapRequestVariable
func CreateNewSoapVariable(name string, value string) SoapDataVariable {
	var soapDataVariable SoapDataVariable

	soapDataVariable.Name = name
	soapDataVariable.Value = value

	return soapDataVariable
}
