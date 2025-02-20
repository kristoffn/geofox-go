package geofox

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	optionReturnError Option = func(a *API) error {
		return errors.New("test error")
	}
)

func TestParseOptions(t *testing.T) {
	api := API{}
	err := api.parseOptions()
	assert.Nil(t, err)

	err = api.parseOptions(optionReturnError)
	assert.ErrorContains(t, err, "test error")
}

func TestOptionDebug(t *testing.T) {
	api := API{}
	assert.False(t, api.debug)
	err := api.parseOptions(Debug())
	assert.Nil(t, err)
	assert.True(t, api.debug)
}
