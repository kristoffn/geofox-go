package main

import (
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

	stations, err := client.ListStations(
		[]geofox.ModificationType{geofox.ModificationTypeMain},
		geofox.CoordinateTypeEPSG4326)
	if err != nil {
		panic(err)
	}
	for k, v := range stations {
		fmt.Printf("%s -> %s\n", k, v.Name)
	}
}
