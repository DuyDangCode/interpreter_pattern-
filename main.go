package main

import "fmt"

func main() {
	maxDistance := DistanceExpression{500}
	maxCost := CostExpression{1000}
	rule1 := AndExpression{&maxDistance, &maxCost}

	maxTime := TimeExpression{5}
	perishableCargo := CargoTypeExpression{"Perishable"}
	rule2 := OrExpression{&maxTime, &perishableCargo}

	route1 := Route{450, 800, 4, "Perishable"}
	route2 := Route{650, 1200, 6, "Non-Perishable"}
	route3 := Route{350, 700, 3, "Non-Perishable"}

	fmt.Println("Route 1, Rule 1:", rule1.interpret(route1))
	fmt.Println("Route 2, Rule 1:", rule1.interpret(route2))
	fmt.Println("Route 3, Rule 1:", rule1.interpret(route3))
	fmt.Println("Route 1, Rule 2:", rule2.interpret(route1))
	fmt.Println("Route 2, Rule 2:", rule2.interpret(route2))
	fmt.Println("Route 3, Rule 2:", rule2.interpret(route3))
}
