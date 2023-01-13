package converter

import (
	"bytes"

	"github.com/microcosm-cc/bluemonday"
	"github.com/segmentfault/pacman/log"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	goldmarkHTML "github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

// Markdown2HTML convert markdown to html
func Markdown2HTML(source string) string {
	mdConverter := goldmark.New(
		goldmark.WithExtensions(&DangerousHTMLFilterExtension{}, extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			goldmarkHTML.WithHardWraps(),
		),
	)
	var buf bytes.Buffer
	if err := mdConverter.Convert([]byte(source), &buf); err != nil {
		log.Error(err)
		return source
	}
	return buf.String()
}

type DangerousHTMLFilterExtension struct {
}

func (e *DangerousHTMLFilterExtension) Extend(m goldmark.Markdown) {
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(&DangerousHTMLRenderer{
			Config: goldmarkHTML.NewConfig(),
			Filter: bluemonday.UGCPolicy(),
		}, 1),
	))
}

type DangerousHTMLRenderer struct {
	goldmarkHTML.Config
	Filter *bluemonday.Policy
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *DangerousHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindHTMLBlock, r.renderHTMLBlock)
	reg.Register(ast.KindRawHTML, r.renderRawHTML)
}

func (r *DangerousHTMLRenderer) renderRawHTML(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkSkipChildren, nil
	}
	n := node.(*ast.RawHTML)
	l := n.Segments.Len()
	for i := 0; i < l; i++ {
		segment := n.Segments.At(i)
		_, _ = w.Write(r.Filter.SanitizeBytes(segment.Value(source)))
	}
	return ast.WalkSkipChildren, nil
}

func (r *DangerousHTMLRenderer) renderHTMLBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.HTMLBlock)
	if entering {
		l := n.Lines().Len()
		for i := 0; i < l; i++ {
			line := n.Lines().At(i)
			r.Writer.SecureWrite(w, r.Filter.SanitizeBytes(line.Value(source)))
		}
	} else {
		if n.HasClosure() {
			closure := n.ClosureLine
			r.Writer.SecureWrite(w, closure.Value(source))
		}
	}
	return ast.WalkContinue, nil
}
