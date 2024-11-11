package dtos

type GetWeatherByCityResponseDto struct {
	Cidade     string  `json:"city"`
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
	Erro       string  `json:"erro,omitempty"`
}
