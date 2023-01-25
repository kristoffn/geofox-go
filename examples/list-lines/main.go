package main

import (
	"context"
	"encoding/json"
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

	lines, err := client.ListLines(
		context.Background(),
		[]geofox.ModificationType{geofox.ModTypeMain},
		false,
	)
	if err != nil {
		panic(err)
	}
	linesBytes, err := json.Marshal(lines)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(linesBytes))
}
