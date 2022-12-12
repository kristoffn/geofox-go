package main

import (
	"context"
	"fmt"

	"github.com/artisanalbushfire/geofox-go"
)

func main() {
	api, err := geofox.New("	", "")
	if err != nil {
		panic(err)
	}
	resp, err := api.Init(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
