/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cli

import (
	"fmt"
	"os"

	"github.com/floshodan/hrobot-cli/internal/cmd/server"
	"github.com/floshodan/hrobot-cli/internal/cmd/sshkey"
	"github.com/floshodan/hrobot-cli/internal/cmd/wol"
	"github.com/floshodan/hrobot-cli/internal/hrapi"
	"github.com/floshodan/hrobot-cli/internal/state"
	"github.com/floshodan/hrobot-go/hrobot"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

const version = "0.0.1"

type App struct {
	client *hrobot.Client
	output string
}

func NewApp(c *hrobot.Client) *App {

	return &App{
		client: c,
	}
}

func NewRootCommand(state *state.State, client hrapi.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "hrobot",
		Short:                 "Hetzner Robot CLI",
		Long:                  "A command-line interface for Hetzner Robot Interface",
		TraverseChildren:      true,
		SilenceUsage:          true,
		SilenceErrors:         true,
		DisableFlagsInUseLine: true,
	}

	cmd.AddCommand(server.NewCommand(state, client))
	cmd.AddCommand(sshkey.NewCommand(state, client))
	cmd.AddCommand(wol.NewCommand(state, client))
	//cmd.AddCommand(app.NewVersionCmd())
	return cmd
}

func (app *App) NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of hrobot-cli",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(fmt.Sprintf("Hetzner Robot cli version %s", version))
		},
	}
}

// rootCmd represents the base command when called without any subcommands

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.

func init() {
	cobra.OnInitialize(initConfig)

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".hrobot-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".hrobot")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
