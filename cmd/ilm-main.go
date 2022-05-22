package cmd

import (
	"github.com/spf13/cobra"
)

// ilmCmd represents the ilm command
var ilmCmd = &cobra.Command{
	Use:   "ilm",
	Short: "manage bucket lifecycle",
	Long:  ``,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("ilm called")
	// },
}

func init() {
	rootCmd.AddCommand(ilmCmd)
}
