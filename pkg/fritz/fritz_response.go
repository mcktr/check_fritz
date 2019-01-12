package fritz

type Response interface {
}

type GetWANPPPConnectionInfoResponse struct {
	Response
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

type GetDeviceInfoResponse struct {
	Response
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

type GetWANCommonInterfaceOnlineMonitorResponse struct {
	Response
	NewTotalNumberSyncGroups string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>NewTotalNumberSyncGroups"`
	NewSyncGroupName         string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>NewSyncGroupName"`
	NewSyncGroupMode         string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>NewSyncGroupMode"`
	Newmax_ds                string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newmax_ds"`
	Newmax_us                string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newmax_us"`
	Newds_current_bps        string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newds_current_bps"`
	Newmc_current_bps        string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newmc_current_bps"`
	Newus_current_bps        string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newus_current_bps"`
	Newprio_realtime_bps     string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newprio_realtime_bps"`
	Newprio_high_bps         string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newprio_high_bps"`
	Newprio_default_bps      string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newprio_default_bps"`
	Newprio_low_bps          string `xml:"Body>X_AVM-DE_GetOnlineMonitorResponse>Newprio_low_bps"`
}

type GetInterfaceInfoResponse struct {
	Response
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
