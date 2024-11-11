package route

import (
	"fmt"

	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/http/route/contracts"

	"go.uber.org/fx"
)

func AsRoute(handler interface{}, routeGroupName string) interface{} {
	return fx.Annotate(
		handler,
		fx.As(new(contracts.Endpoint)),
		fx.ResultTags(fmt.Sprintf(`group:"%s"`, routeGroupName)),
	)
}
