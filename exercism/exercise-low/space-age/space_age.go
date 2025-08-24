package space

type Planet string

// func FactorAgePlanet return orbital period planet of Earth years
func FactorAgePlanet() map[Planet]float64 {
	return map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
}

const OrbitalErathPeriod = 31557600

// func Age calculate how old someone in planet
func Age(seconds float64, planet Planet) float64 {
	factor, ok := FactorAgePlanet()[planet]
	if !ok {
		return -1
	}
	return seconds / OrbitalErathPeriod / factor
}
