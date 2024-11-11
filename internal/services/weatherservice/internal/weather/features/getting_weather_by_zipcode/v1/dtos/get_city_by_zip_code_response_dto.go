package dtos

type GetCityByZipCodeResponseDTO struct {
	Localidade string `json:"localidade,omitempty"`
	Erro       string `json:"erro,omitempty"`
}
