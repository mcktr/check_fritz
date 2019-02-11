package main

import (
	"fmt"
	"strconv"

	"github.com/mcktr/check_fritz/pkg/fritz"
)

// CheckInterfaceUpdate checks if a new firmware is available
func CheckInterfaceUpdate(aI ArgumentInformation) {
	soapReq := fritz.NewSoapRequest(aI.Username, aI.Password, aI.Hostname, aI.Port, "/upnp/control/userif", "UserInterface", "GetInfo")

	err := fritz.DoSoapRequest(&soapReq)

	if HandleError(err) {
		return
	}

	var resp = fritz.GetInterfaceInfoResponse{}

	err = fritz.HandleSoapRequest(&soapReq, &resp)

	if HandleError(err) {
		return
	}

	state, err := strconv.Atoi(resp.NewUpgradeAvailable)

	if HandleError(err) {
		return
	}

	GlobalReturnCode = exitOk

	if state == 0 {
		GlobalReturnCode = exitOk

		fmt.Print("OK - No update avaiable\n")
	} else {
		GlobalReturnCode = exitCritical

		fmt.Print("CRITICAL - Update available\n")
	}
}
