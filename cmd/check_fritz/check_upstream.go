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
	// Start of the initial request to get the total number of sync groups; we always start with group id 0
	resps := make(chan []byte)
	errs := make(chan error)

	initialSoapReq := fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/wancommonifconfig1", "WANCommonInterfaceConfig", "X_AVM-DE_GetOnlineMonitor")
	initialSoapReq.AddSoapDataVariable(fritz.CreateNewSoapVariable("NewSyncGroupIndex", "0"))
	go fritz.DoSoapRequest(&initialSoapReq, resps, errs, aI.Debug)

	initialResponse, err := fritz.ProcessSoapResponse(resps, errs, 1, *aI.Timeout)

	if err != nil {
		fmt.Printf("UNKNOWN - %s\n", err)
		return
	}

	initialSoapResp := fritz.WANCommonInterfaceOnlineMonitorResponse{}
	err = fritz.UnmarshalSoapResponse(&initialSoapResp, initialResponse)

	// Query the total number of sync groups
	totalSyncGroups, err := strconv.Atoi(initialSoapResp.NewTotalNumberSyncGroups)
	if err != nil {
		fmt.Printf("UNKOWN - %s\n", err)
		return
	}

	foundSupportedSyncMode := false
	foundSyncMode := make([]string, 0)
	var finalSoapResponse *fritz.WANCommonInterfaceOnlineMonitorResponse

	for currentSyncGroup := 0; currentSyncGroup < totalSyncGroups; currentSyncGroup++ {
		soapResp := fritz.WANCommonInterfaceOnlineMonitorResponse{}

		// We only need to perform additional queries when there are more than 1 sync groups
		if totalSyncGroups > 1 {
			// Start of additional query attempts (depending on how many sync groups are found)
			responseChan := make(chan []byte)
			errorChan := make(chan error)

			soapReq := fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/wancommonifconfig1", "WANCommonInterfaceConfig", "X_AVM-DE_GetOnlineMonitor")
			soapReq.AddSoapDataVariable(fritz.CreateNewSoapVariable("NewSyncGroupIndex", strconv.Itoa(currentSyncGroup)))
			go fritz.DoSoapRequest(&soapReq, responseChan, errorChan, aI.Debug)

			resp, err := fritz.ProcessSoapResponse(responseChan, errorChan, 1, *aI.Timeout)
			if err != nil {
				fmt.Printf("UNKNOWN - %s\n", err)
				return
			}

			err = fritz.UnmarshalSoapResponse(&soapResp, resp)
			if err != nil {
				fmt.Printf("UNKNOWN - %s\n", err)
				return
			}
		} else {
			soapResp = initialSoapResp
		}

		syncGroupMode := soapResp.NewSyncGroupMode
		foundSyncMode = append(foundSyncMode, syncGroupMode)

		// Search for supported sync groups
		if syncGroupMode == "VDSL" || syncGroupMode == "CABLE" {
			foundSupportedSyncMode = true
			finalSoapResponse = &soapResp

			break
		}
	}

	var upstream float64

	if foundSupportedSyncMode {
		upstream, err = strconv.ParseFloat(finalSoapResponse.NewMaxUS, 64)
		if err != nil {
			fmt.Printf("UNKNOWN - %s\n", err)
			return
		}
	} else {
		fmt.Printf("UNKNOWN - Could not find a supported SyncGroup (VDSL or CABLE); found the following: %s\n", strings.Join(foundSyncMode, ", "))
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

	upstreamMax, err := strconv.ParseFloat(soapResp.NewMaxUS, 64)

	if err != nil {
		panic(err)
	}

	upstreamCurrent = upstreamCurrent * 8 / 1000000
	upstreamMax = upstreamMax * 8 / 1000000

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
