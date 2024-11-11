package params

import (
	"github.com/go-playground/validator"
	"github.com/go-resty/resty/v2"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/config"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/otel/tracing"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ZipCodeRouteParams struct {
	fx.In

	Config     *config.Config
	Log        *zap.Logger
	HttpClient *resty.Client
	Group      *echo.Group `name:"zipcode-echo-group"`
	Validator  *validator.Validate
	Tracer     tracing.AppTracer
}
