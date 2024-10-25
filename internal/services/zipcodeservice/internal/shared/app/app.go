package app

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	// configure dependencies
	appBuilder := NewZipCodeApplicationBuilder()
	// appBuilder.ProvideModule(zipcode.ZipCodeServiceModule)

	app := appBuilder.Build()

	// configure application
	app.ConfigureZipCode()

	app.MapZipCodeEndpoints()

	app.Logger().Info("Starting zipcode_service application")
	app.Run()
}
