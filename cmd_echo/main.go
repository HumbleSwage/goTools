package main

import (
	"fmt"
	"os"
	"strings"
)


func echoArgs1(){
	//方法1
	fmt.Println(strings.Join(os.Args[1:]," "))

}

func echoArgs2(){
	//方法2
	s , sep := "",""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echoArgs3(){
	//方法3
	s , sep := "",""
	for _,args := range os.Args[1:]{
		s += sep + args
		sep = " "
	}
	fmt.Println(s)
}

func main(){
	//比较关键的点在于对os.Args的理解
	echoArgs1()
	echoArgs2()
	echoArgs3()


}