package main

import (
	"fmt"
)

// A Thermostat measures and controls the temperature
type Thermostats struct {
	id    string
	value float64
}

// ID return the thermostat ID
func (t *Thermostats) ID() string {
	return t.id
}

// Value return the current temperature in Celsius
func (t *Thermostats) Value() float64 {
	return t.value
}

// Kind returns the device kind
func (*Thermostats) Kind() string {
	return "thermostats"
}

// Camera is a security camera
type Camera struct {
	id string
}

// ID return the camera ID
func (c *Camera) ID() string {
	return c.id
}

func (*Camera) Kind() string {
	return "camera"
}

type Sensor interface {
	ID() string
	Kind() string
}

func printAll(sensors []Sensor) {
	for _, s := range sensors {
		fmt.Printf("%s <%s>\n", s.ID(), s.Kind())
	}
}

func main() {
	t := Thermostats{"Living Room", 16.2}
	c := Camera{"Baby room"}

	sensors := []Sensor{&t, &c}
	printAll(sensors)
	/*
		Living Room <thermostat>
		Baby room <camera>
	*/
}
