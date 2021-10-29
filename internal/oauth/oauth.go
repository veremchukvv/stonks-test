package oauth

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//var VKconfig *oauth2.Config

func GetOauthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("VK_CLIENT_ID"),
		ClientSecret: os.Getenv("VK_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("VK_REDIRECT_URL"),
		Scopes:       []string{""},
		Endpoint:     vk.Endpoint,
}
}

func GetRandomState() string {
	//TODO randomize state
	return "blabla"
}

func GetUserInfo(ctx context.Context, state string, oauthState string, code string, conf *oauth2.Config) ([]byte, error) {
	log.Print(state)
	log.Print(oauthState)
if state != oauthState {
	return nil, fmt.Errorf("invalid oauth state")
}

token, err :=  conf.Exchange(ctx, code)
if err != nil {
	return nil, fmt.Errorf("code exchange failed: %s", err.Error())
}

response, err := http.Get("https://api.vk.com/method/GetProfileInfo&access_token=" + token.AccessToken)
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




