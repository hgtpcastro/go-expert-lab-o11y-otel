package weatherapi

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/config"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/otel/tracing"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/services/zipcodeservice/internal/zipcode/features/validating_zipcode/v1/dtos"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
)

func GetWeatherByZipCode(
	ctx context.Context,
	log *zap.Logger,
	config *config.Config,
	zipCode string,
	httpClient *resty.Client,
	tracer tracing.AppTracer,
) (dtos.GetWeatherByZipCodeResponseDto, error) {
	ctx, span := tracer.Start(ctx, "function.GetCityByZipCode")
	defer span.End()

	api := config.WeatherApi.Url
	url := fmt.Sprintf(api, zipCode)

	span.SetAttributes(
		attribute.String("url", url),
	)

	req := httpClient.R().SetContext(ctx)
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	resp, err := req.Get(url)
	if err != nil {
		return dtos.GetWeatherByZipCodeResponseDto{}, err
	}

	var data dtos.GetWeatherByZipCodeResponseDto
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return dtos.GetWeatherByZipCodeResponseDto{}, err
	}

	return data, nil
}
