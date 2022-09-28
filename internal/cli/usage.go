package cli

import "fmt"

var usageDoc = `
answer

USAGE
    answer command

COMMANDS
    init        init answer config
    -c          config path, eg: -c data/config.yaml
`

func Usage() {
	fmt.Println(usageDoc)
}
