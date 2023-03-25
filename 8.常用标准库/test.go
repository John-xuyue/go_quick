package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func funcTime() {
	fmt.Printf("hhh\n")
	layout := "2006-01-02 15:04:05"

	now := time.Now()
	nowStr := now.Format(layout)
	fmt.Printf("nowStr: %v\n", nowStr)

	if t, err := time.Parse(layout, nowStr); err == nil {
		fmt.Printf("t: %v\n", t)
	} else {
		fmt.Printf("err: %v\n", err)
	}

	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation(layout, nowStr, loc)
	fmt.Printf("t: %v\n", t)
}

func funcFile() {

	file, err := os.OpenFile("data.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 7777)
	defer file.Close()

	if err != nil {
		fmt.Printf("open err: %v\n", err)
	} else {
		fmt.Println("打开文件成功")
		// 写文件
		writer := bufio.NewWriter(file)
		writer.WriteString("hello\n")
		writer.WriteString("world\n")
		writer.Flush()
	}
}

func funcFile2() {
	file, err := os.OpenFile("data.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 7777)
	defer file.Close()

	if err != nil {
		fmt.Printf("open err: %v\n", err)
	} else {
		fmt.Println("打开文件成功")
		// 读文件
		reader := bufio.NewReader(file)
		for {
			if line, err := reader.ReadString('\n'); err != nil {
				if err == io.EOF {
					if len(line) > 0 {
						fmt.Printf("%v\n", line)
					}
					break
				} else {
					fmt.Printf("read err: %v\n", err)
				}
			} else {
				line = strings.TrimRight(line, "\n")
				fmt.Printf("%v\n", line)
			}
		}
	}

}

func funcLog() {
	log.Printf("log test")

}

func main() {
	// funcTime()
	// funcFile2()
	funcLog()
}
