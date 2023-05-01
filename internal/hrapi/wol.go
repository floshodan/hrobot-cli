package hrapi

import (
	"context"

	"github.com/floshodan/hrobot-go/hrobot"
)

type WakeOnLanClientBase interface {
	GetByServernumber(context.Context, string) (*hrobot.WOL, *hrobot.Response, error)
	Send(context.Context, string) (*hrobot.WOL, *hrobot.Response, error)
}

type WakeOnLanClient interface {
	WakeOnLanClientBase
}

func NewWakeOnLanClient(client WakeOnLanClientBase) WakeOnLanClient {
	return &wolClient{
		WakeOnLanClientBase: client,
	}
}

type wolClient struct {
	WakeOnLanClientBase
}
