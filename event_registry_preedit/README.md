# Event Registry API

## Purpose
The purpose of this project is to utilize the Gin framework in Go to create and register for events, using token-based authentication.

## Instructions
To start this project:
1. Make sure a C compiler (GCC) is installed and available for your system. If using Windows, it's recommended to employ WSL Ubuntu. On Ubuntu, the command to do that would be `sudo apt install gcc`
2. Install Go: `sudo apt install golang-go`
2. Set CGO_ENABLED=1 by running `go env -w CGO_ENABLED=1`
4. Run `go run main.go`. You should see the message "Listening and Serving HTTP on :8080"

## Resources
* [Installing Go](https://go.dev/doc/install)
* [Installing WSL2 and Ubuntu](https://documentation.ubuntu.com/wsl/en/stable/howto/install-ubuntu-wsl2/)
* [sqlite3](https://pkg.go.dev/github.com/mattn/go-sqlite3)