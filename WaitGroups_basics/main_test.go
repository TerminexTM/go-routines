package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printThis(t *testing.T) {
	//when you run the program it is printing to a standard out: stdout
	//we are saving this just to reset at the end
	stdOut := os.Stdout

	//we are now creating our own unique standard out
	r, w, _ := os.Pipe()
	//change the standard out of the library to be w
	os.Stdout = w

	//create a wait group variable
	var wg sync.WaitGroup
	//add a number to wait group
	wg.Add(1)

	//write the test
	go printThis("epsilon", &wg)

	//wait until program finishes running
	wg.Wait()

	//close the pipe
	_ = w.Close()

	//save what was written as a variable
	result, _ := io.ReadAll(r)

	//cast the result into a string format
	output := string(result)

	//set the stdout to what it was before the test ran
	os.Stdout = stdOut

	if !strings.Contains(output, "epsilon") {
		t.Errorf("Expected to find epsilon, but found %s", output)
	}
}
