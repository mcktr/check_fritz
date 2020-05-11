package main

import (
	"fmt"
	"strconv"

	"github.com/mcktr/check_fritz/modules/fritz"
	"github.com/mcktr/check_fritz/modules/perfdata"
)

// CheckDeviceUptime checks the uptime of the device
func CheckDeviceUptime(aI ArgumentInformation) {
	resps := make(chan []byte)
	errs := make(chan error)

	soapReq := fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/deviceinfo", "DeviceInfo", "GetInfo")
	go fritz.DoSoapRequest(&soapReq, resps, errs, aI.Debug)

	res, err := fritz.ProcessSoapResponse(resps, errs, 1, *aI.Timeout)

	if err != nil {
		fmt.Printf("UNKNOWN - %s\n", err)
		return
	}

	soapResp := fritz.DeviceInfoResponse{}
	err = fritz.UnmarshalSoapResponse(&soapResp, res)

	if err != nil {
		panic(err)
	}

	uptime, err := strconv.Atoi(soapResp.NewUpTime)

	if err != nil {
		panic(err)
	}

	days := uptime / 86400
	hours := (uptime / 3600) - (days * 24)
	minutes := (uptime / 60) - (days * 1440) - (hours * 60)
	seconds := uptime % 60
	output := fmt.Sprintf("%dd %dh %dm %ds", days, hours, minutes, seconds)

	perfData := perfdata.CreatePerformanceData("uptime", float64(uptime), "s")

	fmt.Print("OK - Device Uptime: " + fmt.Sprintf("%d", uptime) + " seconds (" + output + ") " + perfData.GetPerformanceDataAsString() + "\n")

	GlobalReturnCode = exitOk
}

// CheckDeviceUpdate checks if a new firmware is available
func CheckDeviceUpdate(aI ArgumentInformation) {
	resps := make(chan []byte)
	errs := make(chan error)

	soapReq := fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/userif", "UserInterface", "GetInfo")
	go fritz.DoSoapRequest(&soapReq, resps, errs, aI.Debug)

	res, err := fritz.ProcessSoapResponse(resps, errs, 1, *aI.Timeout)

	if err != nil {
		fmt.Printf("UNKNOWN - %s\n", err)
		return
	}

	soapResp := fritz.UserInterfaceInfoResponse{}
	err = fritz.UnmarshalSoapResponse(&soapResp, res)

	if err != nil {
		panic(err)
	}

	state, err := strconv.Atoi(soapResp.NewUpgradeAvailable)

	if err != nil {
		panic(err)
	}

	GlobalReturnCode = exitOk

	if state == 0 {
		fmt.Print("OK - No update available\n")
	} else {
		GlobalReturnCode = exitCritical

		fmt.Print("CRITICAL - Update available\n")
	}
}
