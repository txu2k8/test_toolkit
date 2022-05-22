package internal

import (
	"context"
	"encoding/json"
	"test_toolkit/resources/s3/internal/ilm"

	"github.com/minio/cli"
	"github.com/minio/mc/pkg/probe"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/lifecycle"
	"github.com/minio/pkg/console"
)

type ilmAddMessage struct {
	Status string `json:"status"`
	Target string `json:"target"`
	ID     string `json:"id"`
}

func (i ilmAddMessage) String() string {
	return console.Colorize(ilmThemeResultSuccess, "Lifecycle configuration rule added with ID `"+i.ID+"` to "+i.Target+".")
}

func (i ilmAddMessage) JSON() string {
	msgBytes, e := json.MarshalIndent(i, "", " ")
	fatalIf(probe.NewError(e), "Unable to encode as JSON.")
	return string(msgBytes)
}

// Validate user given arguments
func checkILMAddSyntax(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		cli.ShowCommandHelpAndExit(ctx, "add", globalErrorExitStatus)
	}
}

// Calls SetBucketLifecycle with the XML representation of lifecycleConfiguration type.
func MainILMAdd(cliCtx *cli.Context) error {
	ctx, cancelILMAdd := context.WithCancel(globalContext)
	defer cancelILMAdd()

	checkILMAddSyntax(cliCtx)
	setILMDisplayColorScheme()
	args := cliCtx.Args()
	urlStr := args.Get(0)

	client, err := newClient(urlStr)
	fatalIf(err.Trace(urlStr), "Unable to initialize client for "+urlStr)

	// Configuration that is already set.
	lfcCfg, err := client.GetLifecycle(ctx)
	if err != nil {
		if e := err.ToGoError(); minio.ToErrorResponse(e).Code == "NoSuchLifecycleConfiguration" {
			lfcCfg = lifecycle.NewConfiguration()
		} else {
			fatalIf(err.Trace(args...), "Unable to fetch lifecycle rules for "+urlStr)
		}
	}

	opts, err := ilm.GetLifecycleOptions(cliCtx)
	fatalIf(err.Trace(args...), "Unable to generate new lifecycle rules for the input")

	lfcCfg, err = opts.ToConfig(lfcCfg)
	fatalIf(err.Trace(args...), "Unable to generate new lifecycle rules for the input")

	fatalIf(client.SetLifecycle(ctx, lfcCfg).Trace(urlStr), "Unable to add this lifecycle rule")

	printMsg(ilmAddMessage{
		Status: "success",
		Target: urlStr,
		ID:     opts.ID,
	})

	return nil
}
