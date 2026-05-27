

// Package weather represents a current weather condition. 
package weather

var (
    //CurrentCondition stores the latest weather condition.
	CurrentCondition string
    //CurrentLocation stores the name of the city.
	CurrentLocation  string
)

//Forecast updates the location and condition then return a weather summary.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
