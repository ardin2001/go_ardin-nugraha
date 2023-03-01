package main

import "fmt"

type Car struct {
	tipe   string
	fuelln float64
}

func (toyotaf Car) CapacityRange() float64 {
	travel := 0.0
	func() {
		if toyotaf.tipe == "pertalite" {
			travel = 1.5
		} else if toyotaf.tipe == "pertamax" {
			travel = 2
		} else if toyotaf.tipe == "solar" {
			travel = 2.5
		} else {
			travel = 1
		}
	}()
	TotalTravel := toyotaf.fuelln / travel
	return TotalTravel
}
func main() {
	toyota := Car{
		tipe:   "pertalite",
		fuelln: 4.5,
	}

	fmt.Println("Tipe bahan bakar :", toyota.tipe)
	fmt.Println("Bahan bakar yang sedang terisi :", toyota.fuelln, "L")
	fmt.Println("Total jarak tempuh :", toyota.CapacityRange(), "/mill")
}
