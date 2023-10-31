package auth

type AccessToken struct {
	AccessToken           string `json:"access_token"`
	ExpirationTimestampMs int64  `json:"expiration_timestamp_ms"`
}
