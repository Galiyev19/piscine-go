package main

import "fmt"

var Pilot struct {
	Name     string
	Life     int
	Age      int
	Aircraft int
}

func main() {
	donnie := Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1

	fmt.Println(donnie)
}

const AIRCRAFT1 = 1