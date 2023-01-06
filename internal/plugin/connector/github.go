package connector

import (
	"context"
	"encoding/json"
	"fmt"
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

func (g *GitHub) ConnectorSender(ctx *plugin.GinContext, receiverURL string) (redirectURL string) {
	oauth2Config := &oauth2.Config{
		ClientID:     g.ClientID,
		ClientSecret: g.ClientSecret,
		Endpoint:     oauth2GitHub.Endpoint,
		RedirectURL:  receiverURL,
		Scopes:       nil,
	}
	return oauth2Config.AuthCodeURL("state")
}

func (g *GitHub) ConnectorReceiver(ctx *plugin.GinContext) (userInfo plugin.ExternalLoginUserInfo, err error) {
	code := ctx.Query("code")
	// Exchange code for token
	oauth2Config := &oauth2.Config{
		ClientID:     g.ClientID,
		ClientSecret: g.ClientSecret,
		Endpoint:     oauth2GitHub.Endpoint,
	}
	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		return userInfo, err
	}

	// Exchange token for user info
	cli := github.NewClient(oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token.AccessToken},
	)))
	resp, _, err := cli.Users.Get(context.Background(), "")
	if err != nil {
		return userInfo, err
	}

	metaInfo, _ := json.Marshal(resp)
	userInfo = plugin.ExternalLoginUserInfo{
		ExternalID: fmt.Sprintf("%d", resp.GetID()),
		Name:       resp.GetName(),
		Email:      resp.GetEmail(),
		MetaInfo:   string(metaInfo),
	}
	return userInfo, nil
}
