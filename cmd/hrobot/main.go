package main

import (
	"log"
	"os"

	"github.com/floshodan/hrobot-cli/internal/cli"
	"github.com/floshodan/hrobot-cli/internal/hrapi"
	"github.com/floshodan/hrobot-cli/internal/state"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("hrobot: ")
	log.SetOutput(os.Stderr)
}

func main() {

	cliState := state.New()
	cliState.Token = os.Getenv("HROBOT_TOKEN")

	apiClient := hrapi.NewClient(cliState.Client())
	//apiClient := hrapi.NewClient(hrobot.WithToken(os.Getenv("Hrobot_TOKEN")))
	rootCommand := cli.NewRootCommand(cliState, apiClient)

	//app := cli.NewApp(hrobot.NewClient(hrobot.WithToken(os.Getenv("HROBOT_TOKEN"))))

	if err := rootCommand.Execute(); err != nil {
		log.Fatalln(err)
	}

}
