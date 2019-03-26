package main

import (
	"fmt"
	"sort"
)

func alphsort(strinput string) string {
	/*
	   функция, сортирующая символы
	   в строке в лексикографическом порядке
	*/

	var result string
	symlist := make([]string, len(strinput))
	for index, symbol := range strinput {
		symlist[index] = string(symbol)
	}
	sort.Strings(symlist)
	for _, symbol := range symlist {
		result += symbol
	}
	return result
}

func main() {
	fmt.Scan()
	words := make([]string)
	fmt.Println(alphsort(resultu))
}
