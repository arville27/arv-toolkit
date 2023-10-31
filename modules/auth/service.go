package auth

type AuthService interface {
	GenerateAccessToken(username string, password string) (*AccessToken, error)
}
