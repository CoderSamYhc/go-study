package main

import (
	"fmt"
	"io/ioutil"
)


func grade(score int) string {
	s := ""
	switch {
	case 0 > score || score > 100:
		panic(fmt.Sprintf("Error score.. %d", score))
	case score < 60:
		s = "D"
	case score < 80:
		s = "C"
	case score < 90:
		s = "B"
	case score <= 100:
		s = "A"
	default:
		panic(fmt.Sprintf("Error score %d", score))
	}
	return s
}

func getContents() {
	const filename = "abc.txt"
	if c, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("contents: %s\n", c)
	}
}

func main() {

	getContents()
	s := grade(10)
	fmt.Println(s)
}
