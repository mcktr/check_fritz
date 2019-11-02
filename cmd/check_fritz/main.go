package main

import (
	"fmt"
	"os"
	"strconv"

	cmdline "github.com/galdor/go-cmdline"
)

// program version
var version = "1.1.0"

// internal exit codes
const (
	exitOk       = 0
	exitWarning  = 1
	exitCritical = 2
	exitUnknown  = 3
)

// GlobalReturnCode holds always the last set return code
var GlobalReturnCode = exitUnknown

// ArgumentInformation is the data structure for the passed arguments
type ArgumentInformation struct {
	Hostname      *string
	Port          *string
	Username      *string
	Password      *string
	Method        *string
	Warning       *float64
	Critical      *float64
	Index         *string
	InputVariable *string
	Timeout       *int
	Modelgroup    *string
}

func createRequiredArgumentInformation(hostname string, port string, username string, password string, method string, timeout string, modelgroup string) ArgumentInformation {
	var ai ArgumentInformation

	ai.Hostname = &hostname
	ai.Port = &port
	ai.Username = &username
	ai.Password = &password
	ai.Method = &method
	ai.Modelgroup = &modelgroup

	ai.createTimeout(timeout)

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
	ai.Index = &index
}

func (ai *ArgumentInformation) createInputVariable(v string) {
	ai.InputVariable = &v
}

func (ai *ArgumentInformation) createTimeout(t string) {
	timeout, err := strconv.Atoi(t)

	if HandleError(err) {
		return
	}

	ai.Timeout = &timeout
}

func printVersion() {
	fmt.Println("check_fritz v" + version)
	GlobalReturnCode = exitOk
}

func checkRequiredFlags(aI *ArgumentInformation) bool {
	if aI.Hostname == nil || *aI.Hostname == "" {
		fmt.Println("No hostname")
		return false
	}

	if aI.Port == nil || *aI.Port == "" {
		fmt.Println("No port")
		return false
	}

	if aI.Username == nil || *aI.Username == "" {
		fmt.Println("No username")
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

	cmdline.AddOption("H", "hostname", "value", "Specifies the hostname.")
	cmdline.AddOption("P", "port", "value", "Specifies the SSL port.")
	cmdline.AddOption("u", "username", "value", "Specifies the username.")
	cmdline.AddOption("p", "password", "value", "Specifies the password.")
	cmdline.AddOption("m", "method", "value", "Specifies the check method.")
	cmdline.AddOption("w", "warning", "value", "Specifies the warning threshold.")
	cmdline.AddOption("c", "critical", "value", "Specifies the critical threshold.")
	cmdline.AddOption("i", "index", "value", "DEPRECATED: Specifies the index.")
	cmdline.AddOption("a", "ain", "value", "Specifies the AIN for smart devices.")
	cmdline.AddOption("t", "timeout", "value", "Specifies the timeout for the request.")
	cmdline.AddOption("M", "modelgroup", "value", "Specifies the Fritz!Box model group (DSL or Cable).")

	cmdline.AddFlag("V", "version", "Returns the version")

	cmdline.SetOptionDefault("hostname", "fritz.box")
	cmdline.SetOptionDefault("port", "49443")
	cmdline.SetOptionDefault("username", "dslf-config")
	cmdline.SetOptionDefault("method", "connection_status")
	cmdline.SetOptionDefault("timeout", "90")
	cmdline.SetOptionDefault("modelgroup", "DSL")

	cmdline.Parse(os.Args)

	if cmdline.IsOptionSet("version") {
		printVersion()
	} else {

		hostname := cmdline.OptionValue("hostname")
		port := cmdline.OptionValue("port")
		username := cmdline.OptionValue("username")
		password := cmdline.OptionValue("password")
		method := cmdline.OptionValue("method")
		timeout := cmdline.OptionValue("timeout")
		modelgroup := cmdline.OptionValue("modelgroup")

		aI := createRequiredArgumentInformation(hostname, port, username, password, method, timeout, modelgroup)

		if cmdline.IsOptionSet("warning") {
			aI.createWarningThreshold(cmdline.OptionValue("warning"))
		}

		if cmdline.IsOptionSet("critical") {
			aI.createCriticalThreshold(cmdline.OptionValue("critical"))
		}

		if cmdline.IsOptionSet("index") {
			aI.createIndex(cmdline.OptionValue("index"))
		}

		if cmdline.IsOptionSet("ain") {
			aI.createInputVariable(cmdline.OptionValue("ain"))
		}

		if !checkRequiredFlags(&aI) {
			os.Exit(exitUnknown)
		}

		switch *aI.Method {
		case "connection_status":
			CheckConnectionStatus(aI)
		case "connection_uptime":
			CheckConnectionUptime(aI)
		case "device_uptime":
			CheckDeviceUptime(aI)
		case "device_update":
			CheckDeviceUpdate(aI)
		case "downstream_max":
			CheckDownstreamMax(aI)
		case "upstream_max":
			CheckUpstreamMax(aI)
		case "downstream_current":
			CheckDownstreamCurrent(aI)
		case "downstream_usage":
			CheckDownstreamUsage(aI)
		case "upstream_usage":
			CheckUpstreamUsage(aI)
		case "upstream_current":
			CheckUpstreamCurrent(aI)
		case "smart_heatertemperatur":
			if cmdline.IsOptionSet("index") {
				CheckSmartHeaterTemperatur(aI)
			} else {
				CheckSpecificSmartHeaterTemperatur(aI)
			}
		case "smart_socketpower":
			if cmdline.IsOptionSet("index") {
				CheckSmartSocketPower(aI)
			} else {
				CheckSpecificSmartSocketPower(aI)
			}
		case "smart_socketenergy":
			if cmdline.IsOptionSet("index") {
				CheckSmartSocketEnergy(aI)
			} else {
				CheckSpecificSmartSocketEnergy(aI)
			}
		case "smart_status":
			if cmdline.IsOptionSet("index") {
				CheckSmartStatus(aI)
			} else {
				CheckSpecificSmartStatus(aI)
			}
		default:
			fmt.Println("Unknown method.")
			GlobalReturnCode = exitUnknown
		}
	}
	os.Exit(GlobalReturnCode)
}
