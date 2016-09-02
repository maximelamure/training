package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	max := boring("Max your are boring !!!")
	fred := boring2("Fred your are boring !!!")
	scott := FanIn(max, fred)
	for i := 0; i < 10; i++ {
		fmt.Println(<-scott)
	}
}

func FanIn(max, fred chan string) chan string {
	scott := make(chan string)

	go func() {
		for {
			select {
			case c := <-max:
				scott <- c
			case c := <-fred:
				scott <- c
			case <-time.After(1 * time.Second):
				fmt.Println("Sorry too slow !!!")
			}
		}
	}()

	return scott
}

func boring(text string) chan string {
	ch := make(chan string)

	go func() {
		for index := 0; ; index++ {
			time.Sleep(2 * time.Second)
			ch <- text + strconv.Itoa(index)
		}
	}()

	return ch
}

func boring2(text string) chan string {
	ch := make(chan string)

	go func() {
		for index := 0; ; index++ {
			time.Sleep(2 * time.Second)
			ch <- text + strconv.Itoa(index)
		}
	}()

	return ch
}
