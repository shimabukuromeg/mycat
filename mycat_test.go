package mycat

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestReadFromReader(t *testing.T) {
	input := "Hello\nWorld"
	r := strings.NewReader(input)
	var buf bytes.Buffer
	line := 0

	ReadFromReader(&buf, r, true, &line)

	expected := "1: Hello\n2: World\n"
	if buf.String() != expected {
		t.Errorf("Expected output %q, but got %q", expected, buf.String())
	}
}

func TestRead(t *testing.T) {
	// 2つのテスト用ファイルを作成
	contents := []string{
		"Hello\nWorld",
		"Another\nTest",
	}
	var fileNames []string

	for _, content := range contents {
		tmpfile, err := os.CreateTemp("", "example.*.txt")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpfile.Name())

		if _, err := tmpfile.Write([]byte(content)); err != nil {
			tmpfile.Close()
			t.Fatal(err)
		}
		tmpfile.Close()

		fileNames = append(fileNames, tmpfile.Name())
	}

	// Read 関数を実行して出力をキャッチする
	var buf bytes.Buffer
	Read(&buf, fileNames, true)

	expected := "1: Hello\n2: World\n3: Another\n4: Test\n"
	if buf.String() != expected {
		t.Errorf("Expected output %q, but got %q", expected, buf.String())
	}
}
