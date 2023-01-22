# geofox-go
[![Go Report Card](https://goreportcard.com/badge/github.com/kristoffn/geofox-go)](https://goreportcard.com/report/github.com/kristoffn/geofox-go)
## About

geofox-go is a Goclient for the Geofox Thin Interface (GTI) to provide information about Hamburg's public transport system.

## Install
```
go get -u github.com/kristoffn/geofox-go
```

## Example usage
This small application is an example client to send an `init` request to the Geofox API. 
```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kristoffn/geofox-go"
)

func main() {
	username := os.Getenv("GEOFOX_USER")
	password := os.Getenv("GEOFOX_PASSWD")
	client, err := geofox.New(username, password)
	if err != nil {
		panic(err)
	}

	initResponse, err := client.Init(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", initResponse)
}

```

## Note
:exclamation: This library is still in a very experimental state and subject to potential breaking changes. Use at your own discretion.