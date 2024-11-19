package main

import (
	"flag"
	"fmt"
	"go_day02/pkg/find"
)

func main() {
	f := flag.Bool("f", false, "Print only files")
	sl := flag.Bool("sl", false, "Print only symlinks")
	d := flag.Bool("d", false, "Print only directories")
	ext := flag.String("ext", "", "File extension")
	var all bool
	flag.Parse()
	if !(*f) && !(*sl) && !(*d) {
		all = true
	}
	if *ext != "" && !*f {
		fmt.Println("File extension not specified (need to add -f)")
		return
	}
	opt := find.Opt{Files: *f, SymLinks: *sl, Dir: *d, Ext: *ext, All: all}
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("No directories found")
		return
	}
	find.Walk(args[0], &opt)
}
