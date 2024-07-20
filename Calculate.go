package main

import "math"

// StrategyCalculate TODO: What name to use
type Strategy struct {
	initAmount         float32
	simpleInterestRate float32
	years              uint8
	totalInterest      float64
	totalSaving        float64
}
type Calculate interface {
	TotalSavings() float64
	Interest() float64
}

type CompoundStrategy struct {
	Strategy
}

func (compoundStrategy *CompoundStrategy) TotalSavings() float64 {
	return float64(compoundStrategy.initAmount) + compoundStrategy.Interest()
}
func (compoundStrategy *CompoundStrategy) Interest() float64 {
	return float64(compoundStrategy.initAmount) * math.Pow(float64(1+compoundStrategy.simpleInterestRate), float64(compoundStrategy.years))
}

type SimpleStrategy struct {
	Strategy
}

func (simpleStrategy *SimpleStrategy) TotalSavings() float64 {
	interestAmount := simpleStrategy.Interest()
	return interestAmount + float64(simpleStrategy.initAmount)
}
func (simpleStrategy *SimpleStrategy) Interest() float64 {
	amount := interest(simpleStrategy.initAmount, simpleStrategy.simpleInterestRate)
	return amount * float64(simpleStrategy.years)
}

type DollarAverageCostStrategy struct {
	Strategy
}

func (dollarAverageCostStrategy *DollarAverageCostStrategy) TotalSavings() float64 {
	return 0
}
func (dollarAverageCostStrategy *DollarAverageCostStrategy) Interest() float64 {
	amount := interest(dollarAverageCostStrategy.initAmount, dollarAverageCostStrategy.simpleInterestRate)
	return amount * float64(dollarAverageCostStrategy.years)
}

func interest(value float32, rate float32) float64 {
	return float64(value + value*rate)
}
func profit(initValue float32, totalSaving float64) float64 {
	return totalSaving - float64(initValue)
}
