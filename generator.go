package main

import (
	"bufio"
	"fmt"
	"os"
	"flag"
	"time"
	"strconv"
)

func main() {

	currentTime := time.Now()
	fileName := strconv.Itoa(currentTime.Year()) + AppendCharacter(int(currentTime.Month())) + AppendCharacter(currentTime.Day());

	file, err := os.OpenFile(
		fileName + ".log",
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		os.FileMode(0644),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	mode := flag.String("mode", "realtime", "realtime or onetime")
	count := flag.Int("count", 0, "log line count")
	flag.Parse()

	if *mode == "onetime" {
		if *count == 0 {
			fmt.Println("count is not defined")
			os.Exit(0)
		}
		for i:=0; i < *count; i++ {
			w.WriteString(GetLogLine() + "\n")
		}
		w.Flush()
	} else {
		for {
			for i:=0; i < 10; i++ {
				w.WriteString(GetLogLine() + "\n")
			}
			w.Flush()
			time.Sleep(1000 * time.Millisecond)
		}
	}

}