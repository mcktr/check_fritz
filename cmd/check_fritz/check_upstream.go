package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mcktr/check_fritz/pkg/fritz"
)

func CheckUpstreamMax(aI ArgumentInformation) {
	soapReq := fritz.NewSoapRequest(aI.Username, aI.Password, aI.Hostname, aI.Port, "/upnp/control/wancommonifconfig1", "WANCommonInterfaceConfig", "X_AVM-DE_GetOnlineMonitor")
	fritz.AddSoapRequestVariable(&soapReq, fritz.NewSoapRequestVariable("NewSyncGroupIndex", "0"))

	err := fritz.DoSoapRequest(&soapReq)

	if HandleError(err) {
		return
	}

	var resp = fritz.GetWANCommonInterfaceOnlineMonitorResponse{}

	err = fritz.HandleSoapRequest(&soapReq, &resp)

	if HandleError(err) {
		return
	}

	upstream, err := strconv.ParseFloat(resp.Newmax_us, 64)

	if HandleError(err) {
		return
	}

	upstream = upstream * 8 / 1000000

	GlobalReturnCode = exitOk

	if internalCheckLower(aI.Warning, upstream) {
		GlobalReturnCode = exitWarning
	}

	if internalCheckLower(aI.Critical, upstream) {
		GlobalReturnCode = exitCritical
	}

	switch GlobalReturnCode {
	case exitOk:
		fmt.Print("OK - Max Upstream: " + fmt.Sprintf("%.2f", upstream) + " Mbit/s \n")
	case exitWarning:
		fmt.Print("WARNING - Max Upstream " + fmt.Sprintf("%.2f", upstream) + " Mbit/s\n")
	case exitCritical:
		fmt.Print("CRITICAL - Max Upstream: " + fmt.Sprintf("%.2f", upstream) + " Mbit/s \n")
	default:
		GlobalReturnCode = exitUnknown
		fmt.Print("UNKNWON - Not able to calculate maximum upstream\n")
	}
}

func CheckUpstreamCurrent(aI ArgumentInformation) {
	soapReq := fritz.NewSoapRequest(aI.Username, aI.Password, aI.Hostname, aI.Port, "/upnp/control/wancommonifconfig1", "WANCommonInterfaceConfig", "X_AVM-DE_GetOnlineMonitor")
	fritz.AddSoapRequestVariable(&soapReq, fritz.NewSoapRequestVariable("NewSyncGroupIndex", "0"))

	err := fritz.DoSoapRequest(&soapReq)

	if HandleError(err) {
		return
	}

	var resp = fritz.GetWANCommonInterfaceOnlineMonitorResponse{}

	err = fritz.HandleSoapRequest(&soapReq, &resp)

	if HandleError(err) {
		return
	}

	upstreamWithHistory := strings.Split(resp.Newus_current_bps, ",")

	upstream, err := strconv.ParseFloat(upstreamWithHistory[0], 32)

	if HandleError(err) {
		return
	}

	upstream = upstream * 8 / 1000000

	GlobalReturnCode = exitOk

	if internalCheckUpper(aI.Warning, upstream) {
		GlobalReturnCode = exitWarning
	}

	if internalCheckUpper(aI.Critical, upstream) {
		GlobalReturnCode = exitCritical
	}

	switch GlobalReturnCode {
	case exitOk:
		fmt.Print("OK - Current Upstream: " + fmt.Sprintf("%.2f", upstream) + " Mbit/s \n")
	case exitWarning:
		fmt.Print("WARNING - Current Upstream " + fmt.Sprintf("%.2f", upstream) + " Mbit/s\n")
	case exitCritical:
		fmt.Print("CRITICAL - Current Upstream: " + fmt.Sprintf("%.2f", upstream) + " Mbit/s \n")
	default:
		GlobalReturnCode = exitUnknown
		fmt.Print("UNKNWON - Not able to calculate current upstream\n")
	}
}