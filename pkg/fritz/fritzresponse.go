package fritz

type GetInfoResponse struct {
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
