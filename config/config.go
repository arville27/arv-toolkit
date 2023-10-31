package config

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	RestServerIp          string
	RestServerPort        uint16
	SpotifySPDCCookie     string
	SpotifyTokenCachePath string
	AuthTokenSecret       string
	ValidCredentials      *map[string]string
}

func LoadConfigFromEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Failed to load config, %v", err)
	}

	tokenSecret := os.Getenv("AUTH_TOKEN_SECRET")
	if len(tokenSecret) == 0 {
		log.Fatal("Please ensure 'AUTH_TOKEN_SECRET' present in environment variable")
	}

	rawCredentials := os.Getenv("AUTH_VALID_CREDENTIALS")
	validCredentials, err := parseCredentials(rawCredentials)
	if err != nil {
		log.Fatalf("Please ensure 'AUTH_VALID_CREDENTIALS' has valid format, error due to %s", err)
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

	log.Printf("Valid credentials: %s", *validCredentials)
	return &Config{
		SpotifySPDCCookie:     os.Getenv("SPLYR_SP_DC"),
		RestServerIp:          os.Getenv("REST_SERVER_IP"),
		RestServerPort:        uint16(port),
		SpotifyTokenCachePath: tokenCachePath,
		AuthTokenSecret:       tokenSecret,
		ValidCredentials:      validCredentials,
	}
}

func parseCredentials(rawCredentials string) (*map[string]string, error) {
	if len(rawCredentials) == 0 {
		return nil, errors.New("empty credentials")
	}

	listPairCredentials := strings.Split(rawCredentials, ",")

	validCredentials := make(map[string]string)
	for _, rawPairCredential := range listPairCredentials {
		pairCredential := strings.SplitN(rawPairCredential, ";", 2)
		if len(pairCredential) != 2 {
			return nil, errors.New("found invalid credential entry " + rawPairCredential)
		}
		validCredentials[pairCredential[0]] = pairCredential[1]
	}

	return &validCredentials, nil
}
