package hrapi

import (
	"context"

	"github.com/floshodan/hrobot-go/hrobot"
)

type ServerClientBase interface {
	List(context.Context) ([]*hrobot.Server, *hrobot.Response, error)
}

type ServerClient interface {
	ServerClientBase
}

func NewServerClient(client ServerClientBase) ServerClient {
	return &serverClient{
		ServerClientBase: client,
	}
}

type serverClient struct {
	ServerClientBase
}
