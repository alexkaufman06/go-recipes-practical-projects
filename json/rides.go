// What is the maximal ride speed in rides.json?
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func maxRideSpeed(r io.Reader) (float64, error) {
	var rides []struct {
		StartTime string `json:"start"`
		EndTime   string `json:"end"`
		Id        string
		Distance  float64
	}

	dec := json.NewDecoder(r)
	if err := dec.Decode(&rides); err != nil {
		return 0, err
	}

	var maxRideSpeed float64
	const layout = "2006-01-02T15:04"
	for _, ride := range rides {
		startTime, err := time.Parse(layout, ride.StartTime)
		if err != nil {
			return 0, err
		}

		endTime, err := time.Parse(layout, ride.EndTime)
		if err != nil {
			return 0, err
		}

		totalTime := endTime.Sub(startTime)
		totalTimeHours := float64(totalTime) / float64(time.Hour)

		if ride.Distance/totalTimeHours > maxRideSpeed {
			maxRideSpeed = ride.Distance / totalTimeHours
		}
	}
	return maxRideSpeed, nil
}

func main() {
	file, err := os.Open("json/rides.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	speed, err := maxRideSpeed(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(speed) // 40.5
}
