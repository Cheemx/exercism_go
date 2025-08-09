// Package weather provides tools to get current weather condition of a particular city.
package weather

// CurrentCondition variable is defined to store the weather condition of particular location.
var CurrentCondition string
// CurrentLocation variable is defined to store the location for which weather forecast will be requested.
var CurrentLocation string

// Forecast function accepts two parameters the name of the city and it's current weather condition in string format and returns the formatted weather condition in string format.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
