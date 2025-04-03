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
