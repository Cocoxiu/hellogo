package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func writeData(intChan chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		intChan <- i
		fmt.Println("写入的数字是", i)
		time.Sleep(time.Second)
	}
	close(intChan)
}

func readData(intChan chan int) {
	defer wg.Done()
	for v := range intChan {
		fmt.Println("读取的数字是", v)
		time.Sleep(time.Second)
	}

}
func main() {
	intChan := make(chan int, 50)
	wg.Add(2)
	go writeData(intChan)
	go readData(intChan)
	wg.Wait()
}
