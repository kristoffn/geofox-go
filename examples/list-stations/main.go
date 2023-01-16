package main

import (
	"os"

	"github.com/kristoffn/geofox-go"
)

func main() {
	username := os.Getenv("GEOFOX_USER")
	password := os.Getenv("GEOFOX_PASSWD")
	client, err := geofox.New(username, password, geofox.Debug())
	if err != nil {
		panic(err)
	}

	_, err = client.ListStations(
		[]geofox.ModificationType{geofox.ModificationTypeMain},
		geofox.CoordinateTypeEPSG31467)
	if err != nil {
		panic(err)
	}
}
