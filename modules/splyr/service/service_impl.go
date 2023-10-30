package service

import (
	"arville27/arv-toolkit/modules/splyr"
	"arville27/arv-toolkit/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"log/slog"
)

const tokenUrl = "https://open.spotify.com/get_access_token?reason=transport&productType=web_player"
const lyricsUrl = "https://spclient.wg.spotify.com/color-lyrics/v2/track/"

type service struct {
	client                 *http.Client
	spDc                   string
	tokenResponseCachePath string
}

func NewSpotifyLyricsService(client *http.Client, spDc string, tokenResponseCachePath string) splyr.SpotifyLyricsService {
	return &service{client, spDc, tokenResponseCachePath}
}

func (s *service) GetLyrics(trackId string) (*splyr.SpotifyLyrics, error) {
	tokenResponse, err := s.GetToken()
	if err != nil {
		return nil, err
	}

	slog.Info("Request lyric", slog.String("trackId", trackId))
	formattedUrl := fmt.Sprintf("%s%s?format=json&market=from_token", lyricsUrl, trackId)
	request, err := http.NewRequest(http.MethodGet, formattedUrl, nil)
	if err != nil {
		slog.Error("Failed create a request to get a lyric", "error", err)
		return nil, err
	}

	request.Header = http.Header{
		"User-Agent":    {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.0.0 Safari/537.36"},
		"App-platform":  {"WebPlayer"},
		"Authorization": {fmt.Sprintf("Bearer %s", tokenResponse.AccessToken)},
	}

	response, err := s.client.Do(request)
	if err != nil {
		slog.Error("Failed to get lyrics: %v", err)
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		slog.Error("Failed to get lyrics", "response", response)
		return nil, splyr.SplyrError{Reason: "Failed to get lyrics", Cause: err}
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		slog.Error("Failed read lyrics response", "error", err)
		return nil, err
	}

	var spotifyLyricsResponse spotifyLyricsResponse
	err = json.Unmarshal(responseData, &spotifyLyricsResponse)
	if err != nil {
		slog.Error("Failed deserialize lyrics response", "error", err)
		return nil, err
	}

	spotifyLyrics := &splyr.SpotifyLyrics{
		Lines: func() []splyr.SpotifyLyricLine {
			var lines []splyr.SpotifyLyricLine
			for _, line := range spotifyLyricsResponse.Lyrics.Lines {
				lines = append(lines, splyr.SpotifyLyricLine(line))
			}
			return lines
		}(),
		Language:     spotifyLyricsResponse.Lyrics.Language,
		Alternatives: spotifyLyricsResponse.Lyrics.Alternatives,
	}

	return spotifyLyrics, nil
}

func (s *service) GetToken() (*splyr.SpotifyToken, error) {
	var err error
	if tokenResponse == nil {
		tokenResponse, err = loadTokenResponseCache(s.tokenResponseCachePath)
		if err != nil {
			slog.Error("Cannot load token response cache", "error", err)
		} else {
			slog.Debug("Succesfully load access token response from cache")
		}
	}

	if tokenResponse != nil && !tokenResponse.IsTokenExpire() {
		slog.Debug("Token is not expire, reuse access token")
		return tokenResponse, nil
	}

	slog.Info("Request new access token")
	request, err := http.NewRequest(http.MethodGet, tokenUrl, nil)
	if err != nil {
		slog.Error("Failed create a request to get a token", "error", err)
		return nil, err
	}

	request.Header = http.Header{
		"User-Agent":   {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.0.0 Safari/537.36"},
		"App-platform": {"WebPlayer"},
		"content-type": {"text/html; charset=utf-8"},
		"cookie":       {fmt.Sprintf("sp_dc=%s;", s.spDc)},
	}

	response, err := s.client.Do(request)
	if err != nil {
		slog.Error("Failed to get access token: %v", err)
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		slog.Error("Failed to get access token", "error", response)
		return nil, splyr.SplyrError{Reason: "Failed to get access token", Cause: err}
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		slog.Error("Failed read access token response", "error", err)
		return nil, err
	}
	slog.Debug("Get access token response body", "responseBody", string(responseData))

	var spotifyTokenResponse spotifyTokenResponse
	err = json.Unmarshal(responseData, &spotifyTokenResponse)
	if err != nil {
		slog.Error("Failed deserialize access token response", "error", err)
		return nil, err
	}

	tokenResponse = &splyr.SpotifyToken{
		AccessToken:                      spotifyTokenResponse.AccessToken,
		AccessTokenExpirationTimestampMs: spotifyTokenResponse.AccessTokenExpirationTimestampMs,
	}
	err = saveTokenResponseCache(s.tokenResponseCachePath, tokenResponse)
	if err != nil {
		slog.Error("Failed save access token response to cache", "error", err)
	} else {
		slog.Debug("Succesfully saving access token response to cache")
	}

	return tokenResponse, nil
}

func saveTokenResponseCache(path string, tokenResponse *splyr.SpotifyToken) error {
	var err error
	jsonByte, err := json.Marshal(tokenResponse)
	if err != nil {
		return err
	}

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	err = utils.WriteToFile(filepath.Join(path, "cache"), jsonByte)
	if err != nil {
		return err
	}
	return nil
}

func loadTokenResponseCache(path string) (*splyr.SpotifyToken, error) {
	cachedTokenPath, _ := filepath.Abs(filepath.Join(path, "cache"))

	_, err := os.Stat(cachedTokenPath)
	if err == nil {
	} else if os.IsNotExist(err) {
		return nil, errors.New(cachedTokenPath + " does not exist")
	} else {
		return nil, err
	}

	content, err := utils.ReadFile(cachedTokenPath)
	if err != nil {
		return nil, err
	}
	var cachedTokenResponse splyr.SpotifyToken
	err = json.Unmarshal(content, &cachedTokenResponse)
	if err != nil {
		return nil, err
	}

	return &cachedTokenResponse, nil
}

var tokenResponse *splyr.SpotifyToken

// dto
type spotifyTokenResponse struct {
	ClientId                         string `json:"clientId"`
	AccessToken                      string `json:"accessToken"`
	AccessTokenExpirationTimestampMs int64  `json:"accessTokenExpirationTimestampMs"`
	IsAnonymous                      bool   `json:"isAnonymous"`
}

type spotifyLyricsResponse struct {
	Lyrics struct {
		Lines []struct {
			StartTimeMs string `json:"startTimeMs"`
			Words       string `json:"words"`
		} `json:"lines"`
		Language     string   `json:"language"`
		Alternatives []string `json:"alternatives"`
	} `json:"lyrics"`
}
