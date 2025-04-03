package geofox

import "go.uber.org/zap"

type Option func(*API) error

func (a *API) parseOptions(opts ...Option) error {
	for _, option := range opts {
		err := option(a)
		if err != nil {
			return err
		}
	}
	return nil
}

func Debug() Option {
	return func(a *API) error {
		a.debug = true
		return nil
	}
}

func WithLogger(logger *zap.Logger) Option {
	return func(a *API) error {
		if logger != nil {
			a.logger = logger
		}
		return nil
	}
}
