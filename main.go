package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

func main() {

	dirname := os.Args[1]

	if len(os.Args) == 0 {
		fmt.Println("Arg cannot be empty")
		os.Exit(1)
	}

	setOfFile := ReadDir(dirname)
	sorted := Sort(setOfFile)

	for _, val := range sorted {
		fmt.Printf("%s -- %v\n", val.Name, val.Size)
	}
}

type File struct {
	Name string
	Size int64
}

func ReadDir(dirname string) []File {

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	var filesContent []File

	for _, file := range files {
		var fileNode File
		if !file.IsDir() {
			fileNode.Name = file.Name()
			fileNode.Size = file.Size()
		}
		filesContent = append(filesContent, fileNode)
	}

	return filesContent

}

func Sort(setOfFile []File) []File {

	sort.Slice(setOfFile, func(a, b int) bool {
		return setOfFile[a].Size > setOfFile[b].Size
	})

	return setOfFile
}
