package connector

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/answerdev/answer/plugin"
	"golang.org/x/oauth2"
	oauth2Google "golang.org/x/oauth2/google"
)

type Google struct {
	Config *GoogleConfig
}

type GoogleConfig struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type GoogleAuthUserInfo struct {
	ID            string `json:"id"`
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}

func init() {
	plugin.Register(&Google{
		Config: &GoogleConfig{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		},
	})
}

func (g *Google) Info() plugin.Info {
	return plugin.Info{
		Name:        "google connector",
		SlugName:    "google_connector",
		Description: "google connector plugin",
		Version:     "0.0.1",
	}
}

func (g *Google) ConnectorLogoSVG() string {
	//TODO implement me
	panic("implement me")
}

func (g *Google) ConnectorSender(ctx *plugin.GinContext, receiverURL string) (redirectURL string) {
	oauth2Config := &oauth2.Config{
		ClientID:     g.Config.ClientID,
		ClientSecret: g.Config.ClientSecret,
		Endpoint:     oauth2Google.Endpoint,
		RedirectURL:  receiverURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"openid",
		},
	}
	return oauth2Config.AuthCodeURL("state")
}

func (g *Google) ConnectorReceiver(ctx *plugin.GinContext) (userInfo plugin.ExternalLoginUserInfo, err error) {
	code := ctx.Query("code")
	oauth2Config := &oauth2.Config{
		ClientID:     g.Config.ClientID,
		ClientSecret: g.Config.ClientSecret,
		Endpoint:     oauth2Google.Endpoint,
	}

	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		return userInfo, err
	}

	client := oauth2Config.Client(context.TODO(), token)
	client.Timeout = 15 * time.Second
	response, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return userInfo, err
	}
	defer response.Body.Close()
	data, _ := io.ReadAll(response.Body)

	respGoogleAuthUserInfo := &GoogleAuthUserInfo{}
	if err = json.Unmarshal(data, respGoogleAuthUserInfo); err != nil {
		return userInfo, fmt.Errorf("parse google oauth user info response failed: %v", err)
	}

	userInfo = plugin.ExternalLoginUserInfo{
		ExternalID: respGoogleAuthUserInfo.ID,
		Name:       respGoogleAuthUserInfo.Name,
		Email:      respGoogleAuthUserInfo.Email,
		MetaInfo:   string(data),
	}
	return userInfo, nil
}

func (g *Google) ConnectorName() string {
	return "Google"
}

func (g *Google) ConnectorSlugName() string {
	return "google"
}
