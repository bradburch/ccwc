package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var byteFlag, lineFlag, wordFlag, charFlag bool

	flag.BoolVar(&byteFlag, "c", false, "output total bytes in file")
	flag.BoolVar(&lineFlag, "l", false, "output total lines in file")
	flag.BoolVar(&wordFlag, "w", false, "output total words in file")
	flag.BoolVar(&charFlag, "m", false, "output total characters in file")
	flag.Parse()

	if flag.NFlag() == 0 {
		byteFlag = true
		lineFlag = true
		wordFlag = true
	}

	file_name := flag.Arg(0)
	file, err := os.ReadFile(file_name)
	check(err)

	var result strings.Builder

	if byteFlag {
		result.WriteString(numBytes(file) + "\t")
	}
	if lineFlag {
		result.WriteString(numLines(file) + "\t")
	}
	if wordFlag {
		result.WriteString(numWords(file) + "\t")
	}
	if charFlag {
		result.WriteString(numCharacters(file) + "\t")
	}

	fmt.Println(result.String(), file_name)

}

func numBytes(file []byte) string {
	return strconv.Itoa(len(file))
}

func numLines(file []byte) string {
	lineSep := []byte{'\n'}
	return strconv.Itoa(bytes.Count(file, lineSep))
}

func numWords(file []byte) string {
	return strconv.Itoa(len(bytes.Fields(file)))
}

func numCharacters(file []byte) string {
	return strconv.Itoa(len(bytes.Runes(file)))
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
