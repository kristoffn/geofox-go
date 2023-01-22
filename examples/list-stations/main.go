package main

import (
	"context"
	"os"

	"github.com/kristoffn/geofox-go"
	"github.com/kristoffn/geofox-go/coordinates"
	"github.com/kristoffn/geofox-go/model"
)

func main() {
	username := os.Getenv("GEOFOX_USER")
	password := os.Getenv("GEOFOX_PASSWD")
	client, err := geofox.New(username, password, geofox.Debug())
	if err != nil {
		panic(err)
	}

	_, err = client.ListStations(
		context.Background(),
		[]model.ModificationType{model.ModificationTypeMain},
		coordinates.EPSG31467)
	if err != nil {
		panic(err)
	}
}
