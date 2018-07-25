package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"io/ioutil"
)

func readFile(path string){
	file,err := os.Open(path)
	if err != nil{
		fmt.Printf("read file error\n")
		return
	}

	defer file.Close()

	inputReader := bufio.NewReader(file)
	for{
		str,err :=inputReader.ReadString('\n')

		if err == io.EOF{
			fmt.Printf("read eof")
		}

		if err != nil{
			return //err or EOF
		}

		fmt.Printf("the input is:%s",str)
	}
}

func main(){

	readFile("D:/Go/workspace/src/test_go/mytest/test.txt")

	buf, _ := ioutil.ReadFile("D:/Go/workspace/src/test_go/mytest/test.txt")
	fmt.Printf("use io.ioutil read file:%s",buf)

}