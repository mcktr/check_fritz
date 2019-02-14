package main

import (
	"fmt"
	"strconv"

	"github.com/mcktr/check_fritz/pkg/fritz"
	"github.com/mcktr/check_fritz/pkg/perfdata"
)

// CheckDeviceUptime checks the uptime of the device
func CheckDeviceUptime(aI ArgumentInformation) {
	soapReq := fritz.NewSoapRequest(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/deviceinfo", "DeviceInfo", "GetInfo")

	err := fritz.DoSoapRequest(&soapReq)

	if HandleError(err) {
		return
	}

	var resp = fritz.GetDeviceInfoResponse{}

	err = fritz.HandleSoapRequest(&soapReq, &resp)

	if HandleError(err) {
		return
	}

	uptime, err := strconv.ParseFloat(resp.NewUpTime, 64)

	if HandleError(err) {
		return
	}

	perfData := perfdata.CreatePerformanceData("uptime", uptime, "s")

	fmt.Print("OK - Device Uptime: " + fmt.Sprintf("%.0f", uptime) + " " + perfData.GetPerformanceDataAsString() + "\n")

	GlobalReturnCode = exitOk
}
