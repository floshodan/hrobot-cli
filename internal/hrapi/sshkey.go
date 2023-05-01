package hrapi

import (
	"context"

	"github.com/floshodan/hrobot-go/hrobot"
)

type SSHKeyClientBase interface {
	//All(context.Context) ([]*hrobot.SSHKey, error)
	Create(context.Context, *hrobot.CreateKeyOpts) (*hrobot.SSHKey, *hrobot.Response, error)
	Delete(context.Context, string) (*hrobot.Response, error)
	GetByFingerprint(context.Context, string) (*hrobot.SSHKey, *hrobot.Response, error)
	//GetByName(context.Context, string) (*hrobot.SSHKey, *hrobot.Response, error)
	List(context.Context) ([]*hrobot.SSHKey, *hrobot.Response, error)
	Update(context.Context, string, *hrobot.UpdateKeyOpts) (*hrobot.SSHKey, *hrobot.Response, error)
}

type SSHKeyClient interface {
	SSHKeyClientBase
	Names() []string
}

func NewSSHKeyClient(client SSHKeyClientBase) SSHKeyClient {
	return &sshKeyClient{
		SSHKeyClientBase: client,
	}
}

type sshKeyClient struct {
	SSHKeyClientBase
}

// Names obtains a list of available SSH keys. It returns nil if SSH key
// names could not be fetched or none are available.
func (c *sshKeyClient) Names() []string {
	sshKeys, _, err := c.List(context.Background())
	if err != nil || len(sshKeys) == 0 {
		return nil
	}
	names := make([]string, len(sshKeys))
	for i, key := range sshKeys {
		name := key.Name
		if name == "" {
			name = key.Fingerprint
		}
		names[i] = name
	}
	return names
}
