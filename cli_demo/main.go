package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	//os.Stdin用于读取用户在命令行的输入
	fmt.Println("What's your name?")
	reader := bufio.NewReader(os.Stdin)
	text,_ := reader.ReadString('\n') 
	fmt.Printf("Your name is : %s",text)
}