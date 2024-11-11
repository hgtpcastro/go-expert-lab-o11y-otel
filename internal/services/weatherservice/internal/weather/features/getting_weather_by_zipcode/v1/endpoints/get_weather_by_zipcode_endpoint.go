package endpoints

import (
	"net/http"

	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/services/weatherservice/internal/service/weatherapi"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/services/weatherservice/internal/service/zipcodeapi"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/services/weatherservice/internal/weather/contracts/params"
	"github.com/labstack/echo/v4"
)

type getWeatherByZipCodeEndPoint struct {
	params.WeatherRouteParams
}

func NewGetWeatherByZipCodeEndPoint(params params.WeatherRouteParams) *getWeatherByZipCodeEndPoint {
	return &getWeatherByZipCodeEndPoint{WeatherRouteParams: params}
}

func (ep *getWeatherByZipCodeEndPoint) MapEndpoint() {
	ep.Group.GET("/:zipcode", ep.handler())
}

func (ep *getWeatherByZipCodeEndPoint) handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		ctx, span := ep.Tracer.Start(ctx, "getWeatherByZipCodeEndPoint.handler")
		defer span.End()

		zipcode := c.Param("zipcode")

		if zipcode == "" {
			return c.JSON(http.StatusUnprocessableEntity, "invalid zipcode")
		}

		ep.Validator.SetTagName("zipcode")
		err := ep.Validator.Var(zipcode, "required,len=8,numeric")
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, "invalid zipcode")
		}

		dataCity, err := zipcodeapi.GetCityByZipCode(
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

		if dataCity.Erro == `true` {
			return c.JSON(http.StatusNotFound, dataCity)
		}

		dataWeather, err := weatherapi.GetWeatherByCity(
			ctx,
			ep.Log,
			ep.Config,
			dataCity.Localidade,
			ep.HttpClient,
			ep.Tracer,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		if dataWeather.Erro == `true` {
			return c.JSON(http.StatusNotFound, dataWeather)
		}

		return c.JSON(http.StatusOK, dataWeather)
	}
}
