package main

import "testing"

//you need to add the -race flag to your tests to make sure it tests for those conditions in the terminal!
func Test_updateMessage(t *testing.T) {
	msg = "hello world"

	wg.Add(2)

	go updateMessage("Goodbye, cruel world!")
	go updateMessage("Goodbye")

	wg.Wait()

	if msg != "Goodbye, cruel world!" {
		t.Error("Incorrect value in the msg variable")
	}
}
