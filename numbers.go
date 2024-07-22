package main

import (
	"fmt"
	"sync"
)

func numbersPrint(num int) {
	var intData = make(chan *int, 500)

	var recivedChan = make(chan *int, 500)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {

		for i := 1; i < num; i++ {

			temp := i
			intData <- &temp
		}

		close(intData)
		wg.Done()
	}()

	go func() {
		defer close(recivedChan)
		for data := range intData {
			fmt.Println("data from the recived side", *data)

			// recivedChan <- data
		}
	}()

	// for data := range recivedChan {
	// 	fmt.Println("data from the recived side", *data)
	// }
	wg.Wait()
	// fmt.Print("num data", nums)
}
