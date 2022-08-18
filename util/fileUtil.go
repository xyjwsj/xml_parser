package util

import (
	"bufio"
	"io"
	"os"
)

type ReadCallback func(err error, line string)

// ReadFileLine 逐行读取文件
func ReadFileLine(filePath string, callback ReadCallback) {
	fi, err := os.Open(filePath)
	if err != nil {
		callback(err, "")
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			callback(io.EOF, "")
			break
		}
		callback(nil, string(a))
	}
}
