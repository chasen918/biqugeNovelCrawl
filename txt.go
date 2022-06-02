package main

import (
	"bufio"
	"fmt"
	"os"
)

func MkdirForTxt(dirname string) error {
	dir := filePath + dirname
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		return err
	}
	return err
}

func DeleDir(dirname string) error {
	dir := filePath + dirname
	err := os.RemoveAll(dir)
	if err != nil {
		return err
	}
	return err
}

// 将字符串添加写入文本文件
func WriteToTxt(content, bookname, dirname string) {
	filepath := filePath + dirname + "/" + bookname // 存放小说的TXT文件路径
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(content)
	write.Flush()
	fmt.Println("写入", bookname, "完成")
}

func CreateToTxt(content string, fileName string, dirname string) {
	filepath := filePath + dirname + "/" + fileName // 存放小说的TXT文件路径
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(content)
	write.Flush()
	fmt.Println("写入", fileName, "完成")
}
