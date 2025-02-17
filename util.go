package main

import (
	"strconv"
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

func formatNumber(number int64) string {
    var in string = strconv.FormatInt(number, 10)
	var out []byte = make([]byte, len(in) + (len(in) - 1) / 3)
	for i, j, k := len(in) - 1, len(out) - 1, 0; ; i, j = i - 1, j - 1 {
		out[j] = in[i]
		if (i == 0 ){
			return string(out)
		}
        if k++; (k == 3) {
			j, k = j - 1, 0
			out[j] = ','
        }
    }
}
