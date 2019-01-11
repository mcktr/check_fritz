package main

import (
	"flag"
	"fmt"
	"os"
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
