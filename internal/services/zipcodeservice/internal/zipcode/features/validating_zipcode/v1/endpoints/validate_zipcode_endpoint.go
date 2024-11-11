package endpoints

import (
	"net/http"

	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/services/zipcodeservice/internal/zipcode/contracts/params"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/services/zipcodeservice/internal/zipcode/features/validating_zipcode/v1/dtos"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/services/zipcodeservice/internal/zipcode/service/weatherapi"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/attribute"
)

type validateZipCodeEndPoint struct {
	params.ZipCodeRouteParams
}

func NewValidateZipCodeEndPoint(params params.ZipCodeRouteParams) *validateZipCodeEndPoint {
	return &validateZipCodeEndPoint{ZipCodeRouteParams: params}
}

func (ep *validateZipCodeEndPoint) MapEndpoint() {
	ep.Group.POST("/validate", ep.handler())
}

func (ep *validateZipCodeEndPoint) handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		ctx, span := ep.Tracer.Start(ctx, "validateZipCodeEndPoint.handler")
		defer span.End()

		var dto dtos.ValidateZipCodeRequestDto

		err := c.Bind(&dto)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		zipcode := dto.Cep

		if zipcode == "" {
			return c.JSON(http.StatusUnprocessableEntity, "invalid zipcode")
		}

		ep.Validator.SetTagName("cep")
		err = ep.Validator.Var(zipcode, "required,len=8,numeric")
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, "invalid zipcode")
		}

		span.SetAttributes(
			attribute.String("zip-code", zipcode),
		)

		data, err := weatherapi.GetWeatherByZipCode(
			ctx,
			ep.Log,
			ep.Config,
			zipcode,
			ep.HttpClient,
			ep.Tracer,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		if data.Erro == `true` {
			return c.JSON(http.StatusNotFound, "can not find zipcode")
		}

		return c.JSON(http.StatusOK, data)
	}
}
