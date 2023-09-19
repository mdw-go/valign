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

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	flags := flag.NewFlagSet(fmt.Sprintf("%s @ %s", filepath.Base(os.Args[0]), Version), flag.ExitOnError)
	flags.Usage = func() {
		_, _ = fmt.Fprintln(flags.Output(), ""+
			"This program accepts input text from stdin and non-flag command-line arguments and "+
			"vertically aligns matches to command-line flags in the input text.")
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
	for _, match := range flags.Args() {
		lines = valign.On(match, lines...)
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}
