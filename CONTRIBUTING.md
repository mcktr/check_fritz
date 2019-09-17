# Contributing

There are many ways to contribute, you can fork the project and enhance the code base, submitting bug reports or feature
requests and you can help by testing open pull requests. Any help is much appreciated.

For helping on this project it is advantageous to have a Fritz!Box at hand to test.

## Development

This check plugin is written in Go. 

It is recommendable to use an editor with helpers for Go e.g. auto-completion. For instance [Visual Studio Code](https://code.visualstudio.com/) provides
such helper features with the [Go](https://code.visualstudio.com/docs/languages/go) extension. 

You can find the manufacturer documentation [here](https://avm.de/service/schnittstellen/). Have a look into the PDF 
files and learn which data a Fritz!Box can provide.

* The `cmd/check_fritz` folder holds the main package, where all functions are defined to process data received  from a 
Fritz!Box, as well as functions to define the CLI parameters and process those.

* The `modules/fritz` folder holds the `fritz` package, which implements functions to connect, receive data and to 
authenticate against a Fritz!Box.

* The `modules/perfdata` folder holds the `perfdata` package, which implements functions to generate performance data for
check results.

* The `modules/thresholds` holds the `thresholds` package, which implements function to define a warning and/or critical
threshold for a check function.

## Building

The `Makefile` contains various build definitions for Linux, Windows and MacOS. The default target is `all` which will 
generate a binary for every defined platform. Available targets are:

| Target                | Descriptions                            |
|-----------------------|-----------------------------------------|
| `all`                 | Builds for all defined platforms.       |
| `build.linux`         | Builds for all Linux based platforms.   |
| `build.windows`       | Builds for all Windows based platforms. |
| `build.darwin`        | Builds for all Darwin based platforms.  |
| `build.linux.amd64`   | Builds for Linux amd64 architecture.    |
| `build.linux.arm64`   | Builds for Linux arm64 architecture.    |
| `build.linux.arm5`    | Builds for Linux arm5 architecture.     |
| `build.linux.arm6`    | Builds for Linux arm6 architecture.     |
| `build.linux.arm7`    | Builds for Linux arm7 architecture.     |
| `build.windows.amd64` | Builds for Windows amd64 architecture.  |
| `build.darwin.amd64`  | Builds for Darwin amd64 architecture.   |

The binary will be generated with the following naming convention: `check_fritz.$OS.$ARCH`. After the build is finished
a sha256 checksum will be generated for every binary. 

The Makefile script is also used for release builds using TravisCI.