package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mdwhatcott/valign"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	var match string
	flag.StringVar(&match, "match", "", "What to align on.")
	flag.Parse()

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal("read err:", err)
	}

	for _, line := range valign.On(match, lines...) {
		fmt.Println(line)
	}
}
