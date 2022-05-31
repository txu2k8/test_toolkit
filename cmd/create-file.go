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
	Short: "Create local random files",
	Long:  ``,
	Example: `
	// Default create 20 txt files with random of size 1k~10k, 1 data file with size of 1MB
	test_toolkit create_file
	
	// Create 100 txt files with random of size 1k~1MB
	test_toolkit create_file --file_args "txt:20:1k-10k"

	// Create files with Spec local dirs
	test_toolkit create_file --local_dir /data/dir/

	// Create files with 50 percent of files random string
	test_toolkit create_file --random_percent 50

	// Create files with 20 percent of files empty
	test_toolkit create_file --empty_percent 20
	`,
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
	createFileCmd.PersistentFlags().StringArrayVar(&createFileInput.FileArgs, "file_args", []string{"txt:20:1k-10k", "data:1:1mb"}, "Files configs array")
}
