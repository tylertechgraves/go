package space

type Planet string

const earthYearSeconds = 31557600

var planetMap = map[Planet]float64{
	"Mercury": 0.240846 * earthYearSeconds,
	"Venus":   0.61519726 * earthYearSeconds,
	"Earth":   earthYearSeconds,
	"Mars":    1.8808158 * earthYearSeconds,
	"Jupiter": 11.862615 * earthYearSeconds,
	"Saturn":  29.447498 * earthYearSeconds,
	"Uranus":  84.016846 * earthYearSeconds,
	"Neptune": 164.79132 * earthYearSeconds,
}

func Age(seconds float64, planet Planet) float64 {
	return seconds / planetMap[planet]
}
