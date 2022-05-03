package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

type message struct {
	user  string
	words string
}

var messageChannel = make(chan message, 150)
var wg sync.WaitGroup
var queueLengthRecorder = []int{}

func sender(user string, words string) {
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	messageChannel <- message{user, words}
	fmt.Printf("-sent\n")
	receiver()
}

func receiver() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Printf("-received, %v in buffer\n", len(messageChannel))
	queueLengthRecorder = append(queueLengthRecorder, len(messageChannel))
	message := <-messageChannel
	fmt.Printf(">>>Message: %v\n", message)
	defer wg.Done()
}

func drawQueueLengthChart() {
	fmt.Println("-Queue Length Chart:\n ")
	for _, v := range queueLengthRecorder {
		fmt.Printf("%.3v", v)
		fmt.Print(strings.Repeat("=", v), "\n")
	}
}

func main() {
	fmt.Println("-Starting program\n ")
	for i := 0; i < 10; i++ {
		wg.Add(3)
		go sender("Kiwi", "Awww "+strconv.Itoa(i))
		go sender("Tapa", "Meow "+strconv.Itoa(i))
		go sender("Drum", "Hiss "+strconv.Itoa(i))
	}
	wg.Wait()
	close(messageChannel)
	fmt.Println("-All done!\n ")
	drawQueueLengthChart()
}
