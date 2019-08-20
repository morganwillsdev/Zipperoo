package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"errors"
)

var validPath = "^(\\.{1,2}/)*([\\w\\-]+/)*([\\w\\-]+)(\\.[\\w\\-\\.]+)?$"

func parseArgs(args []string) error {
	command := args[0]

	if command == "-zip" {
		var filesToZip []string
		var outputFile string
		var fileFormat string

		match, _ := regexp.MatchString(validPath, args[1])
		if match {
			outputFile = args[1]
			fileFormat = strings.Join(strings.Split(outputFile, ".")[1:], ".")
		} else {
			return errors.New(fmt.Sprintf("invalid file name: %s", args[1]))
		}


		// checks if arguments are valid file or directory names and adds them to the filesToZip slice
		for _, path := range args[2:] {
			match, _ := regexp.MatchString(validPath, path)

			if match {
					filesToZip = append(filesToZip, path)	
			} else {
				return errors.New(fmt.Sprintf("invalid file name: %s", path))
			}
		}

		// check which format files should be compressed to : change to include more formats
		if len(filesToZip) > 0 {
			if fileFormat == "zip" {
				err := ZipFiles(outputFile, filesToZip)
				if err != nil {
					return err
				}
			} else {
				return errors.New(fmt.Sprintf("%s file format currently unsupported\n", fileFormat))
			}
		} else {
			return errors.New("no files specified")
		}
		return nil

	} else if command == "-unzip" {
		var outputPath string
		var fileToUnzip string
		var fileFormat string

		match, _ := regexp.MatchString(validPath, args[1])

		if match {
			fileToUnzip = args[1]
			fileFormat = strings.Join(strings.Split(fileToUnzip, ".")[1:], ".")
		} else {
			return errors.New(fmt.Sprintf("invalid file name: %s", args[1]))
		}

		if len(args) > 2 {
			match, _ := regexp.MatchString(validPath, args[2])

			if match {
				outputPath = args[2]
			} else {
				return errors.New(fmt.Sprintf("invalid file name: %s", args[2]))
			}
		} else {
			outputPath = "zipperoo-default"
		}

		if fileFormat == "zip" {		// change to include more formats
			err := Unzip(fileToUnzip, outputPath)
			if err != nil {
				return err
			}
		} else {
			return errors.New(fmt.Sprintf("%s file format currently unsupported\n", fileFormat))
		}
		return nil
	} else {
		return errors.New(fmt.Sprintf("unknown command: %s\ntry running the -help command\n", command))
	}
}


func main() {
	args := os.Args[1:]

	if len(args) < 1 || (len(args) < 2 && args[0] != "-help") {
		fmt.Printf("not enough arguments\n")
		fmt.Printf("try running the -help command\n")
	} else if args[0] == "-help" {
		fmt.Printf("help\n")
	} else {
		err := parseArgs(args)

		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}
}