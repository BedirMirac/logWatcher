package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	logFiles := []string{"/var/log/pacman.log", "text.txt"} // update this slices for yourself // doesn't matter how many files you wanna add

	for _, file := range logFiles {
		wg.Add(1)
		go watcher(file)
	}

	wg.Wait()
}

func Read(file *os.File, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				time.Sleep(500 * time.Millisecond) // if there is no line added after this program started, wait 500 ms to check again
				continue
			}
			break

		}
		processLine(line) // Sending line to function to see if the word we are looking for in the line or not
	}
}

func processLine(line string) { // where we print the messages as soon as we find a word that we chose
	line = strings.TrimSpace(line)
	lowerLine := strings.ToLower(line)

	if strings.Contains(lowerLine, "warning") {
		fmt.Printf("WARNING: %s\n", line)
	} else if strings.Contains(lowerLine, "error") {
		fmt.Printf("ERROR: %s\n", line)
	}
}

func watcher(file string) { // opens log files heading the end of the file and do progress
	log, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer log.Close()
	log.Seek(0, io.SeekEnd)
	Read(log, &wg)
}
