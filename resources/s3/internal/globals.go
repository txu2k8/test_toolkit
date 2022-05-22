package internal

import (
	"context"
	"crypto/x509"
	"fmt"
	"net/url"
	"os"

	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"github.com/minio/cli"
	"github.com/minio/pkg/console"
)

const (
	globalMCConfigVersion = "10"

	globalMCConfigFile = "config.json"
	globalMCCertsDir   = "certs"
	globalMCCAsDir     = "CAs"

	// session config and shared urls related constants
	globalSessionDir           = "session"
	globalSharedURLsDataDir    = "share"
	globalSessionConfigVersion = "8"

	// Profile directory for dumping profiler outputs.
	globalProfileDir = "profile"

	// Global error exit status.
	globalErrorExitStatus = 1

	// Global CTRL-C (SIGINT, #2) exit status.
	globalCancelExitStatus = 130

	// Global SIGKILL (#9) exit status.
	globalKillExitStatus = 137

	// Global SIGTERM (#15) exit status
	globalTerminatExitStatus = 143
)

var (
	globalQuiet          = false  // Quiet flag set via command line
	globalJSON           = false  // Json flag set via command line
	globalJSONLine       = false  // Print json as single line.
	globalDebug          = false  // Debug flag set via command line
	globalNoColor        = false  // No Color flag set via command line
	globalInsecure       = false  // Insecure flag set via command line
	globalDevMode        = false  // dev flag set via command line
	globalSubnetProxyURL *url.URL // Proxy to be used for communication with subnet

	globalContext, globalCancel = context.WithCancel(context.Background())
)

var (
	// Terminal width
	globalTermWidth int

	// CA root certificates, a nil value means system certs pool will be used
	globalRootCAs *x509.CertPool
)
var (
	// Check if we stderr, stdout are dumb terminals, we do not apply
	// ansi coloring on dumb terminals.
	isTerminal = func() bool {
		return isatty.IsTerminal(os.Stdout.Fd()) && isatty.IsTerminal(os.Stderr.Fd())
	}

	colorCyanBold = func() func(a ...interface{}) string {
		if isTerminal() {
			color.New(color.FgCyan, color.Bold).SprintFunc()
		}
		return fmt.Sprint
	}()

	colorYellowBold = func() func(format string, a ...interface{}) string {
		if isTerminal() {
			return color.New(color.FgYellow, color.Bold).SprintfFunc()
		}
		return fmt.Sprintf
	}()

	colorGreenBold = func() func(format string, a ...interface{}) string {
		if isTerminal() {
			return color.New(color.FgGreen, color.Bold).SprintfFunc()
		}
		return fmt.Sprintf
	}()
)

// Set global states. NOTE: It is deliberately kept monolithic to ensure we dont miss out any flags.
func setGlobals(quiet, debug, json, noColor, insecure, devMode bool) {
	globalQuiet = globalQuiet || quiet
	globalDebug = globalDebug || debug
	globalJSONLine = !isTerminal() && json
	globalJSON = globalJSON || json
	globalNoColor = globalNoColor || noColor || globalJSONLine
	globalInsecure = globalInsecure || insecure
	globalDevMode = globalDevMode || devMode

	// Disable colorified messages if requested.
	if globalNoColor || globalQuiet {
		console.SetColorOff()
	}
}

// Set global states. NOTE: It is deliberately kept monolithic to ensure we dont miss out any flags.
func setGlobalsFromContext(ctx *cli.Context) error {
	quiet := ctx.IsSet("quiet") || ctx.GlobalIsSet("quiet")
	debug := ctx.IsSet("debug") || ctx.GlobalIsSet("debug")
	json := ctx.IsSet("json") || ctx.GlobalIsSet("json")
	noColor := ctx.IsSet("no-color") || ctx.GlobalIsSet("no-color")
	insecure := ctx.IsSet("insecure") || ctx.GlobalIsSet("insecure")
	devMode := ctx.IsSet("dev") || ctx.GlobalIsSet("dev")

	setGlobals(quiet, debug, json, noColor, insecure, devMode)
	return nil
}
