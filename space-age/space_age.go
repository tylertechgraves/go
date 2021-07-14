package space

type Planet string

const earthYearSeconds = 31557600

var planetMap = map[Planet]func(float64, Planet) float64{
	"Mercury": func(seconds float64, planet Planet) float64 {
		return seconds / (earthYearSeconds * 0.240846)
	},
	"Venus": func(seconds float64, planet Planet) float64 {
		return seconds / (0.61519726 * earthYearSeconds)
	},
	"Earth": func(seconds float64, planet Planet) float64 {
		return seconds / earthYearSeconds
	},
	"Mars": func(seconds float64, planet Planet) float64 {
		return seconds / (1.8808158 * earthYearSeconds)
	},
	"Jupiter": func(seconds float64, planet Planet) float64 {
		return seconds / (11.862615 * earthYearSeconds)
	},
	"Saturn": func(seconds float64, planet Planet) float64 {
		return seconds / (29.447498 * earthYearSeconds)
	},
	"Uranus": func(seconds float64, planet Planet) float64 {
		return seconds / (84.016846 * earthYearSeconds)
	},
	"Neptune": func(seconds float64, planet Planet) float64 {
		return seconds / (164.79132 * earthYearSeconds)
	},
}

func Age(seconds float64, planet Planet) float64 {
	return planetMap[planet](seconds, planet)
}
