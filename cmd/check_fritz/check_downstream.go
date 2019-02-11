package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mcktr/check_fritz/pkg/fritz"
	"github.com/mcktr/check_fritz/pkg/thresholds"
)

// CheckDownstreamMax checks the maximum downstream that is available on this internet connection
func CheckDownstreamMax(aI ArgumentInformation) {
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

	downstream, err := strconv.ParseFloat(resp.NewMaxDS, 64)

	if HandleError(err) {
		return
	}

	downstream = downstream * 8 / 1000000

	GlobalReturnCode = exitOk

	if thresholds.CheckLower(aI.Warning, downstream) {
		GlobalReturnCode = exitWarning
	}

	if thresholds.CheckLower(aI.Critical, downstream) {
		GlobalReturnCode = exitCritical
	}

	switch GlobalReturnCode {
	case exitOk:
		fmt.Print("OK - Max Downstream: " + fmt.Sprintf("%.2f", downstream) + " Mbit/s \n")
	case exitWarning:
		fmt.Print("WARNING - Max Downstream " + fmt.Sprintf("%.2f", downstream) + " Mbit/s\n")
	case exitCritical:
		fmt.Print("CRITICAL - Max Downstream: " + fmt.Sprintf("%.2f", downstream) + " Mbit/s \n")
	default:
		GlobalReturnCode = exitUnknown
		fmt.Print("UNKNWON - Not able to calculate maximum downstream\n")
	}
}

// CheckDownstreamCurrent checks the current used downstream
func CheckDownstreamCurrent(aI ArgumentInformation) {
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

	downstreamWithHistory := strings.Split(resp.NewDSCurrentBPS, ",")

	downstream, err := strconv.ParseFloat(downstreamWithHistory[0], 64)

	if HandleError(err) {
		return
	}

	downstream = downstream * 8 / 1000000

	GlobalReturnCode = exitOk

	if thresholds.CheckUpper(aI.Warning, downstream) {
		GlobalReturnCode = exitWarning
	}

	if thresholds.CheckUpper(aI.Critical, downstream) {
		GlobalReturnCode = exitCritical
	}

	switch GlobalReturnCode {
	case exitOk:
		fmt.Print("OK - Current Downstream: " + fmt.Sprintf("%.2f", downstream) + " Mbit/s \n")
	case exitWarning:
		fmt.Print("WARNING - Current Downstream " + fmt.Sprintf("%.2f", downstream) + " Mbit/s\n")
	case exitCritical:
		fmt.Print("CRITICAL - Current Downstream: " + fmt.Sprintf("%.2f", downstream) + " Mbit/s \n")
	default:
		GlobalReturnCode = exitUnknown
		fmt.Print("UNKNWON - Not able to calculate current downstream\n")
	}
}
