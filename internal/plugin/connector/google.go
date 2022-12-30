package connector

import (
	"net/http"

	"github.com/answerdev/answer/internal/plugin"
)

type Google struct {
}

func (g *Google) Info() plugin.Info {
	return plugin.Info{
		Name:        "google connector",
		Description: "google connector plugin",
		Version:     "0.0.1",
	}
}

func (g *Google) ConnectorLogo() []byte {
	return nil
}

func (g *Google) ConnectorLogoContentType() string {
	return "image/png"
}

func (g *Google) ConnectorName() string {
	return "google"
}

func (g *Google) ConnectorSlugName() string {
	return "google"
}

func (g *Google) ConnectorSender(ctx *plugin.GinContext) {
	//TODO implement me
	panic("implement me")
}

func (g *Google) ConnectorReceiver(ctx *plugin.GinContext) {
	ctx.String(http.StatusOK, "OK123")
}

func init() {
	plugin.Register(&Google{})
}
