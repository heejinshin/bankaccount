package main 

import (
	"./mydict"
	"fmt"
)


func main() {
	dictionary := mydict.Dictionary{"first":"First word"}
	definition, err := dictionary.Serach("second")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}