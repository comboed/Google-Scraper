package main

import (
	"bufio"
	"os"
)

func openFile(filename string) (slice []string) {
	var file, _ = os.Open(filename)
	var scan *bufio.Scanner = bufio.NewScanner(file)
	for scan.Scan() {
		slice = append(slice, scan.Text())
	}
	file.Close()
	return slice
}