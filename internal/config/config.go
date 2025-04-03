package config

type Config struct {
	Debug    bool
	Password string
	Username string
	AuthType string
}

func (cfg *Config) Apply(opts ...Option) error {
	for _, opt := range opts {
		err := opt.Apply(cfg)
		if err != nil {
			return err
		}
	}
	return nil
}

type Option interface {
	Apply(*Config) error
}

type OptionFunc func(*Config) error

func (of OptionFunc) Apply(c *Config) error { return of(c) }

func NewConfig(opts ...Option) (*Config, error) {
	cfg := &Config{}
	err := cfg.Apply(opts...)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
