package find

import (
	"fmt"
	"os"
	"path/filepath"
)

type Opt struct {
	Files    bool
	SymLinks bool
	Dir      bool
	All      bool
	Ext      string
}

func Walk(path string, opt *Opt) {
	var walk func(string)
	walk = func(path string) {
		file, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("Unable to get list of files", err)
			return
		}
		for _, f := range file {
			filename := filepath.Join(path, f.Name())
			isDir := f.IsDir()
			isSymLink := f.Type()&os.ModeSymlink != 0
			if (opt.Files && !isDir && !isSymLink) || (opt.Dir && isDir) || (opt.SymLinks && isSymLink) || opt.All {
				if opt.Files && opt.Ext != "" && !isDir {
					if filepath.Ext(filename) != "."+opt.Ext {
						continue
					}
				}
				if isSymLink {
					dest, err := os.Readlink(filename)
					if err != nil {
						return
					}
					if _, err := os.Stat(dest); err == nil {
						fmt.Printf("%s -> %s\n", filename, dest)
					} else if os.IsNotExist(err) {
						fmt.Printf("%s -> [broken]\n", filename)
					} else {
						continue
					}
				} else {
					fmt.Println(filename)
				}
			}
			if f.IsDir() {
				walk(filename)
			}
		}

	}
	walk(path)
}
