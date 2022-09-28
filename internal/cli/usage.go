package cli

import "fmt"

var usageDoc = `
answer

USAGE
    answer command

COMMANDS
    init         init answer config, eg:answer init
    run          config path, eg:answer run -c data/config.yaml
`

func Usage() {
	fmt.Println(usageDoc)
}
