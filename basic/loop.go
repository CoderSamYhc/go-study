package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""

	for ; n > 0; n /= 2 {
		lsp := n % 2
		result = strconv.Itoa(lsp) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() { // 每次读取一行
		fmt.Println(scanner.Text())
	}

}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13))

	printFile("abc.txt")
}
