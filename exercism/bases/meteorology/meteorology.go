package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (t TemperatureUnit) String() string {
	unit := []string{"°C", "°F"}
	return unit[t]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
	return fmt.Sprintf("%v %v", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (t SpeedUnit) String() string {
	unit := []string{"km/h", "mph"}
	return unit[t]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (t Speed) String() string {
	return fmt.Sprintf("%v %v", t.magnitude, t.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (t MeteorologyData) String() string {
	return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity", t.location, t.temperature, t.windDirection, t.windSpeed, t.humidity)
}
