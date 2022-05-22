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
	logger   = logging.MustGetLogger("test")
	runTimes int      // runTimes
	debug    bool     // debug modle
	caseList []string // Case List
	s3Cfg    models.S3Config
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.test_toolkit.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().IntVar(&runTimes, "run_times", 1, "Run test case with iteration loop")
	rootCmd.PersistentFlags().StringArrayVar(&caseList, "case", []string{}, "Test Case Array (default value in sub-command)")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable console debug loglevel if true (default false)")

	rootCmd.PersistentFlags().StringVar(&s3Cfg.S3AccessID, "s3_access_id", "", "S3 access ID")
	rootCmd.PersistentFlags().StringVar(&s3Cfg.S3SecretKey, "s3_secret_key", "", "S3 access secret key")

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
	logger.Infof("Args: platform %s", strings.Join(os.Args[1:], " "))
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
