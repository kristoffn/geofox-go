package model

type Filter string

const (
	NoFilter  Filter = "NO_FILTER"
	HVVListed Filter = "HVV_LISTED"
)

type BaseRequest struct {
	Language   string `json:"language,omitempty"`
	Version    int32  `json:"version,omitempty"`
	FilterType Filter `json:"filterType,omitempty"`
}
