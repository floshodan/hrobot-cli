package state

import (
	"context"
	"os"

	"github.com/floshodan/hrobot-go/hrobot"
)

type State struct {
	Token         string
	Endpoint      string
	Context       context.Context
	ConfigPath    string
	Debug         bool
	DebugFilePath string

	client *hrobot.Client
}

func New() *State {
	s := &State{
		Context: context.Background(),
	}
	if p := os.Getenv("HROBOT_CONFIG"); p != "" {
		s.ConfigPath = p
	}
	return s
}

func (c *State) Client() *hrobot.Client {
	if c.client == nil {
		opts := []hrobot.ClientOption{
			hrobot.WithToken(c.Token),
		}
		if c.Debug {
			if c.DebugFilePath == "" {
				opts = append(opts, hrobot.WithDebugWriter(os.Stderr))
			} else {
				writer, _ := os.Create(c.DebugFilePath)
				opts = append(opts, hrobot.WithDebugWriter(writer))
			}
		}
		c.client = hrobot.NewClient(opts...)
	}
	return c.client
}
