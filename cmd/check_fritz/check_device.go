package main

import (
	"fmt"

	"github.com/mcktr/check_fritz/pkg/fritz"
)

func CheckDeviceUptime(hostname, port string, username string, password string) {
	soapReq := fritz.NewSoapRequest(username, password, hostname, port, "/upnp/control/deviceinfo", "DeviceInfo", "GetInfo")

	err := fritz.DoSoapRequest(&soapReq)

	if HandleError(err) {
		return
	}

	var resp = fritz.GetDeviceInfoResponse{}

	err = fritz.HandleSoapRequest(&soapReq, &resp)

	if HandleError(err) {
		return
	}

	fmt.Print("OK - Device Uptime: " + resp.NewUpTime + "\n")

	GlobalReturnCode = exitOk
}
