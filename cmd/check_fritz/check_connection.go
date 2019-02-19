package main

import (
	"fmt"
	"strconv"

	"github.com/mcktr/check_fritz/pkg/perfdata"

	"github.com/mcktr/check_fritz/pkg/fritz"
)

// CheckConnectionStatus checks the internet connection status
func CheckConnectionStatus(aI ArgumentInformation) {
	soapReq := fritz.NewSoapRequest(aI.Username, aI.Password, aI.Hostname, aI.Port, "/upnp/control/wanpppconn1", "WANPPPConnection", "GetInfo")

	err := fritz.DoSoapRequest(&soapReq)

	if HandleError(err) {
		return
	}

	var resp = fritz.GetWANPPPConnectionInfoResponse{}

	err = fritz.HandleSoapRequest(&soapReq, &resp)

	if HandleError(err) {
		return
	}

	if resp.NewConnectionStatus == "Connected" {
		fmt.Print("OK - Connection Status: " + resp.NewConnectionStatus + "\n")

		GlobalReturnCode = exitOk
	} else {
		fmt.Print("CRITICAL - Connection Status: " + resp.NewConnectionStatus + "\n")

		GlobalReturnCode = exitCritical
	}
}

// CheckConnectionUptime checks the uptime of the internet connection
func CheckConnectionUptime(aI ArgumentInformation) {
	soapReq := fritz.NewSoapRequest(aI.Username, aI.Password, aI.Hostname, aI.Port, "/upnp/control/wanpppconn1", "WANPPPConnection", "GetInfo")

	err := fritz.DoSoapRequest(&soapReq)

	if HandleError(err) {
		return
	}

	var resp = fritz.GetWANPPPConnectionInfoResponse{}

	err = fritz.HandleSoapRequest(&soapReq, &resp)

	if HandleError(err) {
		return
	}

	uptime, err := strconv.ParseFloat(resp.NewUptime, 64)

	if HandleError(err) {
		return
	}

	perfData := perfdata.CreatePerformanceData("uptime", uptime, "s")

	fmt.Print("OK - Connection Uptime: " + fmt.Sprintf("%.f", uptime) + " " + perfData.GetPerformanceDataAsString() + "\n")

	GlobalReturnCode = exitOk
}
