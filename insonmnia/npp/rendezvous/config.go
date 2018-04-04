package rendezvous

import (
	"crypto/ecdsa"
	"net"

	"github.com/jinzhu/configor"
	"github.com/sonm-io/core/accounts"
	"github.com/sonm-io/core/insonmnia/auth"
	"github.com/sonm-io/core/insonmnia/logging"
	"github.com/sonm-io/core/util/netutil"
	"go.uber.org/zap/zapcore"
)

// LoggingConfig represents a logging config.
type LoggingConfig struct {
	Level string `required:"true" default:"debug"`
	level zapcore.Level
}

// ServerConfig represents a Rendezvous server configuration.
type ServerConfig struct {
	// Listening address.
	Addr       net.Addr
	PrivateKey *ecdsa.PrivateKey
	Logging    LoggingConfig
}

// LogLevel returns the minimum logging level configured.
func (c *ServerConfig) LogLevel() zapcore.Level {
	return c.Logging.level
}

type serverConfig struct {
	Addr    netutil.TCPAddr    `yaml:"endpoint" required:"true"`
	Eth     accounts.EthConfig `yaml:"ethereum"`
	Logging LoggingConfig      `yaml:"logging"`
}

// NewServerConfig loads a new Rendezvous server config from a file.
func NewServerConfig(path string) (*ServerConfig, error) {
	cfg := &serverConfig{}
	err := configor.Load(cfg, path)
	if err != nil {
		return nil, err
	}

	privateKey, err := cfg.Eth.LoadKey()
	if err != nil {
		return nil, err
	}

	lvl, err := logging.ParseLogLevel(cfg.Logging.Level)
	if err != nil {
		return nil, err
	}
	cfg.Logging.level = lvl

	return &ServerConfig{
		Addr:       &cfg.Addr,
		PrivateKey: privateKey,
		Logging:    cfg.Logging,
	}, nil
}

type Config struct {
	Endpoints []auth.Addr
}