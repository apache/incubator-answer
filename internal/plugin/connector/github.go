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
	return "GitHubConnector"
}

func (g *GitHub) ConnectorSlugName() string {
	return "github"
}

func (g *GitHub) ConnectorSender(ctx *plugin.GinContext) {
	oauth2Config := &oauth2.Config{
		ClientID:     g.ClientID,
		ClientSecret: g.ClientSecret,
		Endpoint:     oauth2GitHub.Endpoint,
		RedirectURL:  "http://127.0.0.1:8080/answer/api/v1/oauth/redirect/github",
		Scopes:       nil,
	}
	ctx.Redirect(http.StatusFound, oauth2Config.AuthCodeURL(""))
}

func (g *GitHub) ConnectorReceiver(ctx *plugin.GinContext) {
	code := ctx.Query("code")
	//state := ctx.Query("state")

	oauth2Config := &oauth2.Config{
		ClientID:     g.ClientID,
		ClientSecret: g.ClientSecret,
		Endpoint:     oauth2GitHub.Endpoint,
	}
	token, err := oauth2Config.Exchange(context.Background(), code+"1")
	if err != nil {
		ctx.Error(err)
		ctx.Redirect(http.StatusFound, "/50x")
		return
	}

	cli := github.NewClient(oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token.AccessToken},
	)))
	userInfo, _, err := cli.Users.Get(context.Background(), "")
	if err != nil {
		ctx.Error(err)
		ctx.Redirect(http.StatusFound, "/50x")
		return
	}

	fmt.Printf("user info is :%+v", userInfo)
	ctx.Redirect(http.StatusFound, "/")
	return
}
