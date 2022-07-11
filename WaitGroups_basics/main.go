package main

import (
	"fmt"
	"sync"
)

func printThis(this string, wg *sync.WaitGroup) {
	//With Done, after running the wg will we reduce its number by 1. This will remove the wait groups as we finish processing. 9 -> 8 -> 7 etc..
	defer wg.Done()
	fmt.Println(this)
}

func main() {

	//MAJOR POINTS:
	//Wait groups can allow code to be completed before executing further code.
	//Wait groups still wont assign order to a function process! So the order that the list returns in our for loop is still random
	//It is the easiest way to deal with concurrency, not necessarily the most efficient way though.

	//wait group variable
	//enter a number for the number of things you need to wait for 9
	var wg sync.WaitGroup

	//this is the list of things we need to wait for: 9 words
	words := []string{
		"alpha",
		"beta",
		"delta",
		"gama",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}
	//9 items set to wait group
	wg.Add(len(words))

	for i, x := range words {
		//must include pointers to the wg reference
		go printThis(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	//Wait blocks the path to completion until the wg.Add counter reaches 0.
	wg.Wait()
	wg.Add(1)
	printThis("Hello World Second", &wg)
}
