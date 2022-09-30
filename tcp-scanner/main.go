package main

import (
	"fmt"
	"net"
	_"sync"
	_"time"
	"sort"
)

func worker(ports chan int,results chan int){
	for p := range ports {
		address := fmt.Sprintf("20.194.168.28:%d",p)
		conn ,err := net.Dial("tcp",address) 

		if err != nil {
			results <- p
			continue
		}
		conn.Close()
		results <- p
	}
}

func main(){
	//这个地方必须设置一定的缓冲，否则不具备缓冲功能，只能放一个
	ports := make(chan int,100)
	//注意这个results这个没有缓冲，因为只有main进程会使用到它
	results := make(chan int)
	var openPorts []int
	var closePorts []int
	for i := 0 ; i < cap(ports) ; i++ {
		//这里意味着会开启100个goroutine
		go worker(ports,results)
	}
	//收集结果一定要分配工作之前
	go func(){
		for i := 1; i < 1024; i++ {
			port := <- results
			if port != 0 {
				openPorts = append(openPorts,port)
				continue
			}
			closePorts = append(closePorts,port)
		}
	}()
	//分配工作
	for i := 1; i < 1024; i++ {
		ports <- i	
	}
	close(ports)
	close(results)
	//将最后的结果排序
	sort.Ints(openPorts)
	sort.Ints(closePorts)
	for _,port := range closePorts{
		fmt.Printf("端口%d 关闭了\n",port)
	}

}

// //未使用worker池，打印出来的顺序是乱的
// func main(){
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	//使用匿名函数（记得调用）
// 	for i := 21; i < 65535; i++ {
// 		wg.Add(1)
// 		go func(j int){
// 			defer wg.Done()
// 			address := fmt.Sprintf("20.194.168.28:%d",j)
// 			conn , err := net.Dial("tcp",address)
// 			if err != nil {
// 				fmt.Printf("%s 关闭了\n",address)
// 				return
// 			}
// 			conn.Close()
// 			fmt.Printf("%s 打开了\n",address)
// 		}(i)
// 	}
// 	//会在这里阻塞，直到计数器为0
// 	wg.Wait()
// 	elapsed := time.Since(start) / 1e9
// 	fmt.Printf("\n \n扫描21～65535端口总共花费%dseconds",elapsed)
// }

// //单线程，非并发的方式来执行，效率非常的慢
// func main(){
// 	for i := 21; i < 200; i++ {
// 		address := fmt.Sprintf("20.194.168.28:%d",i)
// 		conn , err := net.Dial("tcp",address)
// 		if err != nil {
// 			fmt.Printf("%s 关闭了\n",address)
// 			continue
// 		}
// 		conn.Close()
// 		fmt.Printf("%s 打开了\n",address)
// 	}
// }