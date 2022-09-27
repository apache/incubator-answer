package web

import "embed"

//go:embed html/index.html
var Html []byte

//go:embed html/static
var Static embed.FS
