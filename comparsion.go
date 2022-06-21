package main

import "fmt"

func main() {
	v1 := "CREATE DATABASE pgGolang;"
	v2 := "CREATE DATABASE pgGolang;"

	if v1 != v2 {
		fmt.Println("no comparsion")
	}
	fmt.Println("Fine, comparsion!!!")
}
