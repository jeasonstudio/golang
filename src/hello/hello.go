package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {
	rc, err := zip.OpenReader("a.zip")

	if err != nil {
		defer rc.Close()
	}
	for _, _file := range rc.File {
		fmt.Println(_file.Name)

		f, _ := _file.Open()

		desfile, err1 := os.OpenFile(_file.Name, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err1 == nil {
			fmt.Println("OK")
			fmt.Println(int64(_file.UncompressedSize64))
			io.CopyN(desfile, f, int64(_file.UncompressedSize64))
			desfile.Close()
		} else {
			defer desfile.Close()
		}
	}
}
