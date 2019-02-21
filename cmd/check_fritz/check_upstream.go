package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mcktr/check_fritz/pkg/fritz"
	"github.com/mcktr/check_fritz/pkg/perfdata"
	"github.com/mcktr/check_fritz/pkg/thresholds"
)

// CheckUpstreamMax checks the maximum upstream that is available on this internet connection
func CheckUpstreamMax(aI ArgumentInformation) {
	soapReq := fritz.NewSoapRequest(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/wancommonifconfig1", "WANCommonInterfaceConfig", "X_AVM-DE_GetOnlineMonitor")
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

	upstream, err := strconv.ParseFloat(resp.NewMaxUS, 64)

	if HandleError(err) {
		return
	}

	upstream = upstream * 8 / 1000000
	perfData := perfdata.CreatePerformanceData("upstream_max", upstream, "")

	GlobalReturnCode = exitOk

	if thresholds.IsSet(aI.Warning) {
		perfData.SetWarning(*aI.Warning)

		if thresholds.CheckLower(*aI.Warning, upstream) {
			GlobalReturnCode = exitWarning
		}
	}

	if thresholds.IsSet(aI.Critical) {
		perfData.SetCritical(*aI.Critical)

		if thresholds.CheckLower(*aI.Critical, upstream) {
			GlobalReturnCode = exitCritical
		}
	}

	output := " - Max Upstream: " + fmt.Sprintf("%.2f", upstream) + " Mbit/s " + perfData.GetPerformanceDataAsString()

	switch GlobalReturnCode {
	case exitOk:
		fmt.Print("OK" + output + "\n")
	case exitWarning:
		fmt.Print("WARNING" + output + "\n")
	case exitCritical:
		fmt.Print("CRITICAL" + output + "\n")
	default:
		GlobalReturnCode = exitUnknown
		fmt.Print("UNKNWON - Not able to calculate maximum upstream\n")
	}
}

// CheckUpstreamCurrent checks the current used upstream
func CheckUpstreamCurrent(aI ArgumentInformation) {
	soapReq := fritz.NewSoapRequest(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/wancommonifconfig1", "WANCommonInterfaceConfig", "X_AVM-DE_GetOnlineMonitor")
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

	upstreamWithHistory := strings.Split(resp.NewUSCurrentBPS, ",")

	upstream, err := strconv.ParseFloat(upstreamWithHistory[0], 32)

	if HandleError(err) {
		return
	}

	upstream = upstream * 8 / 1000000
	perfData := perfdata.CreatePerformanceData("upstream_current", upstream, "")

	GlobalReturnCode = exitOk

	if thresholds.IsSet(aI.Warning) {
		perfData.SetWarning(*aI.Warning)

		if thresholds.CheckUpper(*aI.Warning, upstream) {
			GlobalReturnCode = exitWarning
		}
	}

	if thresholds.IsSet(aI.Critical) {
		perfData.SetCritical(*aI.Critical)

		if thresholds.CheckUpper(*aI.Critical, upstream) {
			GlobalReturnCode = exitCritical
		}
	}

	output := " - Current Upstream: " + fmt.Sprintf("%.2f", upstream) + " Mbit/s " + perfData.GetPerformanceDataAsString()

	switch GlobalReturnCode {
	case exitOk:
		fmt.Print("OK" + output + "\n")
	case exitWarning:
		fmt.Print("WARNING" + output + "\n")
	case exitCritical:
		fmt.Print("CRITICAL" + output + "\n")
	default:
		GlobalReturnCode = exitUnknown
		fmt.Print("UNKNWON - Not able to calculate current upstream\n")
	}
}
