package connector

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/answerdev/answer/internal/plugin"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
	oauth2GitHub "golang.org/x/oauth2/github"
)

type GitHub struct {
	ClientID     string
	ClientSecret string
}

func init() {
	plugin.Register(&GitHub{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	})
}

func (g *GitHub) Info() plugin.Info {
	return plugin.Info{
		Name:        "github connector",
		Description: "github connector plugin",
		Version:     "0.0.1",
	}
}

func (g *GitHub) ConnectorLogo() []byte {
	response, err := resty.New().R().Get("https://cdn-icons-png.flaticon.com/32/25/25231.png")
	if err != nil {
		return nil
	}
	return response.Body()
}

func (g *GitHub) ConnectorLogoContentType() string {
	return "image/png"
}

func (g *GitHub) ConnectorName() string {
	return "GitHub"
}

func (g *GitHub) ConnectorSlugName() string {
	return "github"
}

func (g *GitHub) ConnectorSender(ctx *plugin.GinContext) {
	oauth2Config := &oauth2.Config{
		ClientID:     g.ClientID,
		ClientSecret: g.ClientSecret,
		Endpoint:     oauth2GitHub.Endpoint,
		RedirectURL:  "http://127.0.0.1:8080/answer/api/v1/oauth/redirect/github", // TODO: Pass by parameter
		Scopes:       nil,
	}
	ctx.Redirect(http.StatusFound, oauth2Config.AuthCodeURL("state"))
}

func (g *GitHub) ConnectorReceiver(ctx *plugin.GinContext) {
	code := ctx.Query("code")
	//state := ctx.Query("state")

	// Exchange code for token
	oauth2Config := &oauth2.Config{
		ClientID:     g.ClientID,
		ClientSecret: g.ClientSecret,
		Endpoint:     oauth2GitHub.Endpoint,
	}
	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/50x")
		return
	}

	// Exchange token for user info
	cli := github.NewClient(oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token.AccessToken},
	)))
	userInfo, _, err := cli.Users.Get(context.Background(), "")
	if err != nil {
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

func (g *GitHub) ConnectorLoginURL(redirectURL, state string) (loginURL string) {
	oauth2Config := &oauth2.Config{
		ClientID:     g.ClientID,
		ClientSecret: g.ClientSecret,
		Endpoint:     oauth2GitHub.Endpoint,
		RedirectURL:  redirectURL,
	}
	return oauth2Config.AuthCodeURL(state)
}

func (g *GitHub) ConnectorLoginUserInfo(code string) (userInfo *plugin.UserExternalLogin, err error) {
	// Exchange code for token
	oauth2Config := &oauth2.Config{
		ClientID:     g.ClientID,
		ClientSecret: g.ClientSecret,
		Endpoint:     oauth2GitHub.Endpoint,
	}
	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	// Exchange token for user info
	cli := github.NewClient(oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token.AccessToken},
	)))
	resp, _, err := cli.Users.Get(context.Background(), "")
	if err != nil {
		return nil, err
	}

	userInfo = &plugin.UserExternalLogin{
		Provider:    g.ConnectorSlugName(),
		ExternalID:  fmt.Sprintf("%d", resp.GetID()),
		Email:       resp.GetEmail(),
		Name:        resp.GetName(),
		FirstName:   resp.GetName(),
		LastName:    resp.GetName(),
		NickName:    resp.GetName(),
		Description: resp.GetBio(),
		AvatarUrl:   resp.GetAvatarURL(),
	}
	return userInfo, nil
}
