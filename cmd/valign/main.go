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
	for _, match := range flag.Args() {
		lines = valign.On(match, lines...)
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}
