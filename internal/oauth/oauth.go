package oauth

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/veremchukvv/stonks-test/internal/config"
	"github.com/veremchukvv/stonks-test/pkg/logging"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/vk"
)

func GetOauthVKConfig() *oauth2.Config {
	log := logging.NewLogger(false, "console")

	cfg, err := config.GetConfig()
	if err != nil {
		log.Error("Can't read config")
	}

	return &oauth2.Config{
		ClientID:     cfg.OAuth.VkClientID,
		ClientSecret: cfg.OAuth.VkClientSecret,
		RedirectURL:  cfg.OAuth.VkRedirectURL,
		Scopes:       []string{""},
		Endpoint:     vk.Endpoint,
	}
}

func GetOauthGoogleConfig() *oauth2.Config {
	log := logging.NewLogger(false, "console")

	cfg, err := config.GetConfig()
	if err != nil {
		log.Error("Can't read config")
	}

	return &oauth2.Config{
		ClientID:     cfg.OAuth.GoogleClientID,
		ClientSecret: cfg.OAuth.GoogleClientSecret,
		RedirectURL:  cfg.OAuth.GoogleRedirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

func GetRandomState() string {
	// TODO randomize state
	return "blabla1"
}

func GetUserVKInfo(ctx context.Context, state string, oauthState string, code string, conf *oauth2.Config) ([]byte, error) {
	if state != oauthState {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := conf.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %w", err)
	}

	newurl := "https://api.vk.com/method/getProfiles?v=5.131&access_token=" + token.AccessToken
	// nolint:gosec
	response, err := http.Get(newurl)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %w", err)
	}
	defer func() {
		_ = response.Body.Close()
	}()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %w", err)
	}

	return contents, nil
}

func GetUserGoogleInfo(ctx context.Context, state string, oauthState string, code string, conf *oauth2.Config) ([]byte, error) {
	log := logging.FromContext(ctx)

	if state != oauthState {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Infof("code exchange failed: %v", err)
	}

	tokenSource := conf.TokenSource(ctx, token)

	newToken, err := tokenSource.Token()
	if err != nil {
		log.Info(err)
	}

	if newToken.AccessToken != token.AccessToken {
		token = newToken
		log.Info("Saved new token:", newToken.AccessToken)
	}

	url := "https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken
	// nolint:gosec
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %w", err)
	}
	defer func() {
		_ = response.Body.Close()
	}()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %w", err)
	}
	return contents, nil
}
