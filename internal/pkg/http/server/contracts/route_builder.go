package contracts

import "github.com/labstack/echo/v4"

type RouteBuilder struct {
	echo *echo.Echo
}

func NewRouteBuilder(echo *echo.Echo) *RouteBuilder {
	return &RouteBuilder{echo: echo}
}

func (r *RouteBuilder) RegisterGroupFunc(groupName string, builder func(g *echo.Group)) *RouteBuilder {
	builder(r.echo.Group(groupName))

	return r
}
