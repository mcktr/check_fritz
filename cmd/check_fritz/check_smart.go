package main

import (
	"fmt"
	"strconv"

	"github.com/mcktr/check_fritz/modules/fritz"
	"github.com/mcktr/check_fritz/modules/perfdata"
	"github.com/mcktr/check_fritz/modules/thresholds"
)

// CheckSmartStatus checks the connection status of a smart device
func CheckSmartStatus(aI ArgumentInformation) {
	resps := make(chan []byte)
	errs := make(chan error)

	soapReq := fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/x_homeauto", "X_AVM-DE_Homeauto", "GetGenericDeviceInfos")
	soapReq.AddSoapDataVariable(fritz.CreateNewSoapVariable("NewIndex", strconv.Itoa(*aI.Index)))
	go fritz.DoSoapRequest(&soapReq, resps, errs)

	res, err := fritz.ProcessSoapResponse(resps, errs, 1)

	if err != nil {
		panic(err)
	}

	soapResp := fritz.SmartDeviceInfoResponse{}
	err = fritz.UnmarshalSoapResponse(&soapResp, res)

	if err != nil {
		panic(err)
	}

	output := "- " + soapResp.NewProductName + " " + soapResp.NewFirmwareVersion + " - " + soapResp.NewDeviceName + " " + soapResp.NewPresent
	GlobalReturnCode = exitOk

	if soapResp.NewPresent != "CONNECTED" {
		GlobalReturnCode = exitCritical
	}

	switch GlobalReturnCode {
	case exitOk:
		fmt.Print("OK " + output + "\n")
	case exitWarning:
		fmt.Print("WARNING " + output + "\n")
	case exitCritical:
		fmt.Print("CRITICAL " + output + "\n")
	default:
		GlobalReturnCode = exitUnknown
		fmt.Print("UNKNWON - Not able to determine smart device status\n")
	}
}

// CheckSmartHeaterTemperatur checks the temperature of a smart home thermometer device
func CheckSmartHeaterTemperatur(aI ArgumentInformation) {
	resps := make(chan []byte)
	errs := make(chan error)

	soapReq := fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/x_homeauto", "X_AVM-DE_Homeauto", "GetGenericDeviceInfos")
	soapReq.AddSoapDataVariable(fritz.CreateNewSoapVariable("NewIndex", strconv.Itoa(*aI.Index)))
	go fritz.DoSoapRequest(&soapReq, resps, errs)

	res, err := fritz.ProcessSoapResponse(resps, errs, 1)

	if err != nil {
		panic(err)
	}

	soapResp := fritz.SmartDeviceInfoResponse{}
	err = fritz.UnmarshalSoapResponse(&soapResp, res)

	if err != nil {
		panic(err)
	}

	if soapResp.NewTemperatureIsEnabled != "ENABLED" {
		fmt.Print("UNKNOWN - Temperature is not enabled on this smart device\n")
		GlobalReturnCode = exitUnknown
		return
	}

	currentTemp, err := strconv.ParseFloat(soapResp.NewTemperatureCelsius, 64)

	if err != nil {
		panic(err)
	}

	currentTemp = currentTemp / 10.0
	perfData := perfdata.CreatePerformanceData("temperature", currentTemp, "")

	GlobalReturnCode = exitOk

	if thresholds.IsSet(aI.Warning) {
		perfData.SetWarning(*aI.Warning)

		if thresholds.CheckLower(*aI.Warning, currentTemp) {
			GlobalReturnCode = exitWarning
		}
	}

	if thresholds.IsSet(aI.Critical) {
		perfData.SetCritical(*aI.Critical)

		if thresholds.CheckLower(*aI.Critical, currentTemp) {
			GlobalReturnCode = exitCritical
		}
	}

	output := "- " + soapResp.NewProductName + " " + soapResp.NewFirmwareVersion + " - " + soapResp.NewDeviceName + " " + fmt.Sprintf("%.2f", currentTemp) + " °C " + perfData.GetPerformanceDataAsString()

	switch GlobalReturnCode {
	case exitOk:
		fmt.Print("OK " + output + "\n")
	case exitWarning:
		fmt.Print("WARNING " + output + "\n")
	case exitCritical:
		fmt.Print("CRITICAL " + output + "\n")
	default:
		GlobalReturnCode = exitUnknown
		fmt.Print("UNKNWON - Not able to calculate heater temperature\n")
	}
}

