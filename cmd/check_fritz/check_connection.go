package main

import (
	"fmt"
	"strconv"

	"github.com/mcktr/check_fritz/modules/perfdata"

	"github.com/mcktr/check_fritz/modules/fritz"
)

// CheckConnectionStatus checks the internet connection status
func CheckConnectionStatus(aI ArgumentInformation) {
	soapReq := fritz.NewSoapRequest(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/wanpppconn1", "WANPPPConnection", "GetInfo")

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

		fmt.Print("OK - Connection Status: " + resp.NewConnectionStatus + "; External IP: " + resp.NewExternalIPAddress + "\n")

		GlobalReturnCode = exitOk
	} else {
		fmt.Print("CRITICAL - Connection Status: " + resp.NewConnectionStatus + "\n")

		GlobalReturnCode = exitCritical
	}
}

// CheckConnectionUptime checks the uptime of the internet connection
func CheckConnectionUptime(aI ArgumentInformation) {
	soapReq := fritz.NewSoapRequest(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/wanpppconn1", "WANPPPConnection", "GetInfo")

	err := fritz.DoSoapRequest(&soapReq)

	if HandleError(err) {
		return
	}

	var resp = fritz.GetWANPPPConnectionInfoResponse{}

	err = fritz.HandleSoapRequest(&soapReq, &resp)

	if HandleError(err) {
		return
	}

	uptime, err := strconv.Atoi(resp.NewUptime)

	if HandleError(err) {
		return
	}

	days := uptime / 86400
	hours := (uptime / 3600) - (days * 24)
	minutes := (uptime / 60) - (days * 1440) - (hours * 60)
	seconds := uptime % 60
	output := fmt.Sprintf("%dd %dh %dm %ds", days, hours, minutes, seconds)
	perfData := perfdata.CreatePerformanceData("uptime", float64(uptime), "s")

	fmt.Print("OK - Connection Uptime: " + fmt.Sprintf("%d", uptime) + " seconds (" + output + ") " + perfData.GetPerformanceDataAsString() + "\n")

	GlobalReturnCode = exitOk
}
