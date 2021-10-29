package oauth

import (
	"context"
	"fmt"
	"github.com/veremchukvv/stonks-test/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
	"io/ioutil"
	"log"
	"net/http"
)

//var VKconfig *oauth2.Config

func GetOauthConfig() *oauth2.Config {
	cfg, _ := config.GetConfig()
	return &oauth2.Config{
		ClientID:     cfg.OAuth.VkClientID,
		ClientSecret: cfg.OAuth.VkClientSecret,
		RedirectURL:  cfg.OAuth.VkRedirectURL,
		Scopes:       []string{""},
		Endpoint:     vk.Endpoint,
}
}

func GetRandomState() string {
	//TODO randomize state
	return "blabla"
}

func GetUserInfo(ctx context.Context, state string, oauthState string, code string, conf *oauth2.Config) ([]byte, error) {
if state != oauthState {
	return nil, fmt.Errorf("invalid oauth state")
}

token, err :=  conf.Exchange(ctx, code)
if err != nil {
	return nil, fmt.Errorf("code exchange failed: %s", err.Error())
}

newurl := "https://api.vk.com/method/getProfiles?v=5.131&access_token=" + token.AccessToken
log.Print(newurl)
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




