module github.com/kristoffn/geofox-go

go 1.24.0

require (
	github.com/stretchr/testify v1.10.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	[v0.0.1-alpha, v0.0.1, v0.0.2-alpha, v0.0.2, v0.0.3-alpha, v0.0.3, v0.0.4-alpha, v0.0.4]
)
