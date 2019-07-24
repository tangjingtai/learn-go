package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	sfile string // 源文件
	dpath string // 目标路径
	size  int    // 每个文件放置多少行数据
)

func init() {
	flag.StringVar(&sfile, "sfile", "", "指定待分拆的源文件路径")
	flag.StringVar(&dpath, "dpath", "", "指定分拆后保存文件的目录")
	flag.IntVar(&size, "size", 0, "指定每个文件放置多少行数据")
}

func main() {
	flag.Parse()

	s, err := os.Stat(sfile)
	if err != nil {
		log.Fatal(err.Error())
	}
	if os.IsNotExist(err) || s.IsDir() {
		log.Fatal("The specified file could not be found:", sfile)
	}
	s, err = os.Stat(dpath)
	if err != nil {
		log.Fatal(err.Error())
	}
	if os.IsNotExist(err) || !s.IsDir() {
		log.Fatal("The specified directory could not be found:", dpath)
	}
	fileName := filepath.Base(sfile)
	fileSuffix := path.Ext(fileName)
	fileNameWithOutSuffix := strings.TrimSuffix(fileName, fileSuffix) //获取文件名
	fileIndex := 0
	var f *os.File
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	lineCount := 0
	err = ReadLine(sfile, func(bytes []byte) {
		line := string(bytes)
		line = strings.TrimRight(strings.TrimRight(line, "\n"), "\r")
		//items := strings.Split(line, "\t")
		//if len(items) < 3 {
		//	return
		//}
		//if _, err := strconv.Atoi(items[0]); err != nil {
		//	return
		//}
		lineCount++
		if lineCount%size == 1 {
			fileIndex++
			destFileName := path.Join(dpath, fmt.Sprintf("%s%05d%s", fileNameWithOutSuffix, fileIndex, fileSuffix))
			if f != nil {
				f.Close()
				f2, _ := os.OpenFile(destFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0)
				*f = *f2
			} else {
				f, _ = os.OpenFile(destFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0)
			}
		}
		f.WriteString(line)
		if lineCount%size != 0 {
			f.WriteString("\n")
		}
	})
	if err != nil {
		log.Fatal("open file Error,", err)
	}
}

func ReadLine(filePth string, hookfn func([]byte)) error {
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()

	bfRd := bufio.NewReader(f)
	for {
		line, err := bfRd.ReadBytes('\n')
		hookfn(line)    //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

func ReadAllLine(filePth string) ([]string, error) {
	var lines []string
	processLine := func(line []byte) {
		str := strings.TrimRight(string(line), "\r")
		if str == "" {
			return
		}
		lines = append(lines, str)
	}
	err := ReadLine(filePth, processLine)
	return lines, err
}
