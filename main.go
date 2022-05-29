package main

import (
	"fmt"
	"log"
	"os"
)

const (
	msg = "======================================笔趣阁小说爬虫v1.0======================================\n" +
		"免责声明：本爬虫仅供资料学习，请勿滥用，地址：https://github.com/caicheng918/biqugeNovelCrawl\n" +
		"测试阶段存在bug，不是所有小说都可以爬取。\n" +
		"==============================================================================================="

	filePath = "./" //项目根目录前缀
)

func main() {
	var bookname, author string
	println(msg)
	println("第一步，请输入书名和作者，以空格分开:")
	fmt.Scanln(&bookname, &author)
	err := MkdirForTxt(bookname)
	if err == nil {
		println("已创建", bookname, "文件夹")
	}
	var secModeNum int
	println("第二步，请选择模式编号:\n1.下载模式\n2.小说追更\n输入编号:")
	fmt.Scanln(&secModeNum)
	if secModeNum == 1 {
		var thirdModeNum int
		var startUrl string
		go func() {
			startUrl, err = FindStartUrl(bookname, author)
			if err != nil {
				log.Fatal(err)
			}
		}()
		println("已选择下载模式\n第三步，请选择下载模式编号:\n1.单章逐个下载\n2.合集下载\n输入编号立即爬取:")
		fmt.Scanln(&thirdModeNum)
		if thirdModeNum == 1 {
			//单章逐个下载
			println("开始爬取", bookname, "...")
			CrawlOneByOne(startUrl, bookname)
			return
		}
		if thirdModeNum == 2 {
			//合集下载
			filepath := filePath + bookname + ".txt" // 存放小说的TXT文件路径
			file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()
			println("开始爬取", bookname, "...")
			CrawlAll(startUrl, bookname)
			return
		}
	}
	if secModeNum == 2 {
		var thirdModeNum int
		println("已选择追更模式\n第三步，请选择追更模式编号:\n1.设置任务定时时间\n2.启用默认设置（每天12点与18点检查是否更新）\n输入编号:")
		if thirdModeNum == 1 {
			//设置任务定时时间

			return
		}
		if thirdModeNum == 2 {
			//默认设置

			return
		}
	}

}
