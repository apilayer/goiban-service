/*
The MIT License (MIT)

Copyright (c) 2014 Chris Grieger

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package goiban

import (
	"bufio"
	"os"
	"strconv"
)

func isValidChar(ch uint8) bool {
	return (ch > 64 && ch < 91)
}

func isValidNum(ch uint8) bool {
	return (ch > 47 && ch < 58)
}

func toNumericString(val string) string {
	numericVal := ""
	for _, ch := range val {
		// if it's neither a number nor a char
		// fail
		intCh := uint8(ch)
		if !isValidNum(intCh) &&
			!isValidChar(intCh) {
			return ""
		}
		if isValidChar(intCh) {
			numericVal += strconv.Itoa(int(ch) - 55)
		} else {
			numericVal += string(ch)
		}
	}

	return numericVal
}

// code taken from
// http://stackoverflow.com/a/18479916/1408463
// changed to take a channel instead of writing an array
func readLines(path string, out chan string) {
	file, err := os.Open(path)
	if err != nil {
		out <- ""
		return
	}

	//converter,_ := iconv.NewConverter("utf-8", "windows-1252")
	defer file.Close()
	//defer converter.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res := scanner.Text()
		out <- res
	}
	close(out)
}
