package zipcodeapi

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/config"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/otel/tracing"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/services/weatherservice/internal/weather/features/getting_weather_by_zipcode/v1/dtos"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

func GetCityByZipCode(
	ctx context.Context,
	log *zap.Logger,
	config *config.Config,
	zipCode string,
	httpClient *resty.Client,
	tracer tracing.AppTracer,
) (dtos.GetCityByZipCodeResponseDTO, error) {
	ctx, span := tracer.Start(ctx, "function.GetCityByZipCode")
	defer span.End()

	api := config.ZipCodeExternalApi.Url
	url := fmt.Sprintf(api, zipCode)
	log.Sugar().Debug(url)

	span.SetAttributes(
		attribute.String("url", url),
	)

	resp, err := httpClient.R().Get(url)
	if err != nil {
		return dtos.GetCityByZipCodeResponseDTO{}, err
	}

	var data dtos.GetCityByZipCodeResponseDTO
	err = json.Unmarshal(resp.Body(), &data)

	log.Sugar().Debugf("%q", data)

	if err != nil {
		return dtos.GetCityByZipCodeResponseDTO{}, err
	}

	span.SetAttributes(
		attribute.String("city", data.Localidade),
	)

	return data, nil
}
