package geofox

type Option func(*API) error

func Debug() Option {
	return func(a *API) error {
		a.debug = true
		return nil
	}
}

func (a *API) parseOptions(opts ...Option) error {
	for _, option := range opts {
		err := option(a)
		if err != nil {
			return err
		}
	}
	return nil
}
