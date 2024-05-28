package main

import "strings"

type IExpression interface {
	interpret(Route) bool
}

// terminal expression
type DistanceExpression struct {
	maxDistance float64
}

func (d *DistanceExpression) interpret(route Route) bool {
	return route.distance <= d.maxDistance
}

type CostExpression struct {
	maxCost float64
}

func (c *CostExpression) interpret(route Route) bool {
	return route.cost <= c.maxCost
}

type TimeExpression struct {
	maxTime float64
}

func (t *TimeExpression) interpret(route Route) bool {
	return route.time <= t.maxTime
}

type CargoTypeExpression struct {
	cargoType string
}

func (cargo *CargoTypeExpression) interpret(route Route) bool {
	return strings.EqualFold(route.cargoType, cargo.cargoType)
}

// non terminal expression
type AndExpression struct {
	expr1 IExpression
	expr2 IExpression
}

func (ae *AndExpression) interpret(route Route) bool {
	return ae.expr1.interpret(route) && ae.expr2.interpret(route)
}

type OrExpression struct {
	expr1 IExpression
	expr2 IExpression
}

func (oe *OrExpression) interpret(route Route) bool {
	return oe.expr1.interpret(route) || oe.expr2.interpret(route)
}
