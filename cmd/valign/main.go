package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mdwhatcott/valign/v2"
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
		"When set, treat each block of contiguous matching lines as separate blocks.")
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
		lines = append(lines, strings.ReplaceAll(scanner.Text(), "\t", " "))
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal("read err:", err)
	}

	if config.Blocks {
		lines = valign.Blocks(config.Match, lines...)
	} else {
		lines = valign.On(config.Match, lines...)
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}
