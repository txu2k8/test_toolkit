package cmd

import (
	"github.com/minio/mc/cmd/ilm"
	"github.com/spf13/cobra"
)

var lfcCfg = ilm.LifecycleOptions{}

// addCmd represents the add command
var ilmAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add a lifecycle configuration rule for a bucket",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("add called")
		// cliCtx := cli.NewContext()
		// internal.MainILMAdd(cliCtx)
	},
}

func init() {
	ilmCmd.AddCommand(ilmAddCmd)
	ilmAddCmd.PersistentFlags().StringVar(&lfcCfg.ExpiryDate, "expiry-date", "", "ExpiryDate")

}
