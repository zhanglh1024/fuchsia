package main

import (
	"archive/tar"
	//"bufio"
	"compress/gzip"
	"fmt"
	//"github.com/Sirupsen/logrus"
	"io"
	"os"
)

func main() {
	//err:=os.Mkdir("/tmp/test/",007716)
	//file, err := os.Create("/tmp/test/test.go")
	//if err != nil {
	//	fmt.Println(err)
	//	logrus.Errorf("create file error!!")
	//}
	//defer file.Close()
	//str := "hello My test File!!"
	//outPutWtiter := bufio.NewWriter(file)
	//outPutWtiter.WriteString(str)
	//outPutWtiter.Flush()
	fw, err := os.Create("/tmp/test.tar.gz")
	if err != nil {
		panic(err)
	}
	defer fw.Close()
	gw := gzip.NewWriter(fw)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	dir,err:=os.Open("/tmp/test/")
	if err!=nil{
		panic(nil)
	}
	defer dir.Close()
	fis,err:=dir.Readdir(0)
	if err!=nil{
		panic(err)
	}
	for _,fi:=range fis{
		if fi.IsDir(){
			continue
		}
		fmt.Println(fi.Name())
		fr,err:=os.Open(dir.Name()+"/"+fi.Name())
		if err!=nil{
			panic(err)
		}
		defer fr.Close()
		h:=new(tar.Header)
		h.Name=fi.Name()
		h.Size=fi.Size()
		h.Mode=int64(fi.Mode())
		h.ModTime=fi.ModTime()
		if err!=nil{
			panic(err)
		}
		_, err = io.Copy(tw, fr)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}
	fmt.Println("tar.gz ok")

}
