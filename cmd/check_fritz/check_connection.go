package main

import (
	"fmt"

	"github.com/mcktr/check_fritz/pkg/fritz"
)

func CheckConnectionStatus(hostname string, port string, username string, password string) {
	soapReq := fritz.NewSoapRequest(username, password, hostname, port, "/upnp/control/wanpppconn1", "WANPPPConnection", "GetInfo")

	err := fritz.DoSoapRequest(&soapReq)

	if HandleError(err) {
		return
	}

	var resp = fritz.GetWANPPPConnectionInfoResponse{}

	err = fritz.HandleSoapRequest(&soapReq, &resp)

	if HandleError(err) {
		return
	}

	if resp.NewConnectionStatus == "Connected" {
		fmt.Print("OK - Connection Status: " + resp.NewConnectionStatus + "\n")

		GlobalReturnCode = exitOk
	} else {
		fmt.Print("CRITICAL - Connection Status: " + resp.NewConnectionStatus + "\n")

		GlobalReturnCode = exitCritical
	}
}

func CheckConnectionUptime(hostname, port string, username string, password string) {
	soapReq := fritz.NewSoapRequest(username, password, hostname, port, "/upnp/control/wanpppconn1", "WANPPPConnection", "GetInfo")

	err := fritz.DoSoapRequest(&soapReq)

	if HandleError(err) {
		return
	}

	var resp = fritz.GetWANPPPConnectionInfoResponse{}

	err = fritz.HandleSoapRequest(&soapReq, &resp)

	if HandleError(err) {
		return
	}

	fmt.Print("OK - Connection Uptime: " + resp.NewUptime + "\n")

	GlobalReturnCode = exitOk
}
