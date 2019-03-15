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

	uptime, err := strconv.Atoi(resp.NewUpTime)

	if HandleError(err) {
		return
	}

	days := uptime / 86400
	hours := (uptime / 3600) - (days * 24)
	minutes := (uptime / 60) - (days * 1440) - (hours * 60)
	seconds := uptime % 60
	output := fmt.Sprintf("%dd %dh %dm %ds", days, hours, minutes, seconds)

	perfData := perfdata.CreatePerformanceData("uptime", float64(uptime), "s")

	fmt.Print("OK - Device Uptime: " + fmt.Sprintf("%d", uptime) + " seconds (" + output + ") " + perfData.GetPerformanceDataAsString() + "\n")

	GlobalReturnCode = exitOk
}

// CheckDeviceUpdate checks if a new firmware is available
func CheckDeviceUpdate(aI ArgumentInformation) {
	soapReq := fritz.NewSoapRequest(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/userif", "UserInterface", "GetInfo")

	err := fritz.DoSoapRequest(&soapReq)

	if HandleError(err) {
		return
	}

	var resp = fritz.GetInterfaceInfoResponse{}

	err = fritz.HandleSoapRequest(&soapReq, &resp)

	if HandleError(err) {
		return
	}

	state, err := strconv.Atoi(resp.NewUpgradeAvailable)

	if HandleError(err) {
		return
	}

	GlobalReturnCode = exitOk

	if state == 0 {
		GlobalReturnCode = exitOk

		fmt.Print("OK - No update avaiable\n")
	} else {
		GlobalReturnCode = exitCritical

		fmt.Print("CRITICAL - Update available\n")
	}
}
