package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mcktr/check_fritz/pkg/fritz"
)

func CheckDownstreamMax(hostname, port string, username string, password string) {
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

	downstream, err := strconv.Atoi(resp.Newmax_ds)

	if HandleError(err) {
		return
	}

	downstream = downstream * 8 / 1000000

	fmt.Print("OK - Max Downstream: " + strconv.Itoa(downstream) + " Mbit/s \n")

	GlobalReturnCode = exitOk
}

func CheckDownstreamCurrent(hostname string, port string, username string, password string) {
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

	downstreamWithHistory := strings.Split(resp.Newds_current_bps, ",")

	downstream, err := strconv.ParseFloat(downstreamWithHistory[0], 32)

	if HandleError(err) {
		return
	}

	downstream = downstream * 8 / 1000000

	fmt.Print("OK - Current Downstream: " + fmt.Sprintf("%f", downstream) + " Mbit/s \n")

	GlobalReturnCode = exitOk
}
