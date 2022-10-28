package main

import "fmt"

func main() {
	name := "Selamat Malam"

	for _, s := range name {
		fmt.Println(string(s))
	}

	m := make(map[string]int)

	m[""] = 1
	m["a"] = 4
	m["e"] = 1
	m["l"] = 2
	m["m"] = 3
	m["s"] = 1
	m["t"] = 1

	fmt.Println(m)

}
