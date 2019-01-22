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

type ArgumentInformation struct {
	Hostname string
	Port     string
	Username string
	Password string
	Method   string
	Warning  float64
	Critical float64
}

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

func internalCheckLower(threshold float64, val float64) bool {
	return threshold > val
}

func internalCheckUpper(threshold float64, val float64) bool {
	return threshold < val
}

func main() {
	var hostname = flag.String("hostname", "fritz.box", "Specify the hostname")
	var port = flag.String("port", "49443", "SSL port")
	var username = flag.String("username", "dslf-config", "Specify the username")
	var password = flag.String("password", "", "Specify the password")
	var method = flag.String("method", "connection_status", "Specify the used method. (Default: status)")
	var warning = flag.Float64("warning", 0, "Specify the warning threshold")
	var critical = flag.Float64("critical", 0, "Specify the critical threshold")

	flag.Parse()

	aI := ArgumentInformation{*hostname, *port, *username, *password, *method, *warning, *critical}

	if !CheckRequiredFlags(aI.Hostname, aI.Port, aI.Username, aI.Password) {
		os.Exit(exitUnknown)
	}

	switch aI.Method {
	case "connection_status":
		CheckConnectionStatus(aI)
	case "connection_uptime":
		CheckConnectionUptime(aI)
	case "device_uptime":
		CheckDeviceUptime(aI)
	case "downstream_max":
		CheckDownstreamMax(aI)
	case "upstream_max":
		CheckUpstreamMax(aI)
	case "downstream_current":
		CheckDownstreamCurrent(aI)
	case "upstream_current":
		CheckUpstreamCurrent(aI)
	case "interface_update":
		CheckInterfaceUpdate(aI)
	default:
		fmt.Println("Unknown method.")
		GlobalReturnCode = exitUnknown
	}

	os.Exit(GlobalReturnCode)
}
