package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var (
		ext = flag.String("e", "", "target file extension")
		abs = flag.Bool("a", false, "show absolute path")
	)

	flag.Parse()
	args := flag.Args()

	if flag.NArg() == 0 {
		fmt.Println("Please pass a directory path.")
		os.Exit(1)
	}

	passstore := ""
	if strings.HasSuffix(args[0], "/") {
		passstore = args[0]
	} else {
		passstore = args[0] + "/"
	}

	err := filepath.Walk(passstore, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if *ext == "" {
			if *abs == false {
				path = strings.Replace(path, passstore, "", -1)
			}
			fmt.Printf("%s\n", path)
		} else {
			if strings.HasSuffix(path, *ext) {
				if *abs == false {
					path = strings.Replace(path, passstore, "", -1)
				}
				fmt.Printf("%s\n", path)
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}
