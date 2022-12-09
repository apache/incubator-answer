package configs

import _ "embed"

//go:embed  config.yaml
var Config []byte

//go:embed  path_ignore.yaml
var PathIgnore []byte
