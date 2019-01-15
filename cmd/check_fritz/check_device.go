package main

import (
	"fmt"

	"github.com/mcktr/check_fritz/pkg/fritz"
)

func CheckDeviceUptime(aI ArgumentInformation) {
	soapReq := fritz.NewSoapRequest(aI.Username, aI.Password, aI.Hostname, aI.Port, "/upnp/control/deviceinfo", "DeviceInfo", "GetInfo")

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
