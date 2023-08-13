package mycat

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Read(w io.Writer, fileNames []string, n bool) {
	line := 0
	for _, fileName := range fileNames {
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		defer f.Close()

		ReadFromReader(w, f, n, &line)
	}
}

func ReadFromReader(writer io.Writer, reader io.Reader, n bool, line *int) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if n {
			*line++
			fmt.Fprintf(writer, "%d: %s\n", *line, scanner.Text())
		} else {
			fmt.Fprintln(writer, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(writer, "reading input:", err)
	}
}
