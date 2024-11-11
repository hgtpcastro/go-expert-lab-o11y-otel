package zipcode

import (
	"go.uber.org/fx"

	"github.com/go-playground/validator"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/http/route"
	routeContracts "github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/http/route/contracts"
	serverContracts "github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/http/server/contracts"
	validateZipCodeV1 "github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/services/zipcodeservice/internal/zipcode/features/validating_zipcode/v1/endpoints"
	"github.com/labstack/echo/v4"
)

var Module = fx.Module(
	"zipcode",
	fx.Provide(fx.Annotate(func(server serverContracts.EchoHttpServer) *echo.Group {
		var g *echo.Group

		server.RouteBuilder().RegisterGroupFunc("/api/v1", func(v1 *echo.Group) {
			group := v1.Group("/zipcode")
			g = group
		})

		server.SetupDefaultMiddlewares()

		return g
	}, fx.ResultTags(`name:"zipcode-echo-group"`))),

	fx.Provide(
		route.AsRoute(validateZipCodeV1.NewValidateZipCodeEndPoint, "zipcode-routes"),
	),

	fx.Provide(validator.New),

	fx.Options(fx.Invoke(fx.Annotate(func(endpoints []routeContracts.Endpoint) {
		for _, endpoint := range endpoints {
			endpoint.MapEndpoint()
		}
	}, fx.ParamTags(`group:"zipcode-routes"`)))),
)
