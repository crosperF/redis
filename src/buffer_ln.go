package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func buffer_test() {
	input := "$4\r\njohn\r\n"
	reader := bufio.NewReader(strings.NewReader(input))

	char, _ := reader.ReadByte()
	if char != '$' {
		fmt.Println("Invalid type, expecting bulk strings only")
		os.Exit(1)
	}

	size, _ := reader.ReadByte()

	str_size, _ := strconv.ParseInt(string(size), 10, 64)

	reader.ReadByte()
	reader.ReadByte()

	buffer := make([]byte, str_size)
	reader.Read(buffer)

	fmt.Println(string(buffer))
}
