package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/kristoffn/geofox-go"
)

func main() {
	username := os.Getenv("GEOFOX_USER")
	password := os.Getenv("GEOFOX_PASSWD")
	client, err := geofox.New(username, password, geofox.AuthTypeHmacSHA1)
	if err != nil {
		panic(err)
	}
	responseBytes, err := client.Init()
	if err != nil {
		panic(err)
	}
	response := bytes.NewBuffer(responseBytes).String()
	fmt.Println(response)
}
