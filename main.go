package main

import "fmt"

func main() {

	//example := make([]string, 2, 10)
	example := []string{}

	example = append(example, "Nelly")
	example = append(example, "Stanford")
	example = append(example, "Fernandes")
	example = append(example, "Martins")

	fmt.Println(example)
	fmt.Println(len(example))
}
