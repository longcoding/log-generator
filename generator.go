package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"application.log",
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		os.FileMode(0644),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	fmt.Println(GetLogLine())

	w.WriteString("Hello, world\n")
	w.WriteString("Hello, world\n")
	w.Flush()
}