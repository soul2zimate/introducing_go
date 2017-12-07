package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func getHash(filename string) (uint32, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	h := crc32.NewIEEE()

	_, err = io.Copy(h, f)

	if err != nil {
		return 0, err
	}

	return h.Sum32(), nil
}

func main() {
	// hash
	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println(v)

	h1, err := getHash("1.txt")
	if err != nil {
		return
	}

	h2, err := getHash("2.txt")
	if err != nil {
		return
	}

	fmt.Println(h1, h2, h1 == h2)

	// cryptography
	h3 := sha1.New()
	h3.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}
