package weather

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/http/route"
	routeContracts "github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/http/route/contracts"
	serverContracts "github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/http/server/contracts"
	getWeatherByZipCodeV1 "github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/services/weatherservice/internal/weather/features/getting_weather_by_zipcode/v1/endpoints"
)

var Module = fx.Module(
	"weather",
	fx.Provide(fx.Annotate(func(server serverContracts.EchoHttpServer) *echo.Group {
		var g *echo.Group

		server.RouteBuilder().RegisterGroupFunc("/api/v1", func(v1 *echo.Group) {
			group := v1.Group("/weather")
			g = group
		})

		server.SetupDefaultMiddlewares()

		return g
	}, fx.ResultTags(`name:"weather-echo-group"`))),

	fx.Provide(
		route.AsRoute(getWeatherByZipCodeV1.NewGetWeatherByZipCodeEndPoint, "weather-routes"),
	),

	fx.Provide(validator.New),

	fx.Options(fx.Invoke(fx.Annotate(func(endpoints []routeContracts.Endpoint) {
		for _, endpoint := range endpoints {
			endpoint.MapEndpoint()
		}
	}, fx.ParamTags(`group:"weather-routes"`)))),
)
