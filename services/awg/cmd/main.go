package main

import (
	getenv "awg-service/internal/getEnv"
	"awg-service/internal/transport"
	"os"
	"path/filepath"

	awgctrlgo "github.com/OgayOleg/awgctrl-go"
)

var (
	DefaultHTTP   = ""
	DefaultDEVICE = ""
	DefaultAWG    = ""
)

func main() {
	cfg, err := getenv.NewObfuscation()

	if err != nil {
		panic(err)
	}

	tunnelName := getOpt(os.Getenv("DEVICE"), DefaultDEVICE)
	awgEndpoint := getOpt(os.Getenv("AWG_ENDPOINT"), DefaultAWG)
	httpEndpoint := getOpt(os.Getenv("HTTP_ENDPOINT"), DefaultHTTP)

	storagePath, err := filepath.Abs("./data")

	if err != nil {
		panic(err)
	}

	awg, err := awgctrlgo.New(tunnelName, awgEndpoint, storagePath, cfg)

	if err != nil {
		panic(err)
	}
	service := transport.New(awg, storagePath)
	service.Start(httpEndpoint)
}

func getOpt(value string, defaultValue string) string {
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
