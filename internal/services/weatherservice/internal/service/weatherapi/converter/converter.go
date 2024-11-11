package converter

type weatherConverter struct {
}

func NewConverter() *weatherConverter {
	return &weatherConverter{}
}

func (c *weatherConverter) CelsiusToFahrenheit(v float64) float64 {
	return v*1.8 + 32
}

func (c *weatherConverter) CelsiusToKelvin(v float64) float64 {
	return v + 273.15
}
