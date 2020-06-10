package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mcktr/check_fritz/modules/fritzutil"

	"github.com/mcktr/check_fritz/modules/fritz"
	"github.com/mcktr/check_fritz/modules/perfdata"
	"github.com/mcktr/check_fritz/modules/thresholds"
)

// CheckDownstreamMax checks the maximum downstream that is available on this internet connection
func CheckDownstreamMax(aI ArgumentInformation) {
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
		if fritzutil.Contains(supportedSyncGroupModes, syncGroupMode) {
			foundSupportedSyncMode = true
			finalSoapResponse = &soapResp

			break
		}
	}

	var downstream float64

	if foundSupportedSyncMode {
		downstream, err = strconv.ParseFloat(finalSoapResponse.NewMaxDS, 64)
		if err != nil {
			fmt.Printf("UNKNOWN - %s\n", err)
			return
		}
	} else {
		fmt.Printf("UNKNOWN - Could not find a supported SyncGroup (%s); found the following: %s\n", strings.Join(supportedSyncGroupModes, ", "), strings.Join(foundSyncMode, ", "))
		return
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
	go fritz.DoSoapRequest(&soapReq, resps, errs, aI.Debug)

	res, err := fritz.ProcessSoapResponse(resps, errs, 1, *aI.Timeout)

	if err != nil {
		fmt.Printf("UNKNOWN - %s\n", err)
		return
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

// CheckDownstreamUsage checks the current used downstream
func CheckDownstreamUsage(aI ArgumentInformation) {
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

	downstreamWithHistory := strings.Split(soapResp.NewDSCurrentBPS, ",")
	downstreamCurrent, err := strconv.ParseFloat(downstreamWithHistory[0], 64)

	if err != nil {
		panic(err)
	}

	downstreamMax, err := strconv.ParseFloat(soapResp.NewMaxDS, 64)

	if err != nil {
		panic(err)
	}

	downstreamCurrent = downstreamCurrent * 8 / 1000000
	downstreamMax = downstreamMax * 8 / 1000000

	if downstreamMax == 0 {
		fmt.Printf("UNKNOWN - Maximum Downstream is 0\n")
		return
	}

	downstreamUsage := 100 / downstreamMax * downstreamCurrent
	perfData := perfdata.CreatePerformanceData("downstream_usage", downstreamUsage, "")

	perfData.SetMinimum(0.0)
	perfData.SetMaximum(100.0)

	GlobalReturnCode = exitOk

	if thresholds.IsSet(aI.Warning) {
		perfData.SetWarning(*aI.Warning)

		if thresholds.CheckUpper(*aI.Warning, downstreamUsage) {
			GlobalReturnCode = exitWarning
		}
	}

	if thresholds.IsSet(aI.Critical) {
		perfData.SetCritical(*aI.Critical)

		if thresholds.CheckUpper(*aI.Critical, downstreamUsage) {
			GlobalReturnCode = exitCritical
		}
	}

	output := " - " + fmt.Sprintf("%.2f", downstreamUsage) + "% Downstream utilization (" + fmt.Sprintf("%.2f", downstreamCurrent) + " Mbit/s of " + fmt.Sprintf("%.2f", downstreamMax) + " Mbits) " + perfData.GetPerformanceDataAsString()

	switch GlobalReturnCode {
	case exitOk:
		fmt.Print("OK" + output + "\n")
	case exitWarning:
		fmt.Print("WARNING" + output + "\n")
	case exitCritical:
		fmt.Print("CRITICAL" + output + "\n")
	default:
		GlobalReturnCode = exitUnknown
		fmt.Print("UNKNWON - Not able to calculate downstream utilization\n")
	}
}
