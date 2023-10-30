package splyr

type SpotifyLyricsService interface {
	GetLyrics(trackId string) (*SpotifyLyrics, error)
	GetToken() (*SpotifyToken, error)
}
