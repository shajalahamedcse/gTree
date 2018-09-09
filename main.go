package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	args := []string{"."}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	//fmt.Println(args)

	for _, arg := range args {
		err := tree(arg)
		if err != nil {
			log.Printf("tree %s: %vn", arg, err)
		}
	}
}

func tree(root string) error {
	//fmt.Println(path)
	err := filepath.Walk(root, func(path string, fileinfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fileinfo.Name()[0] == '.' {
			return filepath.SkipDir
		}

		fmt.Println(path)
		return nil
	})
	return err
}
