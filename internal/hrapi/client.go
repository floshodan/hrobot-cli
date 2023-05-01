package hrapi

import (
	"github.com/floshodan/hrobot-go/hrobot"
)

type Client interface {
	SSHKey() SSHKeyClient
	WOL() WakeOnLanClient
	Server() ServerClient
}

type client struct {
	client          *hrobot.Client
	sshKeyClient    SSHKeyClient
	wakeOnLanClient WakeOnLanClient
	serverClient    ServerClient
}

func NewClient(c *hrobot.Client) Client {

	return &client{
		client: c,
	}
}

func (c *client) SSHKey() SSHKeyClient {
	//c.mu.Lock()
	if c.sshKeyClient == nil {
		c.sshKeyClient = NewSSHKeyClient(&c.client.SSHKey)
	}
	//defer c.mu.Unlock()
	return c.sshKeyClient
}

func (c *client) WOL() WakeOnLanClient {
	if c.wakeOnLanClient == nil {
		c.wakeOnLanClient = NewWakeOnLanClient(&c.client.WakeOnLane)
	}
	return c.wakeOnLanClient
}

func (c *client) Server() ServerClient {
	if c.serverClient == nil {
		c.serverClient = NewServerClient(&c.client.Server)
	}
	return c.serverClient
}
