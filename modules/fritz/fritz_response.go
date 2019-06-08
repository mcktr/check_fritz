package fritz

import (
	"encoding/xml"
	"time"
)

// TR064Response is the overriding data structure for the SOAP responses
type TR064Response interface {
}

// WANPPPConnectionResponse is the data structure for responses from WANPPPConnection
type WANPPPConnectionResponse struct {
	TR064Response
	NewEnable                  string `xml:"Body>GetInfoResponse>NewEnable"`
	NewConnectionStatus        string `xml:"Body>GetInfoResponse>NewConnectionStatus"`
	NewPossibleConnectionTypes string `xml:"Body>GetInfoResponse>NewPossibleConnectionTypes"`
	NewConnectionType          string `xml:"Body>GetInfoResponse>NewConnectionType"`
	NewUptime                  string `xml:"Body>GetInfoResponse>NewUptime"`
	NewUpstreamMaxBitRate      string `xml:"Body>GetInfoResponse>NewUpstreamMaxBitRate"`
	NewDownstreamMaxBitRate    string `xml:"Body>GetInfoResponse>NewDownstreamMaxBitRate"`
	NewLastConnectionError     string `xml:"Body>GetInfoResponse>NewLastConnectionError"`
	NewIdleDisconnectTime      string `xml:"Body>GetInfoResponse>NewIdleDisconnectTime"`
	NewRSIPAvailable           string `xml:"Body>GetInfoResponse>NewRSIPAvailable"`
	NewUserName                string `xml:"Body>GetInfoResponse>NewUserName"`
	NewNATEnabled              string `xml:"Body>GetInfoResponse>NewNATEnabled"`
	NewExternalIPAddress       string `xml:"Body>GetInfoResponse>NewExternalIPAddress"`
	NewDNSServers              string `xml:"Body>GetInfoResponse>NewDNSServers"`
	NewMACAddress              string `xml:"Body>GetInfoResponse>NewMACAddress"`
	NewConnectionTrigger       string `xml:"Body>GetInfoResponse>NewConnectionTrigger"`
	NewLastAuthErrorInfo       string `xml:"Body>GetInfoResponse>NewLastAuthErrorInfo"`
	NewMaxCharsUsername        string `xml:"Body>GetInfoResponse>NewMaxCharsUsername"`
	NewMinCharsUsername        string `xml:"Body>GetInfoResponse>NewMinCharsUsername"`
	NewAllowedCharsUsername    string `xml:"Body>GetInfoResponse>NewAllowedCharsUsername"`
	NewMaxCharsPassword        string `xml:"Body>GetInfoResponse>NewMaxCharsPassword"`
	NewMinCharsPassword        string `xml:"Body>GetInfoResponse>NewMinCharsPassword"`
	NewAllowedCharsPassword    string `xml:"Body>GetInfoResponse>NewAllowedCharsPassword"`
	NewTransportType           string `xml:"Body>GetInfoResponse>NewTransportType"`
	NewRouteProtocolRx         string `xml:"Body>GetInfoResponse>NewRouteProtocolRx"`
	NewPPPoEServiceName        string `xml:"Body>GetInfoResponse>NewPPPoEServiceName"`
	NewRemoteIPAddress         string `xml:"Body>GetInfoResponse>NewRemoteIPAddress"`
	NewPPPoEACName             string `xml:"Body>GetInfoResponse>NewPPPoEACName"`
	NewDNSEnabled              string `xml:"Body>GetInfoResponse>NewDNSEnabled"`
	NewDNSOverrideAllowed      string `xml:"Body>GetInfoResponse>NewDNSOverrideAllowed"`
}

