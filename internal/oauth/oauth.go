package oauth

import (
	"context"
	"fmt"
	"github.com/veremchukvv/stonks-test/internal/config"
	"github.com/veremchukvv/stonks-test/pkg/logging"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/vk"
	"io/ioutil"
	"net/http"
)

//var VKconfig *oauth2.Config

func GetOauthVKConfig() *oauth2.Config {
	cfg, _ := config.GetConfig()
	return &oauth2.Config{
		ClientID:     cfg.OAuth.VkClientID,
		ClientSecret: cfg.OAuth.VkClientSecret,
		RedirectURL:  cfg.OAuth.VkRedirectURL,
		Scopes:       []string{""},
		Endpoint:     vk.Endpoint,
	}
}

func GetOauthGoogleConfig() *oauth2.Config {
	cfg, _ := config.GetConfig()
	return &oauth2.Config{
		ClientID:     cfg.OAuth.GoogleClientID,
		ClientSecret: cfg.OAuth.GoogleClientSecret,
		RedirectURL:  cfg.OAuth.GoogleRedirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

func GetRandomState() string {
	//TODO randomize state
	return "blabla1"
}

func GetUserVKInfo(ctx context.Context, state string, oauthState string, code string, conf *oauth2.Config) ([]byte, error) {
	log := logging.FromContext(ctx)

	if state != oauthState {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := conf.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	newurl := "https://api.vk.com/method/getProfiles?v=5.131&access_token=" + token.AccessToken
	log.Info(newurl)
	response, err := http.Get(newurl)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
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
		log.Infof("code exchange failed: %s", err)
		//return nil, fmt.Errorf("code exchange failed: %s", err.Error())
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

	//url := conf.AuthCodeURL(oauthState, oauth2.AccessTypeOffline, oauth2.ApprovalForce)

	url := "https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken

	log.Info(url)

	//client := oauth2.NewClient(ctx, tokenSource)

	//response, err := client.Get(url)
	response, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	log.Info(string(contents))
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}
