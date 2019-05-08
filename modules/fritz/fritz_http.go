package fritz

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

// DoSoapRequest does two request to authenticate and handle the SOAP request
func DoSoapRequest(fSR *SoapRequest) error {
	internalCreateNewSoapClient(fSR)
	internalPrepareHTTPClient(fSR)
	internalNewSoapRequestBody(fSR)

	req, err := http.NewRequest("POST", fSR.URL, bytes.NewBuffer([]byte(fSR.soapRequestBody)))

	if err != nil {
		return err
	}

	fSR.soapRequest = *req

	internalPrepareRequestHeader(fSR, fSR.Action, fSR.Service)

	// first request is always unauthenticated due to how digest authenticatioin works

	resp, err := fSR.soapClient.Do(&fSR.soapRequest)

	if err != nil {
		return err
	}

	fSR.soapResponse = *resp

	resp.Body.Close()

	// create immediately a new request (execution later)

	req, err = http.NewRequest("POST", fSR.URL, bytes.NewBuffer([]byte(fSR.soapRequestBody)))

	if err != nil {
		return err
	}

	fSR.soapRequest = *req

	internalPrepareRequestHeader(fSR, fSR.Action, fSR.Service)

	if fSR.soapResponse.StatusCode == http.StatusUnauthorized {
		DoDigestAuthentication(fSR)
	}

	if err != nil {
		return err
	}

	// second request is authenticated

	resp, err = fSR.soapClient.Do(&fSR.soapRequest)

	if err != nil {
		return err
	}

	fSR.soapResponse = *resp

	if fSR.soapResponse.StatusCode == http.StatusUnauthorized {
		error := errors.New("Unauthorized: wrong username or password")
		return error
	}

	return nil
}

// HandleSoapRequest handles the SOAP response
func HandleSoapRequest(fSR *SoapRequest, response Response) error {
	body, err := ioutil.ReadAll(fSR.soapResponse.Body)

	if err != nil {
		panic(err)
	}

	err = xml.Unmarshal(body, &response)

	if err != nil {
		return err
	}

	// Enable for development to output response body
	// fmt.Println(string(body))

	return nil
}

func internalCreateNewSoapClient(fSR *SoapRequest) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	fSR.soapClient = http.Client{Transport: tr}
}

func internalNewSoapRequestBody(SoapRequest *SoapRequest) {
	var request bytes.Buffer

	request.WriteString("<?xml version='1.0'?>\n")
	request.WriteString("<s:Envelope xmlns:s='http://schemas.xmlsoap.org/soap/envelope/' s:encodingStyle='http://schemas.xmlsoap.org/soap/encoding/'>\n")
	request.WriteString("<s:Body>\n")
	request.WriteString("<u:" + SoapRequest.Action + " xmlns:u='urn:dslforum-org:service:" + SoapRequest.Service + ":1'>\n")

	if (SoapRequest.XMLVariable != SoapRequestVariable{}) {
		request.WriteString("<" + SoapRequest.XMLVariable.Name + ">\n")
		request.WriteString(SoapRequest.XMLVariable.Value)
		request.WriteString("</" + SoapRequest.XMLVariable.Name + ">\n")
	}

	request.WriteString("</u:" + SoapRequest.Action + ">\n")
	request.WriteString("</s:Body>\n")
	request.WriteString("</s:Envelope>")

	SoapRequest.soapRequestBody = request.String()
}

func internalPrepareHTTPClient(fSR *SoapRequest) {
	// TODO: Remove ?
}

func internalPrepareRequestHeader(fSR *SoapRequest, action string, service string) {
	fSR.soapRequest.Header.Set("Content-Type", "text/xml; charset='utf-8'")
	fSR.soapRequest.Header.Set("SoapAction", "urn:dslforum-org:service:"+service+":1#"+action)
}
