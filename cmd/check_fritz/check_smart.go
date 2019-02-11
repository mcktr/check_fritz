package main

import (
	"fmt"
	"strconv"

	"github.com/mcktr/check_fritz/pkg/fritz"
	"github.com/mcktr/check_fritz/pkg/thresholds"
)

func CheckSmartThermometer(aI ArgumentInformation) {
	soapReq := fritz.NewSoapRequest(aI.Username, aI.Password, aI.Hostname, aI.Port, "/upnp/control/x_homeauto", "X_AVM-DE_Homeauto", "GetGenericDeviceInfos")
	fritz.AddSoapRequestVariable(&soapReq, fritz.NewSoapRequestVariable("NewIndex", strconv.Itoa(aI.Index)))

	err := fritz.DoSoapRequest(&soapReq)

	if HandleError(err) {
		return
	}

	var resp = fritz.GetSmartDeviceInfoResponse{}

	err = fritz.HandleSoapRequest(&soapReq, &resp)

	if HandleError(err) {
		return
	}

	if resp.NewTemperatureIsEnabled != "ENABLED" {
		fmt.Print("UNKNOWN - Temperature is not enabled on this smart device\n")
		GlobalReturnCode = exitUnknown
		return
	}

	currentTemp, err := strconv.ParseFloat(resp.NewTemperatureCelsius, 64)

	if HandleError(err) {
		return
	}

	currentTemp = currentTemp / 10.0

	GlobalReturnCode = exitOk

	if thresholds.CheckLower(aI.Warning, currentTemp) {
		GlobalReturnCode = exitWarning
	}

	if thresholds.CheckLower(aI.Critical, currentTemp) {
		GlobalReturnCode = exitCritical
	}
	output := "- " + resp.NewProductName + " " + resp.NewFirmwareVersion + " - " + resp.NewDeviceName + " " + resp.NewPresent + " " + fmt.Sprintf("%.2f", currentTemp) + " °C"

	switch GlobalReturnCode {
	case exitOk:
		fmt.Print("OK " + output + "\n")
	case exitWarning:
		fmt.Print("WARNING " + output + "\n")
	case exitCritical:
		fmt.Print("CRITICAL " + output + "\n")
	default:
		GlobalReturnCode = exitUnknown
		fmt.Print("UNKNWON - Not able to calculate thermostat temperature\n")
	}
}

func CheckSmartSocketPower(aI ArgumentInformation) {
	soapReq := fritz.NewSoapRequest(aI.Username, aI.Password, aI.Hostname, aI.Port, "/upnp/control/x_homeauto", "X_AVM-DE_Homeauto", "GetGenericDeviceInfos")
	fritz.AddSoapRequestVariable(&soapReq, fritz.NewSoapRequestVariable("NewIndex", strconv.Itoa(aI.Index)))

	err := fritz.DoSoapRequest(&soapReq)

	if HandleError(err) {
		return
	}

	var resp = fritz.GetSmartDeviceInfoResponse{}

	err = fritz.HandleSoapRequest(&soapReq, &resp)

	if HandleError(err) {
		return
	}

	currentPower, err := strconv.ParseFloat(resp.NewMultimeterPower, 64)

	if HandleError(err) {
		return
	}

	currentPower = currentPower / 100.0

	GlobalReturnCode = exitOk

	if thresholds.CheckUpper(aI.Warning, currentPower) {
		GlobalReturnCode = exitWarning
	}

	if thresholds.CheckUpper(aI.Critical, currentPower) {
		GlobalReturnCode = exitCritical
	}

	output := "- " + resp.NewProductName + " " + resp.NewFirmwareVersion + " - " + resp.NewDeviceName + " " + resp.NewPresent + " " + fmt.Sprintf("%.2f", currentPower) + " W"

	switch GlobalReturnCode {
	case exitOk:
		fmt.Print("OK " + output + "\n")
	case exitWarning:
		fmt.Print("WARNING " + output + "\n")
	case exitCritical:
		fmt.Print("CRITICAL " + output + "\n")
	default:
		GlobalReturnCode = exitUnknown
		fmt.Print("UNKNWON - Not able to fetch socket power\n")
	}
}

func CheckSmartSocketEnergy(aI ArgumentInformation) {
	soapReq := fritz.NewSoapRequest(aI.Username, aI.Password, aI.Hostname, aI.Port, "/upnp/control/x_homeauto", "X_AVM-DE_Homeauto", "GetGenericDeviceInfos")
	fritz.AddSoapRequestVariable(&soapReq, fritz.NewSoapRequestVariable("NewIndex", strconv.Itoa(aI.Index)))

	err := fritz.DoSoapRequest(&soapReq)

	if HandleError(err) {
		return
	}

	var resp = fritz.GetSmartDeviceInfoResponse{}

	err = fritz.HandleSoapRequest(&soapReq, &resp)

	if HandleError(err) {
		return
	}

	currentEnergy, err := strconv.ParseFloat(resp.NewMultimeterEnergy, 64)

	if HandleError(err) {
		return
	}

	currentEnergy = currentEnergy / 1000.0

	GlobalReturnCode = exitOk

	if thresholds.CheckUpper(aI.Warning, currentEnergy) {
		GlobalReturnCode = exitWarning
	}

	if thresholds.CheckUpper(aI.Critical, currentEnergy) {
		GlobalReturnCode = exitCritical
	}

	output := "- " + resp.NewProductName + " " + resp.NewFirmwareVersion + " - " + resp.NewDeviceName + " " + resp.NewPresent + " " + fmt.Sprintf("%.2f", currentEnergy) + " kWh"

	switch GlobalReturnCode {
	case exitOk:
		fmt.Print("OK " + output + "\n")
	case exitWarning:
		fmt.Print("WARNING " + output + "\n")
	case exitCritical:
		fmt.Print("CRITICAL " + output + "\n")
	default:
		GlobalReturnCode = exitUnknown
		fmt.Print("UNKNWON - Not able to fetch socket energy\n")
	}
}
