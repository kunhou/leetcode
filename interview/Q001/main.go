package main

import (
	"fmt"
	"sync"
)

func main() {
	letterCh, numberCh := make(chan bool), make(chan bool)
	var wg sync.WaitGroup
	go func() {
		i := 1
		for {
			<-letterCh
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			numberCh <- true
		}
	}()

	wg.Add(1)
	go func() {
		i := 'A'
		for {
			<-numberCh
			if i > 'Z' {
				break
			}
			fmt.Print(string(i))
			i++
			fmt.Print(string(i))
			i++
			letterCh <- true
		}
		wg.Done()
	}()
	letterCh <- true
	wg.Wait()
}
