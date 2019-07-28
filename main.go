package main

import "fmt"

func main() {
	alphabet := "az"
	for i := 0; i < len(alphabet); i++ {
		fmt.Printf("%v: %d\n", alphabet[i], int(alphabet[i]))
	}

}
