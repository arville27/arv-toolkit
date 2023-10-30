package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	RestServerIp          string
	RestServerPort        uint16
	SpotifySPDCCookie     string
	SpotifyTokenCachePath string
}

func LoadConfigFromEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load config, %v", err)
	}

	portEnv := os.Getenv("REST_SERVER_PORT")
	port := uint64(8080)
	if len(portEnv) != 0 {
		port, err = strconv.ParseUint(portEnv, 16, 16)
		if err != nil {
			log.Fatalf("Failed to parse REST_SERVER_PORT environment variable, received '%s'", portEnv)
		}
	}

	tokenCachePath := os.Getenv("SPLYR_TOKEN_CACHE_PATH")
	if len(tokenCachePath) == 0 {
		tokenCachePath = "."
	}
	return &Config{
		SpotifySPDCCookie:     os.Getenv("SPLYR_SP_DC"),
		RestServerIp:          os.Getenv("REST_SERVER_IP"),
		RestServerPort:        uint16(port),
		SpotifyTokenCachePath: tokenCachePath,
	}
}
