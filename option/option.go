package option

import (
	"github.com/kristoffn/geofox-go/internal/config"
)

// Debug returns an Option that sets the debug flag to true
func Debug() config.Option {
	return config.OptionFunc(func(c *config.Config) error {
		c.Debug = true
		return nil
	})
}
