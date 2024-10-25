package main

import (
	"os"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "weather-microservice",
	Short:            "weather-microservice",
	Long:             `This is a command runner or cli for api architecture in golang.`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		//app.NewApp().Run()
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
