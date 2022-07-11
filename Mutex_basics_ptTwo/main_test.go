package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {

	//creates a unique standard out from pipe
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	//sets us up to grab whatever is being sent to be written

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r) //read what was written
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "$34320") {
		t.Error("Wrong balance returned")
	}
}
