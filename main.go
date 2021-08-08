package main

import (
	"log"
	"sort"
)

func main() {
	var mySlice []int

	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 2)

	sort.Ints(mySlice)

	log.Println(mySlice)

	// another way
	numbers := []int{1,2,3,4,5,6,7}

	log.Println(numbers)

	// only get the range
	log.Println(numbers[3:5])
}