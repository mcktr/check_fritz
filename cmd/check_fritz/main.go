package main

import (
	"fmt"
	"os"
	"strconv"

	cmdline "github.com/galdor/go-cmdline"
)

// program version
var version = "0.1"

// internal exit codes
var exitOk = 0
var exitWarning = 1
var exitCritical = 2
var exitUnknown = 3

// GlobalReturnCode holds always the last set return code
var GlobalReturnCode = exitUnknown

// ArgumentInformation is the data structure for the passed arguments
type ArgumentInformation struct {
	Hostname *string
	Port     *string
	Username *string
	Password *string
	Method   *string
	Warning  *float64
	Critical *float64
	Index    *int
}

func createRequiredArgumentInformation(hostname string, port string, username string, password string, method string) ArgumentInformation {
	var ai ArgumentInformation

	ai.Hostname = &hostname
	ai.Port = &port
	ai.Username = &username
	ai.Password = &password
	ai.Method = &method

	return ai
}

func (ai *ArgumentInformation) createWarningThreshold(warning string) {
	warn, err := strconv.ParseFloat(warning, 64)

	if HandleError(err) {
		return
	}

	ai.Warning = &warn

}

func (ai *ArgumentInformation) createCriticalThreshold(critical string) {
	crit, err := strconv.ParseFloat(critical, 64)

	if HandleError(err) {
		return
	}

	ai.Critical = &crit
}

func (ai *ArgumentInformation) createIndex(index string) {
	ind, err := strconv.Atoi(index)

	if HandleError(err) {
		return
	}

	ai.Index = &ind
}

func getVersion() string {
	return "check_fritz version " + version
}

func checkRequiredFlags(hostname string, port string, username string, password string) bool {
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

// HandleError is the global error handler for the programm
func HandleError(err error) bool {
	if err != nil {
		fmt.Println(err)
		GlobalReturnCode = exitUnknown
		return true
	}

	return false
}

func main() {
	cmdline := cmdline.New()

	cmdline.AddOption("H", "hostname", "value", "Specifies the hostname. (Default: fritz.box)")
	cmdline.AddOption("P", "port", "value", "Specifies the SSL port. (Default: 49443)")
	cmdline.AddOption("u", "username", "value", "Specifies the username. (Default: dslf-config)")
	cmdline.AddOption("p", "password", "value", "Specifies the password.")
	cmdline.AddOption("m", "method", "value", "Specifies the check method. (Default: connection_status)")
	cmdline.AddOption("w", "warning", "value", "Specifies the warning threshold.")
	cmdline.AddOption("c", "critical", "value", "Specifies the critical threshold.")
	cmdline.AddOption("i", "index", "value", "Specifies the index.")

	cmdline.SetOptionDefault("hostname", "fritz.box")
	cmdline.SetOptionDefault("port", "49443")
	cmdline.SetOptionDefault("username", "dslf-config")
	cmdline.SetOptionDefault("method", "connection_status")

	cmdline.Parse(os.Args)

	hostname := cmdline.OptionValue("hostname")
	port := cmdline.OptionValue("port")
	username := cmdline.OptionValue("username")
	password := cmdline.OptionValue("password")
	method := cmdline.OptionValue("method")

	aI := createRequiredArgumentInformation(hostname, port, username, password, method)

	if cmdline.IsOptionSet("warning") {
		aI.createWarningThreshold(cmdline.OptionValue("warning"))
	}

	if cmdline.IsOptionSet("critical") {
		aI.createCriticalThreshold(cmdline.OptionValue("critical"))
	}

	if cmdline.IsOptionSet("index") {
		aI.createIndex(cmdline.OptionValue("index"))
	}

	if !checkRequiredFlags(*aI.Hostname, *aI.Port, *aI.Username, *aI.Password) {
		os.Exit(exitUnknown)
	}

	switch *aI.Method {
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
	case "smart_thermometer":
		CheckSmartThermometer(aI)
	case "smart_socketpower":
		CheckSmartSocketPower(aI)
	case "smart_socketenergy":
		CheckSmartSocketEnergy(aI)
	default:
		fmt.Println("Unknown method.")
		GlobalReturnCode = exitUnknown
	}

	os.Exit(GlobalReturnCode)
}
