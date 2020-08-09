package stream

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
	TB
)

const (
	tmpDir       = "./tmp"
	dataFile     = "tmp/data.txt"
	dataFileGzip = "tmp/data.txt.gz"
	cmdInput     = "tmp/cmdInput1.txt.gz"
	cmdInput2    = "tmp/cmdInput2.txt.gz"
)

//func Initialize() {
//	f, err := os.OpenFile(dataFile, os.O_CREATE|os.O_RDWR| os.O_APPEND, 0666)
//	if err != nil {
//
//	}
//	for i := 0; i < 100000; i++ {
//		str := fmt.Sprintf("select * from tables where hoge = \"%d\"\n", i)
//		f.WriteString(str)
//	}
//}

type Hoge struct {
	writer io.Writer
}

func (h *Hoge) Write(p []byte) (n int, err error) {
	gw := gzip.NewWriter(h.writer)
	defer gw.Close()

	n, err = gw.Write(p)
	if err != nil {
		return n, err
	}
	return n, nil
}

func createGzip(dst, src string) {
	h := new(Hoge)
	writer, err := os.Create(dst)
	if err != nil {
		fmt.Println("open err ", err)
	}

	h.writer = writer

	reader, err := os.Open(src)
	if err != nil {
		fmt.Println("read err ", err)
	}

	_, err = io.Copy(h, reader)
	if err != nil {
		fmt.Println("copy", err)
	}
}

func GzipDivider() {
	fmt.Println(KB)
	// createGzip(dataFileGzip, dataFile)
	devide(dataFileGzip)
}

func fileNameGenerator(basename string) func() string {
	ext := filepath.Ext(basename)[1:]
	filenameWithoutExt := basename[:len(basename)-len(ext)-1]
	count := 0

	return func() string {
		count++
		return fmt.Sprintf(
			"%s.%d.%s",
			filenameWithoutExt,
			count,
			ext,
		)

	}
}

func devide(gzipFilePath string) {
	basename := filepath.Base(gzipFilePath)
	fGenerator := fileNameGenerator(basename)

	f, err := os.Open(gzipFilePath)
	if err != nil {
		os.Exit(0)
	}

	fileInfo, err := f.Stat()
	if err != nil {
		os.Exit(0)
	}

	size := int(fileInfo.Size())
	fmt.Println(size)

	restrictedByte := 50 * KB

	for {
		var tmp int
		if size > restrictedByte {
			tmp = restrictedByte
			size -= restrictedByte
		} else {
			tmp = size
		}

		b := make([]byte, tmp)
		fmt.Println(tmp)

		n, err := f.Read(b)
		if err != nil {
			fmt.Println(err)
		}
		if n == 0 {
			break
		}

		if err := ioutil.WriteFile(filepath.Join(tmpDir, fGenerator()), b, 0666); err != nil {
			fmt.Println(err)
		}
	}

}

func CommandStream() {
	cmd := exec.Command("cat")

	cmdWriter, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}

	filenames := []string{
		"data.txt.1.gz",
		//"data.txt.2.gz",
		//"data.txt.3.gz",
		//"data.txt.4.gz",
		//"data.txt.5.gz",
		//"data.txt.6.gz",
	}

	go func() {
		for _, filename := range filenames {
			func() {
				readPath := filepath.Join(tmpDir, filename)

				reader, err := os.Open(readPath)
				if err != nil {
					fmt.Println("read err ", err)
				}
				defer reader.Close()

				gzipReader, err  := gzip.NewReader(reader)
				if err != nil {
					fmt.Println("read err ", err)
				}
				defer gzipReader.Close()

				_, err = io.Copy(cmdWriter, gzipReader)
				if err != nil {
					fmt.Println(err)
				}
			}()
		}
		cmdWriter.Close()
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}

func MergeFiles() {

	filenames := []string{
		"data.txt.1.gz",
		"data.txt.2.gz",
		"data.txt.3.gz",
		"data.txt.4.gz",
		"data.txt.5.gz",
		"data.txt.6.gz",
	}

	var merged []byte

	for _, filename := range filenames {
		readPath := filepath.Join(tmpDir, filename)
		b, err := ioutil.ReadFile(readPath)
		if err != nil {
			fmt.Println(err)
		}
		merged = append(merged, b...)
	}

	ioutil.WriteFile("./tmp/verification.txt.gz", merged, 0666)

}
