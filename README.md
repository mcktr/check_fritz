# check_fritz

> **Note:**
>
> This is an early development version. Things are unstable and possibly don't work correctly.
> It is also possible that things change in the final version.

**Table of Contents**

1. Introduction
2. Support
3. Requirements
4. Usage

## Introduction

This is a check plugin, written in Go, for Icinga 2 to monitor a Fritz!Box.

## Support

Please ask your questions in the community channels.

## Requirements

### Fritz!Box configuration

The TR-064 feature must be enabled on the Fritz!Box.

You can enable the feature in the following menu:

```
Heimnetz -> Netzwerk -> Netzwerkeinstellungen ->  Heimnetzfreigaben -> Zugriff für Anwendungen zulassen
```

![Fritz!Box configuration](doc/images/fritzbox-configuration-tr064.png)

## Usage

| Parameter   | Description                                                                                   |
| ----------- | --------------------------------------------------------------------------------------------- |
| `-hostname` | **Optional.** IP-Address or Hostname of the Fritz!Box. Defaults to `fritz.box`.               |
| `-port`     | **Optional.** Port for TR-064 over SSL. Defaults to `49443`.                                  |
| `-username` | **Optional.** Fritz!Box web interface Username for authentication. Defaults to `dslf-config`. |
| `-password` | **Required.** Fritz!Box web interface Password for authentication.                            |
| `-method`   | **Optional.** Defines the used check method. Defaults to `connection_status`.                 |

> **Note:**
>
> If you don't use the authentication method with username and password on your Fritz!Box, leave the username blank.


### Methods

| Name                | Description                       |
| ------------------- | --------------------------------- |
| `connection_status` | WAN connection status.            |
| `connection_uptime` | WAN connection uptime in seconds. |
| `device_uptime`     | Device uptime in seconds.         |
| `downstream_max`    | Maximum Downstream.               |
| `upstream_max`      | Maximum Upstream.                 | 
