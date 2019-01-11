package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/mcktr/check_fritz/pkg/fritz"
)

// program version
var version string = "0.1"

// internal exit codes
var exitOk = 0
var exitWarning = 1
var exitCritical = 2
var exitUnknown = 3

var GlobalReturnCode = exitUnknown

func GetVersion() string {
	return "check_fritz version " + version
}

func CheckRequiredFlags(hostname string, port string, username string, password string) bool {
	if hostname == "" {
		fmt.Println("No hostname")
		return false
	}

	if port == "" {
		fmt.Println("No port")
		return false
	}

	if username == "" {
		fmt.Println("No username")
		return false
	}

	if password == "" {
		fmt.Println("No password")
		return false
	}

	return true
}

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

func CheckDeviceUptime(hostname, port string, username string, password string) {
	soapReq := fritz.NewSoapRequest(username, password, hostname, port, "/upnp/control/deviceinfo", "DeviceInfo", "GetInfo")

	err := fritz.DoSoapRequest(&soapReq)

	if HandleError(err) {
		return
	}

	var resp = fritz.GetDeviceInfoResponse{}

	err = fritz.HandleSoapRequest(&soapReq, &resp)

	if HandleError(err) {
		return
	}

	fmt.Print("OK - Device Uptime: " + resp.NewUpTime + "\n")

	GlobalReturnCode = exitOk
}

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

func HandleError(err error) bool {
	if err != nil {
		fmt.Println(err)
		GlobalReturnCode = exitUnknown
		return true
	}

	return false
}

func main() {
	var hostname = flag.String("hostname", "fritz.box", "Specify the hostname")
	var port = flag.String("port", "49443", "SSL port")
	var username = flag.String("username", "dslf-config", "Specify the username")
	var password = flag.String("password", "", "Specify the password")
	var method = flag.String("method", "connection_status", "Specify the used method. (Default: status)")

	flag.Parse()

	if !CheckRequiredFlags(*hostname, *port, *username, *password) {
		os.Exit(exitUnknown)
	}

	switch *method {
	case "connection_status":
		CheckConnectionStatus(*hostname, *port, *username, *password)
	case "connection_uptime":
		CheckConnectionUptime(*hostname, *port, *username, *password)
	case "device_uptime":
		CheckDeviceUptime(*hostname, *port, *username, *password)
	case "downstream_max":
		CheckDownstreamMax(*hostname, *port, *username, *password)
	case "upstream_max":
		CheckUpstreamMax(*hostname, *port, *username, *password)
	default:
		fmt.Println("Unknown method.")
		GlobalReturnCode = exitUnknown
	}

	os.Exit(GlobalReturnCode)
}
