package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"
	"test_toolkit/config"
	"test_toolkit/models"
	"test_toolkit/pkg/tlog"

	"github.com/chenhg5/collection"
	"github.com/op/go-logging"
	"github.com/spf13/cobra"
)

var (
	logger      = logging.MustGetLogger("test")
	repeatCount int      // repeat count
	debug       bool     // console show debug modle
	caseList    []string // Case List
	s3Cfg       models.S3Config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "test_toolkit",
	Short: "A project for test toolkit",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	initLogging()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVar(&repeatCount, "count", 1, "Run test case repeat count")
	rootCmd.PersistentFlags().StringArrayVar(&caseList, "case", []string{}, "Test Case Array (default value in sub-command)")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable console debug loglevel if true (default false)")

	rootCmd.PersistentFlags().StringVar(&s3Cfg.Endpoint, "endpoint", "", "S3 endpoint")
	rootCmd.PersistentFlags().StringVar(&s3Cfg.AccessKey, "access-key", "", "S3 access ID")
	rootCmd.PersistentFlags().StringVar(&s3Cfg.SecretKey, "secret-key", "", "S3 access secret key")

}

// initLogging initialize the logging configs
func initLogging() {
	dir, _ := os.Getwd()
	fileLogName := "test_toolkit"
	fileLogPath := path.Join(dir, "log")
	timeStr := config.TimeStr // time.Now().Format("20060102150405")
	for _, v := range stripArgs() {
		fileLogName = fmt.Sprintf("%s-%s", fileLogName, v)
		fileLogPath = path.Join(fileLogPath, v)
	}
	fileLogName = fmt.Sprintf("%s-%s.log", fileLogName, timeStr)
	fileLogPath = path.Join(fileLogPath, fileLogName)
	consoleLoglevel := logging.INFO
	if collection.Collect(os.Args).Contains("--debug") {
		consoleLoglevel = logging.DEBUG
	}

	conf := tlog.NewOptions(
		tlog.OptionSetFileLogPath(fileLogPath),
		tlog.OptionSetConsoleLogLevel(consoleLoglevel),
	)
	conf.InitLogging()
	logger.Infof("Args: test_toolkit %s", strings.Join(os.Args[1:], " "))
}

func stripArgs() []string {
	commands := []string{}
	args := os.Args[1:]
	ps := ""
	for len(args) > 0 {
		s := args[0]
		args = args[1:]
		switch {
		case s == "--":
			// "--" terminates the flags
			break
		case strings.HasPrefix(s, "--") && !strings.Contains(s, "="):
			// If '--flag arg' then
			// delete arg from args.
			fallthrough // (do the same as below)
		case strings.HasPrefix(s, "-") && !strings.Contains(s, "=") && len(s) == 2:
			// If '-f arg' then
			// delete 'arg' from args or break the loop if len(args) <= 1.
			if len(args) <= 1 {
				break
			} else {
				args = args[1:]
				continue
			}
		case s != "" && !strings.HasPrefix(s, "-") && !strings.HasPrefix(ps, "-"):
			commands = append(commands, s)
		}
		ps = s
	}

	return commands
}
