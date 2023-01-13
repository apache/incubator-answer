package connector

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/answerdev/answer/internal/plugin"
	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
	oauth2GitHub "golang.org/x/oauth2/github"
)

type GitHub struct {
	Config *GitHubConfig
}

type GitHubConfig struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func init() {
	plugin.Register(&GitHub{
		Config: &GitHubConfig{
			ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
			ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		},
	})
}

func (g *GitHub) Info() plugin.Info {
	return plugin.Info{
		Name:        "github connector",
		SlugName:    "github_connector",
		Description: "github connector plugin",
		Version:     "0.0.1",
	}
}

func (g *GitHub) ConnectorLogoSVG() string {
	return `M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z`
}

func (g *GitHub) ConnectorName() string {
	return "GitHub"
}

func (g *GitHub) ConnectorSlugName() string {
	return "github"
}

func (g *GitHub) ConnectorSender(ctx *plugin.GinContext, receiverURL string) (redirectURL string) {
	oauth2Config := &oauth2.Config{
		ClientID:     g.Config.ClientID,
		ClientSecret: g.Config.ClientSecret,
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
		ClientID:     g.Config.ClientID,
		ClientSecret: g.Config.ClientSecret,
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

func (g *GitHub) ConfigFields() []plugin.ConfigField {
	return []plugin.ConfigField{
		{
			Name:        "ClientID",
			Description: "Client ID of your GitHub application.",
			Required:    true,
			Type:        plugin.ConfigTypeInput,
			Items: []plugin.ConfigFieldItem{
				{
					Name:  "client_id",
					Label: "ClientID",
					Value: g.Config.ClientID,
				},
			},
		},
		{
			Name:        "ClientSecret",
			Description: "Client secret of your GitHub application.",
			Required:    true,
			Type:        plugin.ConfigTypeInput,
			Items: []plugin.ConfigFieldItem{
				{
					Name:  "client_secret",
					Label: "ClientSecret",
					Value: g.Config.ClientSecret,
				},
			},
		},
	}
}

func (g *GitHub) ConfigReceiver(config []byte) error {
	c := &GitHubConfig{}
	_ = json.Unmarshal(config, c)
	g.Config = c
	return nil
}
