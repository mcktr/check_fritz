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

	upstreamCombined := 0.0
	output := ""
	performanceData := []perfdata.PerformanceData{}
	GlobalReturnCode = exitOk

	for _, r := range soapResponses {
		if totalNumberSyncGroups > 1 {
			if isInIgnoreList(aI.SyncGroupIgnoreList, r.NewSyncGroupName) {
				if aI.Debug {
					fmt.Printf("Ignoring sync group '%s' since it is in the ignore list\n", r.NewSyncGroupName)
				}

				continue
			}
		} else {
			if aI.Debug {
				fmt.Printf("The total number of sync groups is > 1, there is noting to ignore\n")
			}
		}

		upstream, err := strconv.ParseFloat(r.NewMaxUS, 64)
		if err != nil {
			fmt.Printf("UNKNOWN - %s\n", err)
			return
		}

		upstreamCombined += upstream
		upstream = upstream * 8 / 1000000

		output += fmt.Sprintf(", Upstream SyncGroup '%s': %.2f Mbit/s", r.NewSyncGroupName, upstream)

		pd := perfdata.CreatePerformanceData("upstream_max_"+fmt.Sprintf("%s", r.NewSyncGroupName), upstream, "")

		if thresholds.IsSet(aI.Warning) {
			pd.SetWarning(*aI.Warning)

			if thresholds.CheckLower(*aI.Warning, upstream) {
				GlobalReturnCode = exitWarning
			}
		}

		if thresholds.IsSet(aI.Critical) {
			pd.SetCritical(*aI.Critical)

			if thresholds.CheckLower(*aI.Critical, upstream) {
				GlobalReturnCode = exitCritical
			}
		}

		performanceData = append(performanceData, *pd)
	}

	upstreamCombined = upstreamCombined * 8 / 1000000
	pdUpstreamCombined := perfdata.CreatePerformanceData("upstream_max", upstreamCombined, "")

	if thresholds.IsSet(aI.Warning) {
		pdUpstreamCombined.SetWarning(*aI.Warning)

		if thresholds.CheckLower(*aI.Warning, upstreamCombined) {
			GlobalReturnCode = exitWarning
		}
	}

	if thresholds.IsSet(aI.Critical) {
		pdUpstreamCombined.SetCritical(*aI.Critical)

		if thresholds.CheckLower(*aI.Critical, upstreamCombined) {
			GlobalReturnCode = exitCritical
		}
	}

	performanceData = append(performanceData, *pdUpstreamCombined)

	output = " - Max Upstream: " + fmt.Sprintf("%.2f", upstreamCombined) + " Mbit/s" + output + perfdata.FormatAsString(performanceData)

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