// DeviceInfoResponse is the data structure for responses from DeviceInfo
type DeviceInfoResponse struct {
	TR064Response
	NewManufacturerName string `xml:"Body>GetInfoResponse>NewManufacturerName"`
	NewManufacturerOUI  string `xml:"Body>GetInfoResponse>NewManufacturerOUI"`
	NewModelName        string `xml:"Body>GetInfoResponse>NewModelName"`
	NewDescription      string `xml:"Body>GetInfoResponse>NewDescription"`
	NewProductClass     string `xml:"Body>GetInfoResponse>NewProductClass"`
	NewSerialNumber     string `xml:"Body>GetInfoResponse>NewSerialNumber"`
	NewSoftwareVersion  string `xml:"Body>GetInfoResponse>NewSoftwareVersion"`
	NewHardwareVersion  string `xml:"Body>GetInfoResponse>NewHardwareVersion"`
	NewSpecVersion      string `xml:"Body>GetInfoResponse>NewSpecVersion"`
	NewProvisioningCode string `xml:"Body>GetInfoResponse>NewProvisioningCode"`
	NewUpTime           string `xml:"Body>GetInfoResponse>NewUpTime"`
	NewDeviceLog        string `xml:"Body>GetInfoResponse>NewDeviceLog"`
}

// UserInterfaceInfoResponse is the data structure for responses from UserInterface
type UserInterfaceInfoResponse struct {
	TR064Response
	NewUpgradeAvailable       string `xml:"Body>GetInfoResponse>NewUpgradeAvailable"`
	NewPasswordRequired       string `xml:"Body>GetInfoResponse>NewPasswordRequired"`
	NewPasswordUserSelectable string `xml:"Body>GetInfoResponse>NewPasswordUserSelectable"`
	NewWarrantyDate           string `xml:"Body>GetInfoResponse>NewWarrantyDate"`
	NewXAVMDEVersion          string `xml:"Body>GetInfoResponse>NewX_AVM-DE_Version"`
	NewXAVMDEDownloadURL      string `xml:"Body>GetInfoResponse>NewX_AVM-DE_DownloadURL"`
	NewXAVMDEInfoURL          string `xml:"Body>GetInfoResponse>NewX_AVM-DE_InfoURL"`
	NewXAVMDEUpdateState      string `xml:"Body>GetInfoResponse>NewX_AVM-DE_UpdateState"`
	NewXAVMDELaborVersion     string `xml:"Body>GetInfoResponse>NewX_AVM-DE_LaborVersion"`
}

// WANCommonInterfaceOnlineMonitorResponse is the data structure for responses from WANCommonInterfaceConfig
type WANCommonInterfaceOnlineMonitorResponse struct {
	TR064Response
	NewTotalNumberSyncGroups string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>NewTotalNumberSyncGroups"`
	NewSyncGroupName         string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>NewSyncGroupName"`
	NewSyncGroupMode         string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>NewSyncGroupMode"`
	NewMaxDS                 string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newmax_ds"`
	NewMaxUS                 string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newmax_us"`
	NewDSCurrentBPS          string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newds_current_bps"`
	NewMCCurrentBPS          string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newmc_current_bps"`
	NewUSCurrentBPS          string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newus_current_bps"`
	NewPrioRealtimeBPS       string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newprio_realtime_bps"`
	NewPrioHighBPS           string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newprio_high_bps"`
	NewPrioDefaultBPS        string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newprio_default_bps"`
	NewPrioLowBPS            string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newprio_low_bps"`
}

