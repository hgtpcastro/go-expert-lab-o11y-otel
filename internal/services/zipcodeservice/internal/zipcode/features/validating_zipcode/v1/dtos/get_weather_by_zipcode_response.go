package dtos

type GetWeatherByZipCodeResponseDto struct {
	Cidade     string  `json:"city,omitempty"`
	Celsius    float64 `json:"temp_C,omitempty"`
	Fahrenheit float64 `json:"temp_F,omitempty"`
	Kelvin     float64 `json:"temp_K,omitempty"`
	Erro       string  `json:"erro,omitempty"`
}
