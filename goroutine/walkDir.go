package goroutine

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"
	"time"
)

func walkDir(dir string, n *sync.WaitGroup, fileSize chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := path.Join(dir, entry.Name())
			walkDir(subdir, n, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

// 限制并发数量 计数信号量
var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err :%v", err)
		return nil
	}
	return entries

}

var verbose = flag.Bool("v", false, "show verbose progress.")

func DoWalkDir() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(50 * time.Millisecond)
	}
	// 并行遍历文件树

	for _, root := range roots {
		fmt.Println(root, "-----------------")
		var fileSize = make(chan int64)
		var n sync.WaitGroup
		n.Add(1)
		go walkDir(root, &n, fileSize)
		go func() {
			n.Wait()
			close(fileSize)
		}()
		var nfiles, nbyte int64
		func(root string) {

		loop:
			for {
				select {
				case size, ok := <-fileSize:
					if !ok {
						break loop
					}
					nfiles++
					nbyte += size
				case <-tick:
					printDiskUsage(root, nfiles, nbyte)
				}

			}
			log.Print("break loop")
			printDiskUsage(root, nfiles, nbyte)

		}(root)
	}
}
func printDiskUsage(root string, nfiles, nbyte int64) {
	_, err := fmt.Printf("root:%s, %d files \t %.1f GB\n", root, nfiles, float64(nbyte)/1e9)
	if err != nil {
		log.Fatal(err)
	}
}
