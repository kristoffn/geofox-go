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

	initResponse, err := client.Init()
	if err != nil {
		panic(err)
	}

	stations, err := client.ListStations(
		initResponse.DataID,
		[]geofox.ModificationType{geofox.ModificationTypePosition},
		geofox.CoordinateTypeEPSG4326)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Stations: %v\n", stations)
}
