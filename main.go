package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("ontime project")
	fmt.Println("a personal approach on implementing a public transport tracking and information system")
	fmt.Println("building this while either waiting for or using public transport")
	// Notes: LiveHopOnOff by App

}

// Location is a collection of coordinates combined with optional metadata
type Location struct {
	Lat, Lon int

	Name     string
	Address  string
	Locality string
	State    string
	Country  string
}

// Station is the physical location where passenger units may board or leave a Vehicle
type Station struct {
	ID   string
	Name string

	Type string

	Location *Location
}

// Stop is the timed stop of a Vehicle at a Station
type Stop struct {
	ID      string
	Station *Station

	Arrival   time.Time
	Departure time.Time
}

// Line is a chain of Stops connecting an origin and destination stop
type Line struct {
	ID   string
	Name string

	Origin      *Stop
	Stops       []*Stop
	Destination *Stop
}

// Vehicle is the physical transportation unit
type Vehicle struct {
	ID    string
	Plate string

	Seats int

	Location *Location
}

// Driver is the executing unit of a Tour controlling the Vehicle
type Driver struct {
	ID       string
	Fullname string
}

// Tour is a unique one time execution of a Line
type Tour struct {
	ID   string
	Line *Line

	Vehicle *Vehicle
	Driver  *Driver
	// TODO: Allow multiple Drivers (and Vehicles)

	// Target is the Stop currently targeted next
	Target *Stop
}

// User signed up for using the transportation service
type User struct {
	ID   string
	Name string
	Mail string

	Location *Location
}

// Passenger is a User who signed up to join a Tour at Stop
type Passenger struct {
	ID   string
	User *User

	Tour   *Tour
	HopOn  *Stop
	HopOff *Stop
}
