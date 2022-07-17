//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Lifter interface {
	LiftVehicle()
}

type Vehicle struct {
	model string
	vtype int
}

const (
	Motorcycles = iota
	Cars
	Truck
)

func (vehicle Vehicle) LiftVehicle() {
	fmt.Printf("Ready to lift : %v  \n", vehicle.model)

	switch vehicle.vtype {
	case Motorcycles:
		fmt.Println("USE SMALL LIFT")
	case Cars:
		fmt.Println("USE STANDARAD LIFT")
	case Truck:
		fmt.Println("USE LARGE LIFT")
	default:
		fmt.Println("DO NOT LIFT, SIZE NOT AVAILABLE")
	}
}

func prepareLift(vehicles []Lifter) {
	fmt.Println("Vechiles Queued to be lifted!")
	for i := 0; i < len(vehicles); i++ {
		vehicle := vehicles[i]
		fmt.Printf("--- Ready to lift---\n")
		vehicle.LiftVehicle()
	}
}

func main() {
	vehicles := []Lifter{
		Vehicle{model: "Land Rover", vtype: Cars},
		Vehicle{model: "Pulsar", vtype: Motorcycles},
		Vehicle{model: "Ashok Leyland", vtype: Truck},
		Vehicle{model: "Nano", vtype: Cars},
	}
	prepareLift(vehicles)
}
