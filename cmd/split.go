/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"

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
		fmt.Printf("split file %s, %d times\n", args[0], count)

		SplitFile(args[0], count)

	},
}

// command flags
var count int

func init() {
	rootCmd.AddCommand(splitCmd)
	splitCmd.Flags().IntVarP(&count, "number", "n", 2, "number of output files")
}

type FileStats struct {
	Info  fs.FileInfo
	Lines int32
	// can't read stats for file with more that 2147483647 lines (limited by int32)
}

func GetFileStats(filename string) (FileStats, error) {
	file, err := os.Open(filename)
	if err != nil {
		return FileStats{}, err
	}
	defer file.Close()
	// get number of lines from file
	scanner := bufio.NewScanner(file)
	var total_lines int32
	for scanner.Scan() {
		total_lines++
	}
	return FileStats{Lines: total_lines}, nil

}

func SplitFile(filename string, number_of_files int) error {
	// What to do with empty lines in file? leave as is - the concat should reproduce the original file, tool is not semantic based!
	// get number of lines from file to split evenly amoungst new files

	total_lines, err := GetFileStats(filename)
	if err != nil {
		return errors.New("could not get number of lines from file for splitting")
	}
	if total_lines.Lines < int32(number_of_files) {
		return errors.New("you requested more files that total number of lines")
	}

	split_file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer split_file.Close()

	scanner := bufio.NewScanner(split_file)
	// scan file and generate subfiles

	// create new files
	new_files := []*os.File{}
	for i := 1; i <= number_of_files; i++ {
		new_file_name := fmt.Sprintf("%s_%d", filename, i)
		new_file_name = path.Base(new_file_name)
		new_file, err := os.Create(new_file_name)
		if err != nil {
			return fmt.Errorf("could not create file: %s\n", new_file_name)
		}
		defer new_file.Close()
		new_files = append(new_files, new_file)

	}

	lines_per_file := total_lines.Lines / int32(number_of_files)
	for line_num := 0; scanner.Scan(); line_num++ {

		line := fmt.Sprintf("%s\n", scanner.Text())
		file_index := int32(line_num) / lines_per_file
		if file_index+1 > int32(number_of_files) {
			file_index = int32(number_of_files) - 1
		}
		new_files[file_index].WriteString(line)
	}

	return nil
}
