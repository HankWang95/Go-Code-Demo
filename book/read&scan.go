package book

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"strings"
)

// 读取文件的信息，并计数
func ReadFile1_3() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprint(os.Stderr, "dup2 :%v \n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] = 1
	}
}

func ReadFile1_4() {
	counts := make(map[string]int)
	var flag = false
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ReadFile1_4: %v \n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			if counts[line] > 1 {
				flag = true
			}
		}
		if flag == true {
			fmt.Printf("filename:%s\n", filename)
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("time:%d\tword:%s\n", n, line)
			}
		}
		counts = make(map[string]int)
	}

}
