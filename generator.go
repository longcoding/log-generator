package main

import (
	"bufio"
	"fmt"
	"os"
	"flag"
	"time"
	"strconv"
)

func openFile(timestamp string) os.File {
	file, err := os.OpenFile(
		timestamp + ".log",
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		os.FileMode(0644),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return *file
}

func generateTimeStamp() string {
	currentTime := time.Now()
	return strconv.Itoa(currentTime.Year()) + AppendCharacter(int(currentTime.Month())) + AppendCharacter(currentTime.Day())
}

func main() {

	timeStamp := generateTimeStamp()
	file := openFile(timeStamp)
	defer file.Close()
	w := bufio.NewWriter(&file)

	mode := flag.String("mode", "", "realtime or onetime")
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
	} else if *mode == "realtime" {
		var logCount int
		if *count < 10 {
			logCount = 10
		} else {
			logCount = *count
		}
		for {
			for i:=0; i < logCount; i++ {
				w.WriteString(GetLogLine() + "\n")
			}
			w.Flush()
			time.Sleep(1000 * time.Millisecond)

			if generateTimeStamp() != timeStamp {
				timeStamp = generateTimeStamp()
				file.Close()
				file = openFile(timeStamp)
				w = bufio.NewWriter(&file)
			}
		}
	} else {
		fmt.Println("mode is not defined")
	}

}