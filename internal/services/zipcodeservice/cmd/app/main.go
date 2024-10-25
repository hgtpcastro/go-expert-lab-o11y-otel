package main

import (
	"os"

	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/services/zipcodeservice/internal/shared/app"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "zipcode-microservice",
	Short:            "zipcode-microservice",
	Long:             `This is a command runner or cli for api architecture in golang.`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		app.NewApp().Run()
	},
}

func main() {
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("Zip Code", pterm.FgLightGreen.ToStyle()),
		putils.LettersFromStringWithStyle(" Service", pterm.FgLightYellow.ToStyle())).
		Render()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
