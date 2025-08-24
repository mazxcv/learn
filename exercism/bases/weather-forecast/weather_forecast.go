// Package weather provides current weather condition.
package weather

// CurrentCondition represents current weather condition.
var CurrentCondition string

// CurrentLocation represents current city location.
var CurrentLocation string

// Forecast shows the weather forecast in a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
