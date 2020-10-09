package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mcktr/check_fritz/modules/fritz"
	"github.com/mcktr/check_fritz/modules/perfdata"
	"github.com/mcktr/check_fritz/modules/thresholds"
)

// CheckUpstreamMax checks the maximum upstream that is available on this internet connection
func CheckUpstreamMax(aI ArgumentInformation) {
	resps := make(chan []byte)
	errs := make(chan error)

	var soapReq fritz.SoapData

	isDSL := false

	if strings.ToLower(*aI.Modelgroup) == "dsl" {
		isDSL = true
	}

	if isDSL {
		soapReq = fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/wandslifconfig1", "WANDSLInterfaceConfig", "GetInfo")
	} else {
		soapReq = fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/wancommonifconfig1", "WANCommonInterfaceConfig", "GetCommonLinkProperties")
	}

	go fritz.DoSoapRequest(&soapReq, resps, errs, aI.Debug)

	res, err := fritz.ProcessSoapResponse(resps, errs, 1, *aI.Timeout)

	if err != nil {
		fmt.Printf("UNKNOWN - %s\n", err)
		return
	}

	var upstream float64

	if isDSL {
		soapResp := fritz.WANDSLInterfaceGetInfoResponse{}
		err = fritz.UnmarshalSoapResponse(&soapResp, res)

		if err != nil {
			panic(err)
		}

		ups, err := strconv.ParseFloat(soapResp.NewUpstreamCurrRate, 64)

		if err != nil {
			panic(err)
		}

		upstream = ups / 1000
	} else {
		soapResp := fritz.WANCommonInterfaceCommonLinkPropertiesResponse{}
		err = fritz.UnmarshalSoapResponse(&soapResp, res)

		if err != nil {
			panic(err)
		}

		ups, err := strconv.ParseFloat(soapResp.NewLayer1UpstreamMaxBitRate, 64)

		if err != nil {
			panic(err)
		}

		upstream = ups / 1000000
	}

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
	resps := make(chan []byte)
	errs := make(chan error)

	soapReq := fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/wancommonifconfig1", "WANCommonInterfaceConfig", "X_AVM-DE_GetOnlineMonitor")
	soapReq.AddSoapDataVariable(fritz.CreateNewSoapVariable("NewSyncGroupIndex", "0"))
	go fritz.DoSoapRequest(&soapReq, resps, errs, aI.Debug)

	res, err := fritz.ProcessSoapResponse(resps, errs, 1, *aI.Timeout)

	if err != nil {
		fmt.Printf("UNKNOWN - %s\n", err)
		return
	}

	soapResp := fritz.WANCommonInterfaceOnlineMonitorResponse{}
	err = fritz.UnmarshalSoapResponse(&soapResp, res)

	upstreamWithHistory := strings.Split(soapResp.NewUSCurrentBPS, ",")

	upstream, err := strconv.ParseFloat(upstreamWithHistory[0], 32)

	if err != nil {
		panic(err)
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

// CheckUpstreamUsage checks the total and current upstream and calculates the utilization
func CheckUpstreamUsage(aI ArgumentInformation) {
	resps := make(chan []byte)
	errs := make(chan error)

	soapReq := fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/wancommonifconfig1", "WANCommonInterfaceConfig", "X_AVM-DE_GetOnlineMonitor")
	soapReq.AddSoapDataVariable(fritz.CreateNewSoapVariable("NewSyncGroupIndex", "0"))
	go fritz.DoSoapRequest(&soapReq, resps, errs, aI.Debug)

	res, err := fritz.ProcessSoapResponse(resps, errs, 1, *aI.Timeout)

	if err != nil {
		fmt.Printf("UNKNOWN - %s\n", err)
		return
	}

	soapResp := fritz.WANCommonInterfaceOnlineMonitorResponse{}
	err = fritz.UnmarshalSoapResponse(&soapResp, res)

	upstreamWithHistory := strings.Split(soapResp.NewUSCurrentBPS, ",")
	upstreamCurrent, err := strconv.ParseFloat(upstreamWithHistory[0], 64)

	if err != nil {
		panic(err)
	}

	isDSL := false

	if strings.ToLower(*aI.Modelgroup) == "dsl" {
		isDSL = true
	}

	if isDSL {
		soapReq = fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/wandslifconfig1", "WANDSLInterfaceConfig", "GetInfo")
	} else {
		soapReq = fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/wancommonifconfig1", "WANCommonInterfaceConfig", "GetCommonLinkProperties")
	}

	go fritz.DoSoapRequest(&soapReq, resps, errs, aI.Debug)

	res, err = fritz.ProcessSoapResponse(resps, errs, 1, *aI.Timeout)

	if err != nil {
		fmt.Printf("UNKNOWN - %s\n", err)
		return
	}

	var upstreamMax float64

	if isDSL {
		soapResp := fritz.WANDSLInterfaceGetInfoResponse{}
		err = fritz.UnmarshalSoapResponse(&soapResp, res)

		if err != nil {
			panic(err)
		}

		ups, err := strconv.ParseFloat(soapResp.NewUpstreamCurrRate, 64)

		if err != nil {
			panic(err)
		}

		upstreamMax = ups / 1000
	} else {
		soapResp := fritz.WANCommonInterfaceCommonLinkPropertiesResponse{}
		err = fritz.UnmarshalSoapResponse(&soapResp, res)

		if err != nil {
			panic(err)
		}

		ups, err := strconv.ParseFloat(soapResp.NewLayer1UpstreamMaxBitRate, 64)

		if err != nil {
			panic(err)
		}

		upstreamMax = ups / 1000000
	}

	upstreamCurrent = upstreamCurrent * 8 / 1000000

	if upstreamMax == 0 {
		fmt.Printf("UNKNOWN - Maximum Downstream is 0\n")
		return
	}

	upstreamUsage := 100 / upstreamMax * upstreamCurrent
	perfData := perfdata.CreatePerformanceData("upstream_usage", upstreamUsage, "")

	perfData.SetMinimum(0.0)
	perfData.SetMaximum(100.0)

	GlobalReturnCode = exitOk

	if thresholds.IsSet(aI.Warning) {
		perfData.SetWarning(*aI.Warning)

		if thresholds.CheckUpper(*aI.Warning, upstreamUsage) {
			GlobalReturnCode = exitWarning
		}
	}

	if thresholds.IsSet(aI.Critical) {
		perfData.SetCritical(*aI.Critical)

		if thresholds.CheckUpper(*aI.Critical, upstreamUsage) {
			GlobalReturnCode = exitCritical
		}
	}

	output := " - " + fmt.Sprintf("%.2f", upstreamUsage) + "% Upstream utilization (" + fmt.Sprintf("%.2f", upstreamCurrent) + " Mbit/s of " + fmt.Sprintf("%.2f", upstreamMax) + " Mbits) " + perfData.GetPerformanceDataAsString()

	switch GlobalReturnCode {
	case exitOk:
		fmt.Print("OK" + output + "\n")
	case exitWarning:
		fmt.Print("WARNING" + output + "\n")
	case exitCritical:
		fmt.Print("CRITICAL" + output + "\n")
	default:
		GlobalReturnCode = exitUnknown
		fmt.Print("UNKNWON - Not able to calculate upstream utilization\n")
	}
}
