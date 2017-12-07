package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}

	bs := make([]byte, stat.Size())
	n, err := file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
	fmt.Println(n)

	// ioutil
	bs1, err := ioutil.ReadFile("main.go")
	if err != nil {
		return
	}
	str1 := string(bs1)
	fmt.Println(str1)

	//create file
	newFile, err := os.Create("new.txt")
	if err != nil {
		fmt.Println("error message")
	}
	defer newFile.Close()
	newFile.WriteString("test")

	// directory
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
