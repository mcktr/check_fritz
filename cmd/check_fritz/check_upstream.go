package main

import (
	"fmt"
	"strconv"

	"github.com/mcktr/check_fritz/pkg/fritz"
)

func CheckUpstreamMax(hostname, port string, username string, password string) {
	soapReq := fritz.NewSoapRequest(username, password, hostname, port, "/upnp/control/wancommonifconfig1", "WANCommonInterfaceConfig", "X_AVM-DE_GetOnlineMonitor")
	fritz.AddSoapRequestVariable(&soapReq, fritz.NewSoapRequestVariable("NewSyncGroupIndex", "0"))

	err := fritz.DoSoapRequest(&soapReq)

	if HandleError(err) {
		return
	}

	var resp = fritz.GetWANCommonInterfaceOnlineMonitorResponse{}

	err = fritz.HandleSoapRequest(&soapReq, &resp)

	if HandleError(err) {
		return
	}

	upstream, err := strconv.Atoi(resp.Newmax_us)

	if HandleError(err) {
		return
	}

	upstream = upstream * 8 / 1000000

	fmt.Print("OK - Max Upstream: " + strconv.Itoa(upstream) + " Mbit/s \n")

	GlobalReturnCode = exitOk
}
