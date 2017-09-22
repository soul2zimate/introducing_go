package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	x2 := make(map[int]int)
	x2[1] = 12
	fmt.Println(x2[1])

	// delete
	delete(x2, 1)
	fmt.Println(x2)

	//example
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])

	aName, aOk := elements["C"]
	fmt.Println(aName)
	fmt.Println(aOk)

	anotherName := elements["Be"]
	fmt.Println(anotherName)

	if name, ok := elements["He"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("find nothing")
	}
}
