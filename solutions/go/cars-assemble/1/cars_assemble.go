package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	production := float64(productionRate)
    assemblyPerHour := successRate / 100
    return assemblyPerHour * production
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    carsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
    return int(carsPerHour / 60) 
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    count := uint(carsCount)
	costs := count / 10
    remainCars := count % 10
    totalCount := (costs * 95000) + (remainCars * 10000)
    return totalCount
}
