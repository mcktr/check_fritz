package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli"
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
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "hostname, H",
			Value: "fritz.box",
			Usage: "Specifies the hostname.",
		},
		cli.StringFlag{
			Name:  "port, P",
			Value: "49443",
			Usage: "Specifies the SSL port.",
		},
		cli.StringFlag{
			Name:  "username, u",
			Value: "dslf-config",
			Usage: "Specifies the username.",
		},
		cli.StringFlag{
			Name:  "password, p",
			Usage: "Specifies the password.",
		},
		cli.StringFlag{
			Name:  "method, m",
			Value: "connection_status",
			Usage: "Specifies the check method.",
		},
		cli.StringFlag{
			Name:  "warning, w",
			Usage: "Specifies the warning threshold.",
		},
		cli.StringFlag{
			Name:  "critical, c",
			Usage: "Specifies the critical threshold.",
		},
		cli.StringFlag{
			Name:  "ain, a",
			Usage: "Specifies the AIN for smart devices.",
		},
		cli.StringFlag{
			Name:  "timeout, t",
			Value: "90",
			Usage: "Specifies the timeout for the request.",
		},
		cli.StringFlag{
			Name:  "modelgroup, M",
			Value: "DSL",
			Usage: "Specifies the Fritz!Box model group (DSL or Cable).",
		},
	}

	app.Action = func(c *cli.Context) error {

		hostname := c.String("hostname")
		port := c.String("port")
		username := c.String("username")
		password := c.String("password")
		method := c.String("method")
		timeout := c.String("timeout")
		modelgroup := c.String("modelgroup")

		argInfo := createRequiredArgumentInformation(hostname, port, username, password, method, timeout, modelgroup)

		if c.IsSet("warning") {
			argInfo.createWarningThreshold(c.String("warning"))
		}

		if c.IsSet("critical") {
			argInfo.createCriticalThreshold(c.String("critical"))
		}

		if c.IsSet("ain") {
			argInfo.createInputVariable(c.String("ain"))
		}

		if !checkRequiredFlags(&argInfo) {
			os.Exit(exitUnknown)
		}

		switch *argInfo.Method {
		case "connection_status":
			CheckConnectionStatus(argInfo)
		case "connection_uptime":
			CheckConnectionUptime(argInfo)
		case "device_uptime":
			CheckDeviceUptime(argInfo)
		case "device_update":
			CheckDeviceUpdate(argInfo)
		case "downstream_max":
			CheckDownstreamMax(argInfo)
		case "upstream_max":
			CheckUpstreamMax(argInfo)
		case "downstream_current":
			CheckDownstreamCurrent(argInfo)
		case "downstream_usage":
			CheckDownstreamUsage(argInfo)
		case "upstream_usage":
			CheckUpstreamUsage(argInfo)
		case "upstream_current":
			CheckUpstreamCurrent(argInfo)
		case "smart_heatertemperatur":
			CheckSpecificSmartHeaterTemperatur(argInfo)
		case "smart_socketpower":
			CheckSpecificSmartSocketPower(argInfo)
		case "smart_socketenergy":
			CheckSpecificSmartSocketEnergy(argInfo)
		case "smart_status":
			CheckSpecificSmartStatus(argInfo)
		default:
			fmt.Println("Unknown method.")
			GlobalReturnCode = exitUnknown
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}

	os.Exit(GlobalReturnCode)
}
