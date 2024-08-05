package main

import (
	"fmt"
	"net/http"
	"os"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("D:\\M_GO\\GO_gin\\learn_gin_01\\test2\\content.txt")

	if err != nil {
		fmt.Printf("found err in os.Open : %v\n", err)
		return
	}

	defer file.Close()

	var bs = make([]byte, 1023, 1024)
	num, err := file.Read(bs)

	if err != nil {
		fmt.Println("read byte num = ", num)
		fmt.Printf("found err in os.Read : %v\n", err)
		return
	}

	fmt.Fprintln(w, string(bs))

}

func main() {

	http.HandleFunc("/hello", sayhello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http sever failed , err : %v\n", err)
		return
	}

}
