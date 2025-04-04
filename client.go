package geofox

import (
	"github.com/kristoffn/geofox-go/announcement"
	"github.com/kristoffn/geofox-go/internal/config"
)

type Client struct {
	Options      []config.Option
	Announcement *announcement.AnnouncementService
}

func NewClient(opts ...config.Option) (c *Client) {
	c = &Client{Options: opts}

	c.Announcement = announcement.NewAnnouncementService(opts...)
	return
}
