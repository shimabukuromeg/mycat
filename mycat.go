package mycat

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

func Read(fileNames []string, n bool) {
	line := 0
	for _, fileName := range fileNames {
		path := path.Join(".", fileName)
		f, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
			return
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if n {
				line++
				fmt.Fprintf(os.Stdout, "%d: %s\n", line, scanner.Text())
			} else {
				fmt.Println(scanner.Text())
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		// ファイルを閉じる
		defer f.Close()
	}
}
