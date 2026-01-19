package main

import (
	getenv "awg-service/internal/getEnv"
	"awg-service/internal/transport"
	"os"

	awgctrlgo "github.com/slipynil/awgctrl-go"
)

func main() {
	cfg, err := getenv.NewObfuscation()
	if err != nil {
		panic(err)
	}
	tunnelName, awgEndpoint := os.Getenv("DEVICE"), os.Getenv("AWG_ENDPOINT")
	httpEndpoint := os.Getenv("HTTP_ENDPOINT")
	if tunnelName == "" || awgEndpoint == "" || httpEndpoint == "" {
		panic("DEVICE and AWG_ENDPOINT environment variables are required")
	}
	awg, err := awgctrlgo.New(tunnelName, awgEndpoint, cfg)
	if err != nil {
		panic(err)
	}
	service := transport.New(awg)
	service.Start(httpEndpoint)
}