// SmartDeviceInfoResponse is the data structure for responses from X_AVM-DE_Homeauto
type SmartDeviceInfoResponse struct {
	TR064Response
	NewAIN                    string `xml:"Body>GetGenericDeviceInfosResponse>NewAIN"`
	NewDeviceID               string `xml:"Body>GetGenericDeviceInfosResponse>NewDeviceId"`
	NewFunctionBitMask        string `xml:"Body>GetGenericDeviceInfosResponse>NewFunctionBitMask"`
	NewFirmwareVersion        string `xml:"Body>GetGenericDeviceInfosResponse>NewFirmwareVersion"`
	NewManufacturer           string `xml:"Body>GetGenericDeviceInfosResponse>NewManufacturer"`
	NewProductName            string `xml:"Body>GetGenericDeviceInfosResponse>NewProductName"`
	NewDeviceName             string `xml:"Body>GetGenericDeviceInfosResponse>NewDeviceName"`
	NewPresent                string `xml:"Body>GetGenericDeviceInfosResponse>NewPresent"`
	NewMultimeterIsEnabled    string `xml:"Body>GetGenericDeviceInfosResponse>NewMultimeterIsEnabled"`
	NewMultimeterIsValid      string `xml:"Body>GetGenericDeviceInfosResponse>NewMultimeterIsValid"`
	NewMultimeterPower        string `xml:"Body>GetGenericDeviceInfosResponse>NewMultimeterPower"`
	NewMultimeterEnergy       string `xml:"Body>GetGenericDeviceInfosResponse>NewMultimeterEnergy"`
	NewTemperatureIsEnabled   string `xml:"Body>GetGenericDeviceInfosResponse>NewTemperatureIsEnabled"`
	NewTemperatureIsValid     string `xml:"Body>GetGenericDeviceInfosResponse>NewTemperatureIsValid"`
	NewTemperatureCelsius     string `xml:"Body>GetGenericDeviceInfosResponse>NewTemperatureCelsius"`
	NewTemperatureOffset      string `xml:"Body>GetGenericDeviceInfosResponse>NewTemperatureOffset"`
	NewSwitchIsEnabled        string `xml:"Body>GetGenericDeviceInfosResponse>NewSwitchIsEnabled"`
	NewSwitchIsValid          string `xml:"Body>GetGenericDeviceInfosResponse>NewSwitchIsValid"`
	NewSwitchState            string `xml:"Body>GetGenericDeviceInfosResponse>NewSwitchState"`
	NewSwitchMode             string `xml:"Body>GetGenericDeviceInfosResponse>NewSwitchMode"`
	NewSwitchLock             string `xml:"Body>GetGenericDeviceInfosResponse>NewSwitchLock"`
	NewHkrIsEnabled           string `xml:"Body>GetGenericDeviceInfosResponse>NewHkrIsEnabled"`
	NewHkrIsValid             string `xml:"Body>GetGenericDeviceInfosResponse>NewHkrIsValid"`
	NewHkrIsTemperature       string `xml:"Body>GetGenericDeviceInfosResponse>NewHkrIsTemperature"`
	NewHkrSetVentilStatus     string `xml:"Body>GetGenericDeviceInfosResponse>NewHkrSetVentilStatus"`
	NewHkrSetTemperature      string `xml:"Body>GetGenericDeviceInfosResponse>NewHkrSetTemperature"`
	NewHkrReduceVentilStatus  string `xml:"Body>GetGenericDeviceInfosResponse>NewHkrReduceVentilStatus"`
	NewHkrReduceTemperature   string `xml:"Body>GetGenericDeviceInfosResponse>NewHkrReduceTemperature"`
	NewHkrComfortVentilStatus string `xml:"Body>GetGenericDeviceInfosResponse>NewHkrComfortVentilStatus"`
	NewHkrComfortTemperature  string `xml:"Body>GetGenericDeviceInfosResponse>NewHkrComfortTemperature"`
}

// UnmarshalSoapResponse unmarshals the soap response to the data structure
func UnmarshalSoapResponse(resp TR064Response, inputXML [][]byte) error {
	for i := range inputXML {

		err := xml.Unmarshal(inputXML[i], &resp)

		if err != nil {
			return err
		}
	}
	return nil
}

// ProcessSoapResponse handles the SOAP response from channels
func ProcessSoapResponse(resps chan []byte, errs chan error, count int) ([][]byte, error) {
	results := make([][]byte, 0)

	for {
		timedout := false

		select {
		case err := <-errs:
			if err != nil {
				return results, err
			}
		case res := <-resps:
			count--
			results = append(results, res)

			if count <= 0 {
				break
			}
		case <-time.After(60 * time.Second):
			// TODO: Timeout
			panic("Timeout")
		}

		if count <= 0 || timedout {
			break
		}
	}

	return results, nil
}
