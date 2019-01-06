package main

import (
	"flag"
	"fmt"
	"os"

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

func CheckStatus(hostname string, port string, username string, password string) {
	url := "https://" + hostname + ":" + port + "/upnp/control/wanpppconn1"

	soapReq := fritz.NewSoapRequest(url, username, password, "WANPPPConnection", "GetInfo")

	err := fritz.DoSoapRequest(&soapReq)

	if HandleError(err) {
		return
	}

	status, err := fritz.HandleSoapRequest(&soapReq)

	if HandleError(err) {
		return
	}

	if status == "Connected" {
		fmt.Print("OK - Connection Status: " + status + "\n")

		GlobalReturnCode = exitOk
	} else {
		fmt.Print("CRITICAL - Connection Status: " + status + "\n")

		GlobalReturnCode = exitCritical
	}
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

	flag.Parse()

	if !CheckRequiredFlags(*hostname, *port, *username, *password) {
		os.Exit(exitUnknown)
	}

	CheckStatus(*hostname, *port, *username, *password)

	os.Exit(GlobalReturnCode)
}
