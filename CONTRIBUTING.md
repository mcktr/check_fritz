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

The `Makefile` contains the target `build` which will build binaries for the following plattforms. The target will create 
a subdirectory called `build`, the binaries will be created there. The `build` target also creats sha256 checksums for 
every binary.

* linux/amd64
* linux/arm64
* linux/arm
* windows/amd64
* darwin/amd64

The default target `all` will first remove the `build` subdirectory if present, the second step is to build the binaries,
the third and last step will run `go test` on the project directory. Currently there are not that many tests, this will
improved with further development.

The `Makefile` is also used in our CI using [GitHub Actions](https://github.com/mcktr/check_fritz/actions).