package weatherapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/config"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/otel/tracing"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/services/weatherservice/internal/service/weatherapi/converter"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/services/weatherservice/internal/weather/features/getting_weather_by_zipcode/v1/dtos"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

func GetWeatherByCity(
	ctx context.Context,
	log *zap.Logger,
	config *config.Config,
	city string,
	httpClient *resty.Client,
	tracer tracing.AppTracer,
) (dtos.GetWeatherByCityResponseDto, error) {
	ctx, span := tracer.Start(ctx, "function.GetWeatherByCity")
	defer span.End()

	key := config.WeatherExternalApi.Key
	api := config.WeatherExternalApi.Url

	params := url.Values{}
	params.Add("q", strings.ToLower(city))
	url := fmt.Sprintf(api, params.Encode(), key)
	log.Debug(url)

	span.SetAttributes(
		attribute.String("url", url),
	)

	resp, err := httpClient.R().Get(url)
	if err != nil {
		return dtos.GetWeatherByCityResponseDto{}, err
	}

	var weatherByCityDTO getWeatherByCityDTO
	err = json.Unmarshal(resp.Body(), &weatherByCityDTO)
	log.Sugar().Debugf("%v", weatherByCityDTO)

	if err != nil {
		return dtos.GetWeatherByCityResponseDto{}, err
	}

	tempC := weatherByCityDTO.Current.TempC
	tempF := converter.NewConverter().CelsiusToFahrenheit(tempC)
	tempK := converter.NewConverter().CelsiusToKelvin(tempC)

	responseDto := dtos.GetWeatherByCityResponseDto{
		Cidade:     city,
		Celsius:    tempC,
		Fahrenheit: tempF,
		Kelvin:     tempK,
	}

	// span.SetAttributes(
	// 	attribute.String("responseDto", fmt.Sprintf("%q", responseDto)),
	// )

	return responseDto, nil
}

type getWeatherByCityDTO struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}
