package main

import (
	"os"

	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/config"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/http/client"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/http/server"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/log"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/otel/tracing"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/services/weatherservice/internal/weather"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var rootCmd = &cobra.Command{
	Use:              "weather-microservice",
	Short:            "weather-microservice",
	Long:             `This is a command runner or cli for api architecture in golang.`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			config.Module,
			log.Module,
			server.Module,
			client.Module,
			tracing.Module,
			weather.Module,
		)
		app.Run()
	},
}

func main() {
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("Weather", pterm.FgLightGreen.ToStyle()),
		putils.LettersFromStringWithStyle(" Service", pterm.FgLightYellow.ToStyle())).
		Render()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
