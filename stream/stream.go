package stream

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

const (
	input    = "./tmp/input.txt"
	input2   = "./tmp/input2.txt"
	dist     = "./tmp/stream.txt"
	gzipdist = "./tmp/stream.txt.gz"
)

func Stream() {
	writer, err := os.Create(dist)
	if err != nil {
		fmt.Println("open err ", err)
	}

	reader, err := os.Open(gzipdist)
	if err != nil {
		fmt.Println("read err ", err)
	}

	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		fmt.Println("read err ", err)
	}

	n, err := io.Copy(writer, gzipReader)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}

var text = []rune("hoge")

type Test struct{ n int }

func (t *Test) Read(p []byte) (n int, err error) {
	in := len(p)
	nw := 0
	for i := 0; i < in; i++ {
		cb := []byte(string(text[t.n%len(text)]))
		if nw+len(cb) > in {
			break
		}
		cbl := len(cb)
		copy(p[nw:nw+cbl], cb)
		nw += cbl
		t.n++
	}
	return nw, nil
}

func FileCreate() {
	f, err := os.Create(input2)
	if err != nil {
		fmt.Println("file open error.", err)
	}
	defer f.Close()

	io.Copy(f, &Test{})
}

func FileCopy(input string) {
	fr, err := os.Open(input)
	if err != nil {
		fmt.Println("file open error.", err)
	}
	defer fr.Close()

	fw, err := os.Create(dist)
	if err != nil {
		fmt.Println("file open error.", err)
	}
	defer fw.Close()
	io.Copy(fw, fr)
}

func FileMerge() {
	pr, pw := io.Pipe()
	go func() {
		inputs := []string{input, input2}
		for _, input := range inputs {
			func() {
				fr, err := os.Open(input)
				if err != nil {
					fmt.Println("file open error.", err)
				}
				defer fr.Close()

				io.Copy(pw, fr)
			}()
		}
		pw.Close()
	}()

	fw, err := os.Create(dist)
	if err != nil {
		fmt.Println("file open error.", err)
	}
	defer fw.Close()

	io.Copy(fw, pr)
}

func FileMerge2() {
	fw, err := os.Create(dist)
	if err != nil {
		fmt.Println("file open error.", err)
	}
	defer fw.Close()

	inputs := []string{input, input2}
	for _, input := range inputs {
		func() {
			fr, err := os.Open(input)
			if err != nil {
				fmt.Println("file open error.", err)
			}
			defer fr.Close()

			io.Copy(fw, fr)
		}()
	}
}
