package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	maxNum := 50

	rand.Seed(time.Now().UnixNano())
	secretNum := rand.Intn(maxNum)
	// fmt.Println("The sec÷ret number is:", secretNum)

	fmt.Println("Please enter your guess number: ")
	// 从标准输入中读取数据
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading", err)
			return
		}
		input = strings.Trim(input, "\r\n")
		// 将输入字符串转换为整形
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input error, please input integer number:", err)
			return
		}
		fmt.Println("You guess number is: ", guess)
		if guess < secretNum {
			fmt.Printf("You guess num %d is less than secretNum, please reenter your guess number:\n", guess)
		} else if guess > secretNum {
			fmt.Printf("You guess num %d is greater than secretNum, please reenter your guess number:\n", guess)
		} else {
			fmt.Printf("Congratulations you, you have guessed secret number: %d", guess)
			break
		}

	}
}
