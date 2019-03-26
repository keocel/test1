package main

import (
	"fmt"
	"sort"
	"strings"
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
	var (
		wordsNumber int
		tempStr     string
		finalSlice  [][]string
		tempSlice   []string
	)
	//ввод данных
	fmt.Println("number of words")
	fmt.Scanln(&wordsNumber)
	var words []string
	for i := 0; i < wordsNumber; i++ {
		fmt.Scanln(&tempStr)
		words = append(words, strings.ToLower(tempStr))
	}
	//перебор слов, пока список полон
	for len(words) > 0 {
		//создается список анаграмм, и туда "выдавливается" первое слово в списке
		tempSlice = append(tempSlice, words[0])
		tempStr = alphsort(words[0])
		words = words[1:]
		//цикл поиска  анаграмм по сравнению наборов букв, при нахождении
		//она перемещается из исходного списка в список анаграмм
		for i := 0; i < len(words); {
			if tempStr == alphsort(words[i]) {
				tempSlice = append(tempSlice, words[i])
				words = append(words[:i], words[i+1:]...)
			} else {
				i++
			}
		}
		sort.Strings(tempSlice)
		//отсортированный список анаграмм в цикле сравнивается с уже полученными до этого списками анаграмм,
		//по первому слову и помещается на на свое место в списке(в алфавитном порядке)
		for index, item := range finalSlice {
			if tempSlice[0] < item[0] {
				finalSlice = append(append(finalSlice[:index+1], tempSlice), finalSlice[index+1:]...)
				goto exit
			}
		}
		finalSlice = append(finalSlice, tempSlice)
	exit:
		tempSlice = make([]string, 0)
	}
	//вывод
	for _, item := range finalSlice {
		fmt.Println(item)
	}
}
