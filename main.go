package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"asciiart/fs"
)

// main function reads a banner file, creates a map of ASCII art, validates user input,
// and prints the corresponding ASCII art to the output file.
func main() {
	// Check if color flag is not provided correctly, i.e. provided without equal sign.
	properColorFlag := regexp.MustCompile(`^-color(?:=(.+))?$`)
	properOutputFlag := regexp.MustCompile(`^-output(?:=(.+))?$`)
	args := os.Args
	for _, v := range args {
		if properColorFlag.MatchString(v) || v == "--color" {
			fmt.Print("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"\n")
			return
		} else if properOutputFlag.MatchString(v) || v == "--output" {
			fmt.Print("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard\n")
			return
		}
	}

	// Get the flag values for color, letters to colorize, input text and banner file name. Handle possible errors.
	options, err := fs.ParseOptions()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Get ANSI format string to colorize ASCII-art in the output file.
	colorCode, err := fs.SetColor(options.ColorFlag)
	check(err)

	// Read banner file
	if options.BannerFile == "" {
		options.BannerFile = "standard"
	}
	bannerFile, err := fs.ReadBannerFile("./banners/" + options.BannerFile + ".txt")
	check(err)

	// Create map of ASCII art
	ASCIIArtMap, err := fs.MapCreator(bannerFile)
	check(err)

	// Get ASCII art corresponding to input text
	asciiArt, err := fs.ArtRetriever(options.InputText, colorCode, options.ColorizeLetters, ASCIIArtMap)
	check(err)

	// Checking whether the specified output file is a text file.
	outputFile := ""
	if options.OutputFlag != "" {
		if strings.HasPrefix(options.OutputFlag, "./banners/") || strings.HasPrefix(options.OutputFlag, "banners/") {
			fmt.Println("error: cannot write to or modify files in the banners directory")
			return
		} else if strings.HasSuffix(options.OutputFlag, ".txt") {
			outputFile = options.OutputFlag
		} else {
			fmt.Println("error: the output file must be a text file <filename.txt>")
			return
		}
	} else {
		fmt.Print(asciiArt)
		return
	}

	// Write the ASCII art to a file
	err = os.WriteFile(outputFile, []byte(asciiArt), 0o644)
	check(err)
}

// Handle errors
func check(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", e)
		os.Exit(1)
	}
}
