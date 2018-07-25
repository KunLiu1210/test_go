package main

import (
	"time"
	"log"
	"fmt"
	"os"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}


func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}



func main()  {
	//fmt.Println(triple(4)) // "12"
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return
		}
		defer f.Close() // NOTE: risky; could run out of file descriptors
		// ...process f...
	}

	for _, filename := range filenames {
		if err := doFile(filename); err != nil {
			return
		}
	}
}

func doFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// ...process f...

	return err
}

