package main

import (
	"fmt"
	"numbercypher/cypher"
)

func main() {
	var input string
	//fmt.Scanln(&input)
	input = "RRL=R"

	result, err := cypher.CreateNumberCypher(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
