package zipcode

import "github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/fxapp/contracts"

type ZipCodeServiceConfigurator struct {
	contracts.Application
	// infrastructureConfigurator *infrastructure.InfrastructureConfigurator
	// ordersModuleConfigurator   *configurations.OrdersModuleConfigurator
}

func NewZipCodeServiceConfigurator(
	app contracts.Application,
) *ZipCodeServiceConfigurator {
	// infraConfigurator := infrastructure.NewInfrastructureConfigurator(app)
	// ordersModuleConfigurator := configurations.NewOrdersModuleConfigurator(app)

	return &ZipCodeServiceConfigurator{
		Application: app,
		// infrastructureConfigurator: infrastructure.NewInfrastructureConfigurator(app),
		// ordersModuleConfigurator:   configurations.NewOrdersModuleConfigurator(app),
	}
}

func (ic *ZipCodeServiceConfigurator) ConfigureZipCode() {
	// Shared
	// Infrastructure
	// ic.infrastructureConfigurator.ConfigInfrastructures()

	// Shared
	// Zipcode service configurations

	// Modules
	// Order module
	// ic.zipcodeModuleConfigurator.ConfigureZipCodeModule()
}

func (ic *ZipCodeServiceConfigurator) MapZipCodeEndpoints() {}