// CheckSmartSocketPower checks the current watt usage on the smart socket
func CheckSmartSocketPower(aI ArgumentInformation) {
	resps := make(chan []byte)
	errs := make(chan error)

	soapReq := fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/x_homeauto", "X_AVM-DE_Homeauto", "GetGenericDeviceInfos")
	soapReq.AddSoapDataVariable(fritz.CreateNewSoapVariable("NewIndex", strconv.Itoa(*aI.Index)))
	go fritz.DoSoapRequest(&soapReq, resps, errs)

	res, err := fritz.ProcessSoapResponse(resps, errs, 1)

	if err != nil {
		panic(err)
	}

	soapResp := fritz.SmartDeviceInfoResponse{}
	err = fritz.UnmarshalSoapResponse(&soapResp, res)

	if err != nil {
		panic(err)
	}

	currentPower, err := strconv.ParseFloat(soapResp.NewMultimeterPower, 64)

	if err != nil {
		panic(err)
	}

	currentPower = currentPower / 100.0
	perfData := perfdata.CreatePerformanceData("power", currentPower, "")

	GlobalReturnCode = exitOk

	if thresholds.IsSet(aI.Warning) {
		perfData.SetWarning(*aI.Warning)

		if thresholds.CheckUpper(*aI.Warning, currentPower) {
			GlobalReturnCode = exitWarning
		}
	}

	if thresholds.IsSet(aI.Critical) {
		perfData.SetCritical(*aI.Critical)

		if thresholds.CheckUpper(*aI.Critical, currentPower) {
			GlobalReturnCode = exitCritical
		}
	}

	output := "- " + soapResp.NewProductName + " " + soapResp.NewFirmwareVersion + " - " + soapResp.NewDeviceName + " " + fmt.Sprintf("%.2f", currentPower) + " W " + perfData.GetPerformanceDataAsString()

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

// CheckSmartSocketEnergy checks total power consumption of the last year on the smart socket
func CheckSmartSocketEnergy(aI ArgumentInformation) {
	resps := make(chan []byte)
	errs := make(chan error)

	soapReq := fritz.CreateNewSoapData(*aI.Username, *aI.Password, *aI.Hostname, *aI.Port, "/upnp/control/x_homeauto", "X_AVM-DE_Homeauto", "GetGenericDeviceInfos")
	soapReq.AddSoapDataVariable(fritz.CreateNewSoapVariable("NewIndex", strconv.Itoa(*aI.Index)))
	go fritz.DoSoapRequest(&soapReq, resps, errs)

	res, err := fritz.ProcessSoapResponse(resps, errs, 1)

	if err != nil {
		panic(err)
	}

	soapResp := fritz.SmartDeviceInfoResponse{}
	err = fritz.UnmarshalSoapResponse(&soapResp, res)

	if err != nil {
		panic(err)
	}

	currentEnergy, err := strconv.ParseFloat(soapResp.NewMultimeterEnergy, 64)

	if err != nil {
		panic(err)
	}

	currentEnergy = currentEnergy / 1000.0
	perfData := perfdata.CreatePerformanceData("energy", currentEnergy, "")

	GlobalReturnCode = exitOk

	if thresholds.IsSet(aI.Warning) {
		perfData.SetWarning(*aI.Warning)

		if thresholds.CheckUpper(*aI.Warning, currentEnergy) {
			GlobalReturnCode = exitWarning
		}
	}

	if thresholds.IsSet(aI.Critical) {
		perfData.SetCritical(*aI.Critical)

		if thresholds.CheckUpper(*aI.Critical, currentEnergy) {
			GlobalReturnCode = exitCritical
		}
	}

	output := "- " + soapResp.NewProductName + " " + soapResp.NewFirmwareVersion + " - " + soapResp.NewDeviceName + " " + fmt.Sprintf("%.2f", currentEnergy) + " kWh " + perfData.GetPerformanceDataAsString()

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
