package main

import "fmt"

func main() {
	maxMph := 50
	passengerCapacity := 4
	airbags := true
	familyCar := passengerCapacity > 2 && airbags
	sportCat := maxMph > 100 || passengerCapacity == 2
	canCategorize := !familyCar && !sportCat
	fmt.Println(familyCar, sportCat, canCategorize)
}
