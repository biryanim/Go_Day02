package main

import (
	"bufio"
	"flag"
	"fmt"
	"go_day02/pkg/wc"
	"sync"
)

func main() {
	l := flag.Bool("l", false, "Count lines")
	m := flag.Bool("m", false, "Count characters")
	w := flag.Bool("w", false, "Count words")
	flag.Parse()
	if *l && *m || *l && *w || *m && *w {
		fmt.Println("You need to provide -l or -m or -w")
		return
	}
	var strategy func(data []byte, atEOF bool) (advance int, token []byte, err error)
	switch {
	case *l:
		strategy = bufio.ScanLines
	case *m:
		strategy = bufio.ScanRunes
	case *w:
		strategy = bufio.ScanWords
	default:
		strategy = bufio.ScanWords
	}
	filenames := flag.Args()
	var wg sync.WaitGroup
	for _, filename := range filenames {
		wg.Add(1)
		go func() {
			count, err := wc.Counter(filename, &wg, strategy)

			if err != nil {
				fmt.Println(err)

			} else {
				fmt.Printf("%d\t%s\n", count, filename)
			}
		}()
	}
	wg.Wait()
}
