package main

import (
	"context"
	"log"
	"os"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
)

const (
	URL         = "https://eu-central-1-1.aws.cloud2.influxdata.com"
	Database    = "broadband"
	Measurement = "broadband"
	Location    = "home"
	Provider    = "PlusNet"
)

func WriteToInflux(upRate int, downRate int) {

	// Create a new client using an InfluxDB server base URL and an authentication token
	client, err := influxdb3.New(influxdb3.ClientConfig{
		Host:     URL,
		Token:    os.Getenv("INFLUXDB_TOKEN"),
		Database: Database,
	})
	if err != nil {
		panic(err)
	}

	// Close client at the end and escalate error if present
	defer func(client *influxdb3.Client) {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}(client)

	// Create InfluxDB Point and Add to Slice
	point := influxdb3.NewPointWithMeasurement(Measurement).
		SetTag("location", Location).
		SetTag("provider", Provider).
		SetField("upload_speed", upRate).
		SetField("download_speed", downRate)
	points := []*influxdb3.Point{point}

	// Write Point
	if err := client.WritePoints(context.Background(), points); err != nil {
		log.Println(err)
	}
}
