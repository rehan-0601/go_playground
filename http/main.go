package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//bs := []byte{}
	// empty byte slice of size 99999

	// resp.Body implements the Reader interface
	// Reader interface contains read function
	// the library has already implemented the interface.
	// in order to use the read function :
	// read function takes in a byte slice.

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// source of data --> Reader interface --> []byte
	// we get access to byte in our program
	// []byte ---> Writer interface ---> some form of output
	// []byte --> Writer Interface --> terminal/outgoing http

	// io.Copy==>
	// std lib - os.StdOut implements the writer interface
	// resp.Body implements the reader interface

	io.Copy(os.Stdout, resp.Body)
}
