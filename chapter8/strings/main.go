package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("test", "te"))
	fmt.Println(strings.HasSuffix("test", "st"))
	fmt.Println(strings.Index("test", "e"))
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	fmt.Println(strings.Repeat("a", 5))
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))
	fmt.Println(strings.Split("a-b-c-d", "-"))
	fmt.Println(strings.ToLower("TEST"))
	fmt.Println(strings.ToUpper("test"))
	arr := []byte("monday")
	fmt.Println(arr)
	str1 := string([]byte{'g', 'o', 'o', 'd'})
	fmt.Println(str1)
	str2 := string(arr)
	fmt.Println(str2)
}
