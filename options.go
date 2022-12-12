package geofox

// Option is a functional option for configuring the API client
type Option func(*API) error

func UsingLogger(logger Logger) Option {
	return func(api *API) error {
		api.logger = logger
		return nil
	}
}

func (api *API) parseOptions(opts ...Option) error {
	for _, option := range opts {
		err := option(api)
		if err != nil {
			return err
		}
	}
	return nil
}
