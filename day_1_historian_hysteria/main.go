package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func read() ([]int, []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var list1 []int
	var list2 []int

	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		
		for i, number := range numbers {
			num, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
			if i+2 % 2 == 0 {
				list1 = append(list1, num)
			} else {
				list2 = append(list2, num)
			}
		}
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return list1, list2
}

func sortList(list []int) []int {
	sort.Ints(list)

	return list
}

func compareList(list1 []int, list2 []int) []int {
	var differences []int
	if len(list1) == len(list2) {
		for i := range list1 {
			num1 := list1[i]
			num2 := list2[i]
			diff := num1 - num2
			if diff < 0 {
				diff = diff *-1
			}
			differences = append(differences, diff)
		}
	}

	return differences
}

func part1() {
	// first read the lists from the text file
	list1, list2 := read()
	// sort the lists
	list1 = sortList(list1)
	list2 = sortList(list2)
	// compare the lists
	differences := compareList(list1, list2)
	// total
	//fmt.Println(differences)
	var total = 0
	for _, val :=  range differences {
		total += val
	}
	// output answer\
	fmt.Println(total)
}

func countItems(list1 []int, list2 []int) map[int]int {
	myMap := make(map[int]int)
	var counter = 0
	for _, val1 := range list1 {
		for _, val2 := range list2 {
			if val1 == val2 {
				//fmt.Println("matched")
				counter = counter +1
				//fmt.Println(counter)
			}
		}
		myMap[val1] = counter
		counter = 0
	}
	return myMap
}

func part2() {
	// first read the lists from the text file	
	list1, list2 := read()
	// compare the lists to get the number in list 2
	myMap := countItems(list1, list2)
	//fmt.Println(myMap)
	// total

	var total = 0
	for key, value := range myMap {
		val := key * value
		total = total + val
	}
	// print
	fmt.Println(total)
}

func main() {
	part1()
	part2()
}