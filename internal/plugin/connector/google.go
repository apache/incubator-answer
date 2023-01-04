package connector

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/answerdev/answer/internal/plugin"
	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2"
	oauth2Google "golang.org/x/oauth2/google"
)

type Google struct {
	ClientID     string
	ClientSecret string
}

func init() {
	plugin.Register(&Google{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	})
}

func (g *Google) Info() plugin.Info {
	return plugin.Info{
		Name:        "google connector",
		Description: "google connector plugin",
		Version:     "0.0.1",
	}
}

func (g *Google) ConnectorLogo() []byte {
	response, err := resty.New().R().Get("https://cdn-icons-png.flaticon.com/32/300/300221.png")
	if err != nil {
		return nil
	}
	return response.Body()
}

func (g *Google) ConnectorLogoContentType() string {
	return "image/png"
}

func (g *Google) ConnectorName() string {
	return "Google"
}

func (g *Google) ConnectorSlugName() string {
	return "google"
}

func (g *Google) ConnectorSender(ctx *plugin.GinContext) {
	oauth2Config := &oauth2.Config{
		ClientID:     g.ClientID,
		ClientSecret: g.ClientSecret,
		Endpoint:     oauth2Google.Endpoint,
		RedirectURL:  "http://127.0.0.1:8080/answer/api/v1/oauth/redirect/google", // TODO: Pass by parameter
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"openid",
		},
	}
	ctx.Redirect(http.StatusFound, oauth2Config.AuthCodeURL("state"))
}

type GoogleAuthUserInfo struct {
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

func (g *Google) ConnectorReceiver(ctx *plugin.GinContext) {
	code := ctx.Query("code")
	oauth2Config := &oauth2.Config{
		ClientID:     g.ClientID,
		ClientSecret: g.ClientSecret,
		Endpoint:     oauth2Google.Endpoint,
	}

	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/50x")
		return
	}

	client := oauth2Config.Client(context.TODO(), token)
	client.Timeout = 60 * time.Second
	userinfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		ctx.Redirect(http.StatusFound, "/50x")
		return
	}
	defer userinfo.Body.Close()
	data, _ := io.ReadAll(userinfo.Body)

	userInfo := &GoogleAuthUserInfo{}
	if err = json.Unmarshal(data, userInfo); err != nil {
		ctx.Redirect(http.StatusFound, "/50x")
		return
	}

	fmt.Printf("user info is :%+v", userInfo)

	// TODO
	// If user email exists, try to login this user.
	// If user email not exists, try to register this user.

	ctx.Redirect(http.StatusFound, "/login-success?access_token=token")
	return
}
