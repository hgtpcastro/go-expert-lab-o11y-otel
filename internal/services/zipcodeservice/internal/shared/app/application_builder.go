package app

import (
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/fxapp"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/fxapp/contracts"
)

type ZipCodeApplicationBuilder struct {
	contracts.ApplicationBuilder
}

func NewZipCodeApplicationBuilder() *ZipCodeApplicationBuilder {
	return &ZipCodeApplicationBuilder{fxapp.NewApplicationBuilder()}
}

func (a *ZipCodeApplicationBuilder) Build() *ZipCodeApplication {
	return NewZipCodeApplication(
		a.GetProvides(),
		a.GetDecorates(),
		a.Options(),
		a.Logger(),
		a.Environment(),
	)
}
