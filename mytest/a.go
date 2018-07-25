package main

import (
	"fmt"
)

func Abc()  {
	fmt.Println("Foun new Peer Ip: ")
}


func main() {
	var rmdirs []func()
	tempDirs := []string{"1","2","3"}
	for _, dir := range tempDirs{
		//os.MkdirAll(dir, 0755)
		dir := dir
		rmdirs = append(rmdirs, func() {
			//os.RemoveAll(dir) // NOTE: incorrect!
			fmt.Println(dir)
			//fmt.Println(d)
		})
	}

	for _, rmdir := range rmdirs {
		rmdir()
	}
}

