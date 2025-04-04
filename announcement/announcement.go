package announcement

import (
	"github.com/kristoffn/geofox-go/internal/config"
)

type AnnouncementService struct {
	Options []config.Option
}

func NewAnnouncementService(opts ...config.Option) (c *AnnouncementService) {
	c = &AnnouncementService{Options: opts}
	return
}
