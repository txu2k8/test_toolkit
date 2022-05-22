package cmd

import (
	"fmt"

	// "test_toolkit/resources/ilm"

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
		fmt.Println("add called")
	},
}

func init() {
	ilmCmd.AddCommand(ilmAddCmd)
	ilmAddCmd.PersistentFlags().StringVar(&lfcCfg.ExpiryDate, "s3_ip", "", "S3 server IP address")

}

// Calls SetBucketLifecycle with the XML representation of lifecycleConfiguration type.
func mainILMAdd(cmd *cobra.Command) error {
	// connect S3
	// client := s3.NewS3Client(s3Cfg.Endpoint, s3Cfg.S3AccessID, s3Cfg.S3SecretKey)

	// Configuration that is already set.
	// lfcCfg, err := client.GetLifecycle(ctx)
	// if err != nil {
	// 	if e := err.ToGoError(); minio.ToErrorResponse(e).Code == "NoSuchLifecycleConfiguration" {
	// 		lfcCfg = lifecycle.NewConfiguration()
	// 	} else {
	// 		logger.Fatalf(err.Trace(args...), "Unable to fetch lifecycle rules for "+urlStr)
	// 	}
	// }

	// opts, err := ilm.GetLifecycleOptions(cliCtx)
	// fatalIf(err.Trace(args...), "Unable to generate new lifecycle rules for the input")

	// lfcCfg, err = opts.ToConfig(lfcCfg)
	// fatalIf(err.Trace(args...), "Unable to generate new lifecycle rules for the input")

	// fatalIf(client.SetLifecycle(ctx, lfcCfg).Trace(urlStr), "Unable to add this lifecycle rule")

	// printMsg(ilmAddMessage{
	// 	Status: "success",
	// 	Target: urlStr,
	// 	ID:     opts.ID,
	// })

	return nil
}
