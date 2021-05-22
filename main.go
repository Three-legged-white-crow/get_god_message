package main

import (
	"bufio"
	"fmt"
	"get_god_message/stack"
	"math/rand"
	"os"
	"time"
)

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
		return false
	}
	return true
}

func main() {
	var question string
	var yao int
	var pause string
	fmt.Println("请输入您的问题\\Input your question")
	fmt.Scanln(&question)
	// fmt.Printf("回复:%s \n", question)
	myStack := stack.NewStack()
	for i := 0; i < 6; i++ {
		var sum int
		sum = 0
		sequenceNumber := i + 1
		for j := 0; j < 3; j++ {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			result := r.Intn(2)
			if result == 1 {
				// 正面3，阳
				yao = 3
			} else {
				//反面2，阴
				yao = 2
			}
			sum = sum + yao
			fmt.Printf("%d  ", yao)
		}
		fmt.Printf("number %d: \t", sequenceNumber)
		// myStack.Push(sequenceNumber)
		switch sum {
		// 2+2+2,老阴
		case 6:
			fmt.Println("- -,x" + "\n")
			myStack.Push("- -,x" + "\n")
		// 2+2+3，少阳
		case 7:
			fmt.Println("——" + "\n")
			myStack.Push("——" + "\n")
		// 3+3+2，少阴
		case 8:
			fmt.Println("- -" + "\n")
			myStack.Push("- -" + "\n")
		// 3+3+3，老阳
		case 9:
			fmt.Println("——,x" + "\n")
			myStack.Push("——,x" + "\n")
		}
		//模拟手摇
		fmt.Scanln(&pause)
	}
	question = question + "\n"
	myStack.Push(question)

	//创建一个新文件
	Exists("log")
	filePath := "./log/" + time.Now().Format("2006-01-02 15:04:05")
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for i := 0; i < 7; i++ {
		temp := myStack.Pop()
		fmt.Println(temp.(string))
		write.WriteString(temp.(string))
	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()

}
