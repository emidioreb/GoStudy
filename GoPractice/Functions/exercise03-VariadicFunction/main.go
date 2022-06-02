package main

import "fmt"

func calculateAverageWeight(peopleWeight... float64) {
	lenght := float64(len(peopleWeight))
	var sumWeight float64
	sumWeight = 0
	for _, weight := range peopleWeight{
		sumWeight += weight
	}
	result := sumWeight/lenght
	fmt.Println(result)
}
func main() {
	calculateAverageWeight(2.00, 2.00, 2.00)
}