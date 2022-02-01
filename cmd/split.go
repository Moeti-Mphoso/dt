/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// splitCmd represents the split command
var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "split a csv file",
	Long: `split large csv file into smaller files:
dt split file # splits a file in 2 and appends numbers to files, file1 and file2
dt split -n x file # splits a file into x files.`,
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open(args[0])
		defer f.Close()
		if err != nil {
			log.Fatalf("bloody error: %s", err)
		}
		fmt.Printf("split file %s, %d times\n", args[0], count)
	},
}

// command flags
var count int

func init() {
	rootCmd.AddCommand(splitCmd)
	splitCmd.Flags().IntVarP(&count, "number", "n", 2, "number of output files")
}

func splitFile(file io.Reader, count int) {

}
