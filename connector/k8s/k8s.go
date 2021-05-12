package k8s

import (
	"context"
	"fmt"

	"github.com/dexidp/dex/connector"
	"github.com/dexidp/dex/pkg/log"
)

// Config holds configuration options for K8s logins.
type Config struct {
	Namespace string `json:"namespace"`
}

// Open returns an authentication strategy using K8s.
func (c *Config) Open(id string, logger log.Logger) (connector.Connector, error) {
	conn, err := c.OpenConnector(logger)
	if err != nil {
		return nil, err
	}
	return connector.Connector(conn), nil
}

// OpenConnector is the same as Open but returns a type with all implemented connector interfaces.
func (c *Config) OpenConnector(logger log.Logger) (interface {
	connector.Connector
	connector.PasswordConnector
	connector.RefreshConnector
}, error) {
	return c.openConnector(logger)
}

func (c *Config) openConnector(logger log.Logger) (*KubernetesConnector, error) {
	requiredFields := []struct {
		name string
		val  string
	}{
		{"namespace", c.Namespace},
	}

	for _, field := range requiredFields {
		if field.val == "" {
			return nil, fmt.Errorf("kubernetes: missing required field %q", field.name)
		}
	}

	return &KubernetesConnector{*c, c.Namespace, logger}, nil
}

type KubernetesConnector struct {
	Config
	Namespace string
	Logger    log.Logger
}

func (k *KubernetesConnector) Prompt() string {
	return "TODO"
}

func (k *KubernetesConnector) Login(ctx context.Context, s Scopes, username, password string) (identity Identity, validPassword bool, err error) {
	return nil, false, nil
}
