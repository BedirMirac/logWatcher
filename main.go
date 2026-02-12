package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const logPath = "/var/log/pacman.log" // the path of log file that you wanna watch

func main() {
	file := OpenFile()
	defer file.Close()
	file.Seek(0, io.SeekEnd)
	ReadPacman(file)

}

func OpenFile() *os.File {
	file, err := os.Open(logPath)
	if err != nil {
		panic(err)
	}
	return file
}

func ReadPacman(file *os.File) {
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				time.Sleep(500 * time.Millisecond)
				continue
			}
			break

		}
		processLine(line)
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
