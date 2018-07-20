package book

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 判断输入是否重复
func Dedup() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

func WordFrequency() {
	var count = make(map[string]int)

	for _, i := range os.Args[1:] {
		data, err := ioutil.ReadFile(i)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Word Frequency:%s", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			wordList := strings.Split(line, " ")
			for _, word := range wordList {
				count[word]++
			}

		}
	}
	for k, v := range count {
		fmt.Printf("word:%s count:%d\n", k, v)
	}
}


