.PHONY : all
all : build.linux build.windows build.darwin

build.linux: build.linux.amd64 build.linux.arm64 build.linux.arm5 build.linux.arm6 build.linux.arm7
build.windows: build.windows.amd64
build.darwin: build.darwin.amd64

build.linux.amd64:
	GOOS=linux GOARCH=amd64 go build -a -v -o "check_fritz.linux.amd64" ./cmd/check_fritz/
	sha256sum check_fritz.linux.amd64 > check_fritz.linux.amd64.sha256

build.linux.arm64:
	GOOS=linux GOARCH=arm64 go build -a -v -o "check_fritz.linux.arm64" ./cmd/check_fritz/
	sha256sum check_fritz.linux.arm64 > check_fritz.linux.arm64.sha256

build.linux.arm5:
	GOOS=linux GOARCH=arm GOARM=5 go build -a -v -o "check_fritz.linux.arm5" ./cmd/check_fritz/
	sha256sum check_fritz.linux.arm5 > check_fritz.linux.arm5.sha256

build.linux.arm6:
	GOOS=linux GOARCH=arm GOARM=6 go build -a -v -o "check_fritz.linux.arm6" ./cmd/check_fritz/
	sha256sum check_fritz.linux.arm6 > check_fritz.linux.arm6.sha256

build.linux.arm7:
	GOOS=linux GOARCH=arm GOARM=7 go build -a -v -o "check_fritz.linux.arm7" ./cmd/check_fritz/
	sha256sum check_fritz.linux.arm7 > check_fritz.linux.arm7.sha256

build.windows.amd64:
	GOOS=windows GOARCH=amd64 go build -a -v -o "check_fritz.windows.amd64.exe" ./cmd/check_fritz/
	sha256sum check_fritz.windows.amd64.exe > check_fritz.windows.amd64.exe.sha256

build.darwin.amd64:
	GOOS=darwin GOARCH=amd64 go build -a -v -o "check_fritz.darwin.amd64" ./cmd/check_fritz/
	sha256sum check_fritz.darwin.amd64 > check_fritz.darwin.amd64.sha256
