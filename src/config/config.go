package config

import (
	"os"
)

const DEFAULT_PORT = "8000"
const DEFAULT_ENDPOINT = "https://api.hitbtc.com/api/2"

var SUPPORTED_SYMBOLS = []string{"BTCUSD", "ETHBTC"}

func GetPort() string {
	if val, ok := os.LookupEnv("HTTP_PORT"); ok {
		return val
	}

	return DEFAULT_PORT
}

func GetHitBTCEndpoint() string {
	if val, ok := os.LookupEnv("HITBTC_ENDPOINT"); ok {
		return val
	}

	return DEFAULT_ENDPOINT
}
