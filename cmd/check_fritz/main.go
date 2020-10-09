package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// program version
var Version = "1.2.0"

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
	Debug         bool
}

func createRequiredArgumentInformation(hostname string, port string, username string, password string, method string, timeout int, modelgroup string) ArgumentInformation {
	var ai ArgumentInformation

	ai.Hostname = &hostname
	ai.Port = &port
	ai.Username = &username
	ai.Password = &password
	ai.Method = &method
	ai.Modelgroup = &modelgroup
	ai.Timeout = &timeout
	ai.Debug = false

	return ai
}

func (ai *ArgumentInformation) createWarningThreshold(warning float64) {
	ai.Warning = &warning
}

func (ai *ArgumentInformation) createCriticalThreshold(critical float64) {
	ai.Critical = &critical
}

func (ai *ArgumentInformation) createInputVariable(v string) {
	ai.InputVariable = &v
}

func (ai *ArgumentInformation) setDebugMode() {
	ai.Debug = true
}

func printVersion() {
	fmt.Println("check_fritz v" + Version)
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

func checkMain(c *cli.Context) error {

	type param struct {
		Hostname   string
		Port       string
		Username   string
		Password   string
		Method     string
		Timeout    int
		ModelGroup string
	}

	p := param{}

	p.Hostname = c.String("hostname")
	p.Port = c.String("port")
	p.Username = c.String("username")

	if c.IsSet("password") {
		p.Password = c.String("password")
	}

	p.Method = c.String("method")
	p.Timeout = c.Int("timeout")
	p.ModelGroup = c.String("modelgroup")

	argInfo := createRequiredArgumentInformation(p.Hostname, p.Port, p.Username, p.Password, p.Method, p.Timeout, p.ModelGroup)

	if c.IsSet("warning") {
		argInfo.createWarningThreshold(c.Float64("warning"))
	}

	if c.IsSet("critical") {
		argInfo.createCriticalThreshold(c.Float64("critical"))
	}

	if c.IsSet("ain") {
		argInfo.createInputVariable(c.String("ain"))
	}

	if c.IsSet("debug") {
		argInfo.setDebugMode()
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

	os.Exit(GlobalReturnCode)

	return nil
}

func main() {
	app := &cli.App{
		Action:  checkMain,
		Name:    "check_fritz",
		Usage:   "Check plugin to monitor a Fritz!Box",
		Version: Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "hostname",
				Aliases:     []string{"H"},
				Value:       "fritz.box",
				DefaultText: "fritz.box",
				Usage:       "Specifies the hostname.",
			},
			&cli.StringFlag{
				Name:        "port",
				Aliases:     []string{"P"},
				Value:       "49443",
				DefaultText: "49443",
				Usage:       "Specifies the SSL port.",
			},
			&cli.StringFlag{
				Name:        "username",
				Aliases:     []string{"u"},
				Value:       "dslf-config",
				DefaultText: "dslf-config",
				Usage:       "Specifies the username.",
			},
			&cli.StringFlag{
				Name:    "password",
				Aliases: []string{"p"},
				Usage:   "Specifies the password.",
			},
			&cli.StringFlag{
				Name:        "method",
				Aliases:     []string{"m"},
				Value:       "connection_status",
				DefaultText: "connection_status",
				Usage:       "Specifies the check method.",
			},
			&cli.StringFlag{
				Name:    "ain",
				Aliases: []string{"a"},
				Usage:   "Specifies the AIN for smart devices.",
			},
			&cli.IntFlag{
				Name:        "timeout",
				Aliases:     []string{"t"},
				Value:       90,
				DefaultText: "90",
				Usage:       "Specifies the timeout for requests.",
			},
			&cli.StringFlag{
				Name:        "modelgroup",
				Value:       "DSL",
				DefaultText: "DSL",
				Aliases:     []string{"M"},
				Usage:       "Specifies the Fritz!Bpx model group (DSL or Cable).",
			},
			&cli.Float64Flag{
				Name:    "warning",
				Aliases: []string{"w"},
				Usage:   "Specifies the warning threshold.",
			},
			&cli.Float64Flag{
				Name:    "critical",
				Aliases: []string{"c"},
				Usage:   "Specifies the critical threshold.",
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "Outputs debug information",
			},
		},
	}

	cli.AppHelpTemplate = `NAME:

	{{.Name}} - {{.Usage}}

USAGE:
   check_fritz [options...]

OPTIONS:

	{{range .VisibleFlags}}{{.}}
	{{end}}

METHODS:
	connection_status       WAN connection status,
	connection_uptime       WAN connection uptime (in seconds),
	device_uptime           device uptime (in seconds),
	device_update           update state,
	downstream_max          maximum downstream,
	upstream_max            maximum downstream,
	downstream_current      current downstream,
	upstream_current        current upstream,
	downstream_usage        current downstream usage,
	upstream_usage          current upstream usage,
	smart_heatertemperatur  current temperature of a a radiator thermostat (requires AIN),
	smart_socketpower       current power consumption of a socket switch (requires AIN),
	smart_socketenergy 		Total power consumption of the last year of a socket switch (requires AIN),	
	smart_status            current smart device status (requires AIN)
`

	cli.VersionFlag = &cli.BoolFlag{
		Name: "version", Aliases: []string{"V"},
		Usage: "print the version",
	}

	app.Run(os.Args)
}
