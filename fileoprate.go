package main

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"fmt"
	"github.com/Sirupsen/logrus"
	"io"
	"os"
)

func main() {
	file, err := os.Create("/tmp/test.go")
	if err != nil {
		fmt.Println(err)
		logrus.Errorf("create file error!!")
	}
	defer file.Close()
	str := "hello My test File!!"
	outPutWtiter := bufio.NewWriter(file)
	outPutWtiter.WriteString(str)
	outPutWtiter.Flush()
	fw, err := os.Create("/tmp/test.tar.gz")
	if err != nil {
		panic(err)
	}
	defer fw.Close()
	gw := gzip.NewWriter(fw)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	fr, err := os.Open("/tmp/test.go")
	if err != nil {
		panic(err)
	}
	defer fr.Close()
	h := new(tar.Header)
	h.Name = fr.Name()

	err = tw.WriteHeader(h)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(tw, fr)
	if err != nil {
		panic(err)
	}
}
