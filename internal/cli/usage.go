package cli

import "fmt"

var usageDoc = `
Welcome to answer

VERSION:
   1.0.0

USAGE:
   answer  [global options] command [command options] [arguments...]

COMMANDS:
   init         Init config, eg:./answer init
   run          Start web server, eg:./answer run -c data/config.yaml
`

func Usage() {
	fmt.Println(usageDoc)
}
