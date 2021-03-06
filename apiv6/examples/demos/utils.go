package demos

import (
	"fmt"

	"github.com/kentik/community_sdk_golang/apiv6/examples"
)

var ExitOnError = examples.PanicOnError
var PrettyPrint = examples.PrettyPrint
var NewClient = examples.NewClient

func Step(msg string) {
	blueBold := "\033[1m\033[34m"
	reset := "\033[0m"
	examples.NewClient()
	fmt.Println()
	fmt.Printf("%s%s%s\n", blueBold, msg, reset)
	fmt.Printf("Press enter to continue...")
	fmt.Scanln()
}
