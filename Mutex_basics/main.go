package main

import (
	"fmt"
	"sync"
)

var msg string        //hold message
var wg sync.WaitGroup //waitgroup

func updateMessage(s string /*, m *sync.Mutex */) {
	defer wg.Done() //after process is done decrement waitgroup

	//REMOVING THE RACE CONDITION!
	//THIS IS KNOWN AS A THREAD SAFE OPERATION!
	//m.Lock() //locks anything within so only the routine running it can make changes!
	msg = s
	//m.Unlock() //unlocks when the function is done needing it
}

func main() {
	msg = "Hello, World!" //initial value

	//var mutex sync.Mutex //mutex

	wg.Add(2) //wait for 2 decrements
	go updateMessage("Hello, Universe!" /*&mutex*/)
	go updateMessage("Hello, Cosmos!" /*&mutex*/)
	wg.Wait() //waits here

	fmt.Println(msg) //print message
	//Issue!!!! Hello, Universe is finishing last and is getting read back to the user! But Hello, Cosmos is run second!
	//go run -race .     //This run in the terminal tells us the code has a race condition, the two routines are racing to finish.
}
