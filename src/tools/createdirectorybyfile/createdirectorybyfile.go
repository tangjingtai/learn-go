package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
)

const (
	dirPath string = "D:\\数据处理\\13. 创建智慧作业的图片\\初中物理"
)

func main() {
	rd, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal("读取目录出错", err.Error())
	}
	index := 0
	for _, fi := range rd {
		if fi.IsDir() {
			continue
		}
		index++
		for {
			d := path.Join(dirPath, strconv.Itoa(index))
			_, err := os.Stat(d)

			if err != nil && os.IsNotExist(err) {
				os.MkdirAll(d, os.ModePerm)
				break
			}
			index++
		}
		fileName := fi.Name()
		os.Rename(path.Join(dirPath, fileName), path.Join(dirPath, strconv.Itoa(index), fileName))
	}
}
