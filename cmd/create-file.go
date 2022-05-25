/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"test_toolkit/config"
	"test_toolkit/models"
	"test_toolkit/resources/data_factory"

	"github.com/spf13/cobra"
)

var createFileInput = models.FileCreateInput{}

// createFileCmd represents the createFile command
var createFileCmd = &cobra.Command{
	Use:   "create_file",
	Short: "Create local files",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create Files ...")
		data_factory.CreateFiles(createFileInput)
	},
}

func init() {
	rootCmd.AddCommand(createFileCmd)

	createFileCmd.PersistentFlags().StringVar(&createFileInput.LocalDataDir, "local_dir", config.RootDir+"/tmp/", "Local dir for save files")
	createFileCmd.PersistentFlags().IntVar(&createFileInput.RandomPercent, "random_percent", 100, "Percent of files with random data(0~100)% (default 100)")
	createFileCmd.PersistentFlags().IntVar(&createFileInput.EmptyPercent, "empty_percent", 0, "Percent of files with empty data(0~100)% (default 0)")
	createFileCmd.PersistentFlags().BoolVar(&createFileInput.RenameFile, "rename", false, "Rename files name every time if true (default false)")
	createFileCmd.PersistentFlags().StringArrayVar(&createFileInput.FileArgs, "file_args", []string{"txt:20:1k-10k", "dd:1:1mb"}, "Files configs array")
}
