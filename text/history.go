package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// cmdFreq returns the frequency of "go" subcommand usage in ZSH history
func cmdFreq(fileName string) (map[string]int, error) {
	freq := make(map[string]int)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		if strings.Contains(s.Text(), ";go ") {
			test := strings.Split(s.Text(), " ")
			freq[test[2]]++
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return freq, nil
}

func main() {
	freqs, err := cmdFreq("text/zsh_history")
	if err != nil {
		log.Fatal(err)
	}

	for cmd, count := range freqs {
		fmt.Printf("%s -> %d\n", cmd, count)
	}
}
