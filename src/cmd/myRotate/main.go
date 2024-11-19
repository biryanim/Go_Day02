package main

import (
	"flag"
	"go_day02/pkg/rotate"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

func main() {
	a := flag.String("a", "", "Directory for archive files")
	flag.Parse()
	if *a != "" {
		if err := os.MkdirAll(*a, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
	files := flag.Args()

	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			fi, err := os.Stat(file)
			if err != nil {
				log.Fatal(err)
			}
			name := strings.TrimSuffix(file, filepath.Ext(file))
			if *a != "" {
				name = *a + "/" + name
			}
			out, err := os.Create(name + "_" + strconv.FormatInt(fi.ModTime().Unix(), 10) + ".tar.gz")
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()
			rotate.CreateArchive(file, out)
		}(&wg)
	}
	wg.Wait()
}
