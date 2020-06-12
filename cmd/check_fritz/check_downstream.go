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
	// First query to collect total number of sync groups
	initialResponses := make(chan []byte)
	initialErrors := make(chan error)

	initialSoapRequest := fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/wancommonifconfig1", "WANCommonInterfaceConfig", "X_AVM-DE_GetOnlineMonitor")
	initialSoapRequest.AddSoapDataVariable(fritz.CreateNewSoapVariable("NewSyncGroupIndex", "0"))
	go fritz.DoSoapRequest(&initialSoapRequest, initialResponses, initialErrors, aI.Debug)

	initialResponse, err := fritz.ProcessSoapResponse(initialResponses, initialErrors, 1, *aI.Timeout)
	if err != nil {
		fmt.Printf("UNKNOWN - %s\n", err)
		return
	}

	initialSoapResponse := fritz.WANCommonInterfaceOnlineMonitorResponse{}
	err = fritz.UnmarshalSoapResponse(&initialSoapResponse, initialResponse)
	if err != nil {
		fmt.Printf("UNKNOWN - %s\n", err)
		return
	}

	soapResponses := []fritz.WANCommonInterfaceOnlineMonitorResponse{}
	soapResponses = append(soapResponses, initialSoapResponse)

	totalNumberSyncGroups, err := strconv.Atoi(initialSoapResponse.NewTotalNumberSyncGroups)
	if err != nil {
		fmt.Printf("UNKNOWN - %s\n", err)
		return
	}

	// If there are more sync groups query all other sync groups, starting with the next one
	if totalNumberSyncGroups > 1 {
		for i := 1; i < totalNumberSyncGroups; i++ {
			responses := make(chan []byte)
			errors := make(chan error)

			soapRequest := initialSoapRequest
			soapRequest.AddSoapDataVariable(fritz.CreateNewSoapVariable("NewSyncGroupIndex", strconv.Itoa(i)))

			go fritz.DoSoapRequest(&soapRequest, responses, errors, aI.Debug)

			response, err := fritz.ProcessSoapResponse(responses, errors, 1, *aI.Timeout)
			if err != nil {
				fmt.Printf("UNKNOWN - %s\n", err)
				return
			}

			soapResponse := fritz.WANCommonInterfaceOnlineMonitorResponse{}
			err = fritz.UnmarshalSoapResponse(&soapResponse, response)
			if err != nil {
				fmt.Printf("UNKNOWN - %s\n", err)
				return
			}

			soapResponses = append(soapResponses, soapResponse)
		}
	}

	downstreamCombined := 0.0
	output := ""
	performanceData := []perfdata.PerformanceData{}
	GlobalReturnCode = exitOk

	for _, r := range soapResponses {
		downstream, err := strconv.ParseFloat(r.NewMaxDS, 64)
		if err != nil {
			fmt.Printf("UNKNOWN - %s\n", err)
			return
		}

		downstreamCombined += downstream
		downstream = downstream * 8 / 1000000

		output += fmt.Sprintf(", Downstream SyncGroup '%s': %.2f Mbit/s", r.NewSyncGroupName, downstream)

		pd := perfdata.CreatePerformanceData("downstream_max_"+fmt.Sprintf("%s", r.NewSyncGroupName), downstream, "")

		if thresholds.IsSet(aI.Warning) {
			pd.SetWarning(*aI.Warning)

			if thresholds.CheckLower(*aI.Warning, downstream) {
				GlobalReturnCode = exitWarning
			}
		}

		if thresholds.IsSet(aI.Critical) {
			pd.SetCritical(*aI.Critical)

			if thresholds.CheckLower(*aI.Critical, downstream) {
				GlobalReturnCode = exitCritical
			}
		}

		performanceData = append(performanceData, *pd)
	}

	downstreamCombined = downstreamCombined * 8 / 1000000
	pdDownstreamCombined := perfdata.CreatePerformanceData("downstream_max", downstreamCombined, "")

	if thresholds.IsSet(aI.Warning) {
		pdDownstreamCombined.SetWarning(*aI.Warning)

		if thresholds.CheckLower(*aI.Warning, downstreamCombined) {
			GlobalReturnCode = exitWarning
		}
	}

	if thresholds.IsSet(aI.Critical) {
		pdDownstreamCombined.SetCritical(*aI.Critical)

		if thresholds.CheckLower(*aI.Critical, downstreamCombined) {
			GlobalReturnCode = exitCritical
		}
	}

	performanceData = append(performanceData, *pdDownstreamCombined)

	output = " - Max Downstream: " + fmt.Sprintf("%.2f", downstreamCombined) + " Mbit/s" + output + perfdata.FormatAsString(performanceData)

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
