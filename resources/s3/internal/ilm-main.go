package internal

import (
	"github.com/fatih/color"
	"github.com/minio/pkg/console"
)

const (
	ilmMainHeader         string = "Main-Heading"
	ilmThemeHeader        string = "Row-Header"
	ilmThemeRow           string = "Row-Normal"
	ilmThemeTick          string = "Row-Tick"
	ilmThemeExpiry        string = "Row-Expiry"
	ilmThemeResultSuccess string = "SuccessOp"
	ilmThemeResultFailure string = "FailureOp"
)

// Color scheme for the table
func setILMDisplayColorScheme() {
	console.SetColor(ilmMainHeader, color.New(color.Bold, color.FgHiRed))
	console.SetColor(ilmThemeRow, color.New(color.FgHiWhite))
	console.SetColor(ilmThemeHeader, color.New(color.Bold, color.FgHiGreen))
	console.SetColor(ilmThemeTick, color.New(color.FgGreen))
	console.SetColor(ilmThemeExpiry, color.New(color.BlinkRapid, color.FgGreen))
	console.SetColor(ilmThemeResultSuccess, color.New(color.FgGreen, color.Bold))
	console.SetColor(ilmThemeResultFailure, color.New(color.FgHiYellow, color.Bold))
}
