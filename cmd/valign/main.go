package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mdwhatcott/valign"
)

var Version = "dev"

type Config struct {
	Match  string
	Blocks bool
}

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	var config Config
	flags := flag.NewFlagSet(fmt.Sprintf("%s @ %s", filepath.Base(os.Args[0]), Version), flag.ExitOnError)
	flags.StringVar(&config.Match, "match", "", "The value to match when vertically aligning the input text.")
	flags.BoolVar(&config.Blocks, "blocks", false, ""+
		"When set, only consider the first non-flag command-line argument for matching, and "+
		"treat each block of contiguous matching lines as separate blocks.")
	flags.Usage = func() {
		_, _ = fmt.Fprintln(flags.Output(), ""+
			"This program accepts input text from stdin and command-line arguments to "+
			"vertically aligns matches to in the input text.")
		flags.PrintDefaults()
	}
	_ = flags.Parse(os.Args[1:])

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, strings.ReplaceAll(line, "\t", " "))
		_, _ = fmt.Fprintln(os.Stderr, line)
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal("read err:", err)
	}

	matches := flags.Args()

	if config.Blocks {
		blocks := valign.Blocks(matches[0], lines...)
		lines = nil
		for _, block := range blocks {
			lines = append(lines, valign.On(matches[0], block...)...)
		}
	} else {
		for _, match := range matches {
			lines = valign.On(match, lines...)
		}
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}
