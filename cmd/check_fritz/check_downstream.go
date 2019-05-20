package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mcktr/check_fritz/modules/fritz"
	"github.com/mcktr/check_fritz/modules/perfdata"
	"github.com/mcktr/check_fritz/modules/thresholds"
)

// CheckDownstreamMax checks the maximum downstream that is available on this internet connection
func CheckDownstreamMax(aI ArgumentInformation) {
	resps := make(chan []byte)
	errs := make(chan error)

	soapReq := fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/wancommonifconfig1", "WANCommonInterfaceConfig", "X_AVM-DE_GetOnlineMonitor")
	soapReq.AddSoapDataVariable(fritz.CreateNewSoapVariable("NewSyncGroupIndex", "0"))
	go fritz.DoSoapRequest(&soapReq, resps, errs)

	res, err := fritz.ProcessSoapResponse(resps, errs, 1)

	if err != nil {
		panic(err)
	}

	soapResp := fritz.WANCommonInterfaceOnlineMonitorResponse{}
	err = fritz.UnmarshalSoapResponse(&soapResp, res)

	downstream, err := strconv.ParseFloat(soapResp.NewMaxDS, 64)

	if err != nil {
		panic(err)
	}

	downstream = downstream * 8 / 1000000
	perfData := perfdata.CreatePerformanceData("downstream_max", downstream, "")

	GlobalReturnCode = exitOk

	if thresholds.IsSet(aI.Warning) {
		perfData.SetWarning(*aI.Warning)

		if thresholds.CheckLower(*aI.Warning, downstream) {
			GlobalReturnCode = exitWarning
		}
	}

	if thresholds.IsSet(aI.Critical) {
		perfData.SetCritical(*aI.Critical)

		if thresholds.CheckLower(*aI.Critical, downstream) {
			GlobalReturnCode = exitCritical
		}
	}

	output := " - Max Downstream: " + fmt.Sprintf("%.2f", downstream) + " Mbit/s " + perfData.GetPerformanceDataAsString()

	switch GlobalReturnCode {
	case exitOk:
		fmt.Print("OK" + output + "\n")
	case exitWarning:
		fmt.Print("WARNING" + output + "\n")
	case exitCritical:
		fmt.Print("CRITICAL" + output + "\n")
	default:
		GlobalReturnCode = exitUnknown
		fmt.Print("UNKNWON - Not able to calculate maximum downstream\n")
	}
}

// CheckDownstreamCurrent checks the current used downstream
func CheckDownstreamCurrent(aI ArgumentInformation) {
	resps := make(chan []byte)
	errs := make(chan error)

	soapReq := fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/wancommonifconfig1", "WANCommonInterfaceConfig", "X_AVM-DE_GetOnlineMonitor")
	soapReq.AddSoapDataVariable(fritz.CreateNewSoapVariable("NewSyncGroupIndex", "0"))
	go fritz.DoSoapRequest(&soapReq, resps, errs)

	res, err := fritz.ProcessSoapResponse(resps, errs, 1)

	if err != nil {
		panic(err)
	}

	soapResp := fritz.WANCommonInterfaceOnlineMonitorResponse{}
	err = fritz.UnmarshalSoapResponse(&soapResp, res)

	downstreamWithHistory := strings.Split(soapResp.NewDSCurrentBPS, ",")

	downstream, err := strconv.ParseFloat(downstreamWithHistory[0], 64)

	if err != nil {
		panic(err)
	}

	downstream = downstream * 8 / 1000000
	perfData := perfdata.CreatePerformanceData("downstream_current", downstream, "")

	GlobalReturnCode = exitOk

	if thresholds.IsSet(aI.Warning) {
		perfData.SetWarning(*aI.Warning)

		if thresholds.CheckUpper(*aI.Warning, downstream) {
			GlobalReturnCode = exitWarning
		}
	}

	if thresholds.IsSet(aI.Critical) {
		perfData.SetCritical(*aI.Critical)

		if thresholds.CheckUpper(*aI.Critical, downstream) {
			GlobalReturnCode = exitCritical
		}
	}

	output := " - Current Downstream: " + fmt.Sprintf("%.2f", downstream) + " Mbit/s " + perfData.GetPerformanceDataAsString()

	switch GlobalReturnCode {
	case exitOk:
		fmt.Print("OK" + output + "\n")
	case exitWarning:
		fmt.Print("WARNING" + output + "\n")
	case exitCritical:
		fmt.Print("CRITICAL" + output + "\n")
	default:
		GlobalReturnCode = exitUnknown
		fmt.Print("UNKNWON - Not able to calculate current downstream\n")
	}
}
