package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var question string
	var yao int
	var pause string
	fmt.Println("请输入您的问题\\Input your question")
	fmt.Scanln(&question)
	// fmt.Printf("回复:%s \n", question)

	for i := 0; i < 6; i++ {
		var sum int
		sum = 0
		sequenceNumber := i + 1
		for j := 0; j < 3; j++ {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			result := r.Intn(100)
			if result >= 50 {
				yao = 9
			} else {
				yao = 6
			}
			sum = sum + yao
			fmt.Printf("%d  ", yao)
			fmt.Scanln(&pause)
		}
		fmt.Printf("number %d: \t", sequenceNumber)
		switch sum {
		// 6+6+9
		case 21:
			fmt.Println("- -" + "\n")
		// 6+6+6
		case 18:
			fmt.Println("- -,x" + "\n")
		// 9+9+6
		case 24:
			fmt.Println("——" + "\n")
		// 9+9+9
		case 27:
			fmt.Println("——,x" + "\n")
		}
	}

}
