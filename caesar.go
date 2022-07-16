package main

import (
	"fmt"
)

func main1() {
	caesar()
}

func caesar() {
	var offset int32
	_, err := fmt.Scan(&offset)
	if err != nil {
		return
	}
	var s string
	var r []byte
	_, err = fmt.Scan(&s)
	for _, b := range s {
		r = append(r, byte(97+(b-97+offset)%26))
	}
	fmt.Println(string(r))
}
