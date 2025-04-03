package option

import (
	"github.com/kristoffn/geofox-go/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDebugOption(t *testing.T) {
	conf, err := config.NewConfig()
	assert.Nil(t, err)
	assert.Equal(t, false, conf.Debug)

	conf, err = config.NewConfig(
		Debug(),
	)
	assert.Nil(t, err)
	assert.Equal(t, true, conf.Debug)
}

func TestWithPassword(t *testing.T) {
	conf, err := config.NewConfig()
	assert.Nil(t, err)
	assert.Equal(t, "", conf.Password)

	conf, err = config.NewConfig(
		WithPassword("test-password"),
	)
	assert.Nil(t, err)
	assert.Equal(t, "test-password", conf.Password)
}

func TestWithUsername(t *testing.T) {
	conf, err := config.NewConfig()
	assert.Nil(t, err)
	assert.Equal(t, "", conf.Username)

	conf, err = config.NewConfig(
		WithUsername("test-username"),
	)
	assert.Nil(t, err)
	assert.Equal(t, "test-username", conf.Username)
}
