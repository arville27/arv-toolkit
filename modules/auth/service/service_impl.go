package service

import (
	"arville27/arv-toolkit/modules/auth"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	SigningKeySecret string
	ValidCredentials *map[string]string
}

func NewAuthService(signingKeySecret string, validCredentials *map[string]string) auth.AuthService {
	return &service{SigningKeySecret: signingKeySecret, ValidCredentials: validCredentials}
}

func (s service) GenerateAccessToken(username string, password string) (*auth.AccessToken, error) {
	isValid := false
	for validUsername, validPassword := range *s.ValidCredentials {
		if username == validUsername && password == validPassword {
			isValid = true
			break
		}
	}

	if !isValid {
		return nil, auth.NewAuthError("Invalid username or password", auth.InvalidCredential, nil)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	expTimestamp := time.Now().Add(time.Hour * 72).Unix()
	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = username
	claims["exp"] = expTimestamp

	accessToken, err := token.SignedString([]byte(s.SigningKeySecret))
	if err != nil {
		slog.Error(
			"Failed generate new access token",
			slog.String("error", err.Error()),
		)
		return nil, err
	}

	slog.Info("Successfully generate an access token")

	return &auth.AccessToken{AccessToken: accessToken, ExpirationTimestampMs: expTimestamp}, nil
}
