package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mcktr/check_fritz/modules/fritz"
	"github.com/mcktr/check_fritz/modules/perfdata"
)

// CheckConnectionStatus checks the internet connection status
func CheckConnectionStatus(aI ArgumentInformation) {
	resps := make(chan []byte)
	errs := make(chan error)

	modelgroup := strings.ToLower(*aI.Modelgroup)

	var soapReq fritz.SoapData

	switch modelgroup {
	case "dsl":
		soapReq = fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/wanpppconn1", "WANPPPConnection", "GetInfo")
	case "cable":
		soapReq = fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/wanipconnection1", "WanIPConnection", "GetInfo")
	default:
		fmt.Printf("UNKNOWN - Fritz!Box modelgroup '%s' is unknown. Supported modelgroups are: DSL, CABLE\n", modelgroup)
		GlobalReturnCode = exitUnknown
		return
	}
	go fritz.DoSoapRequest(&soapReq, resps, errs, aI.Debug)

	res, err := fritz.ProcessSoapResponse(resps, errs, 1, *aI.Timeout)

	if err != nil {
		fmt.Printf("UNKNOWN - %s\n", err)
		return
	}

	soapResp := fritz.WANConnectionInfoResponse{}
	err = fritz.UnmarshalSoapResponse(&soapResp, res)

	if err != nil {
		panic(err)
	}

	output := ""

	if soapResp.NewConnectionStatus == "Connected" {
		output = fmt.Sprintf("OK - Connection Status: %s", soapResp.NewConnectionStatus)

		if soapResp.NewExternalIPAddress != "" {
			output += fmt.Sprintf("; External IP: %s", soapResp.NewExternalIPAddress)
		}

		GlobalReturnCode = exitOk
	} else if soapResp.NewConnectionStatus == "" {
		output = fmt.Sprint("UNKNOWN - Connection Status is empty")

		GlobalReturnCode = exitUnknown
	} else if soapResp.NewConnectionStatus == "Connecting" || soapResp.NewConnectionStatus == "Authenticating" {
		output = fmt.Sprintf("WARNING - Connection Status: %s", soapResp.NewConnectionStatus)

		GlobalReturnCode = exitWarning
	} else {
		output = fmt.Sprintf("CRITICAL - Connection Status: %s", soapResp.NewConnectionStatus)

		GlobalReturnCode = exitCritical
	}

	perfData := perfdata.CreatePerformanceData("status", float64(GlobalReturnCode), "")

	fmt.Printf("%s %s\n", output, perfData.GetPerformanceDataAsString())
}

// CheckConnectionUptime checks the uptime of the internet connection
func CheckConnectionUptime(aI ArgumentInformation) {
	resps := make(chan []byte)
	errs := make(chan error)

	modelgroup := strings.ToLower(*aI.Modelgroup)

	var soapReq fritz.SoapData

	switch modelgroup {
	case "dsl":
		soapReq = fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/wanpppconn1", "WANPPPConnection", "GetInfo")
	case "cable":
		soapReq = fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/wanipconnection1", "WanIPConnection", "GetInfo")
	default:
		fmt.Printf("UNKNOWN - Fritz!Box modelgroup '%s' is unknown. Supported modelgroups are: DSL, CABLE\n", modelgroup)
		GlobalReturnCode = exitUnknown
		return
	}
	go fritz.DoSoapRequest(&soapReq, resps, errs, aI.Debug)

	res, err := fritz.ProcessSoapResponse(resps, errs, 1, *aI.Timeout)

	if err != nil {
		fmt.Printf("UNKNOWN - %s\n", err)
		return
	}

	soapResp := fritz.WANConnectionInfoResponse{}
	err = fritz.UnmarshalSoapResponse(&soapResp, res)

	if err != nil {
		panic(err)
	}

	if soapResp.NewUptime != "" {
		uptime, err := strconv.Atoi(soapResp.NewUptime)

		if err != nil {
			panic(err)
		}

		days := uptime / 86400
		hours := (uptime / 3600) - (days * 24)
		minutes := (uptime / 60) - (days * 1440) - (hours * 60)
		seconds := uptime % 60
		output := fmt.Sprintf("%dd %dh %dm %ds", days, hours, minutes, seconds)
		perfData := perfdata.CreatePerformanceData("uptime", float64(uptime), "s")

		fmt.Print("OK - Connection Uptime: " + fmt.Sprintf("%d", uptime) + " seconds (" + output + ") " + perfData.GetPerformanceDataAsString() + "\n")

		GlobalReturnCode = exitOk
	} else {
		fmt.Print("UNKNOWN - Connection Uptime is empty\n")

		GlobalReturnCode = exitUnknown
	}
}
