package splyr

import "time"

type SpotifyToken struct {
	AccessToken                      string
	AccessTokenExpirationTimestampMs int64
}

type SpotifyLyrics struct {
	Lines        []SpotifyLyricLine `json:"lines"`
	Language     string             `json:"language"`
	Alternatives []string           `json:"alternatives"`
}

type SpotifyLyricLine struct {
	StartTimeMs string `json:"start_time_ms"`
	Words       string `json:"words"`
}

func (spotifyToken SpotifyToken) IsTokenExpire() bool {
	expirationTime := time.Unix(0, spotifyToken.AccessTokenExpirationTimestampMs*int64(time.Millisecond))
	return time.Now().After(expirationTime)
}
