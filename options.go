package geofox

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

func Debug(debug bool) Option {
	return func(a *API) error {
		a.Debug = debug
		return nil
	}
}
