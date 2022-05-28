package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {

	if len(os.Args) < 1 || reflect.TypeOf(os.Args[1]).Kind() != reflect.String {
		log.Fatal("This is a useless check lol")
	}
	fileName := os.Args[1]
	filePtr, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	// what a weird pattern
	byteArr := make([]byte, 9999)

	filePtr.Read(byteArr)

	fmt.Println(string(byteArr))

}
