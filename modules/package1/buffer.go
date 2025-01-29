package package1

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func Buffer_test1() {
	var bf bytes.Buffer

	bf.WriteString("Deepak Darshan")
	bf.Write([]byte("Darshan")) // appends to the current buffer

	fmt.Println(bf.String())
	fmt.Println(bf.Len())
	p := make([]byte,6)
	bf.Read(p)
	fmt.Println("From buffer Read : ",string(p))
}

// buffered writer -> bufio
// provides buffered I/O for reading and writing. It wraps objects like files or network connections to improve efficiency by minimizing system calls.

func Buffered_Writer() {
	file,_ := os.Create("file1.txt")
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("Buffered eg using bufio")
	writer.Flush()
}

func Buffered_Reader() {
	file,_ := os.Open("file1.txt")
	defer file.Close()

	reader := bufio.NewReader(file)

	line,_  := reader.ReadString('\n')
	fmt.Println("FRom file : ",line)
}


// buffering in channels
// Unbuffered channels require both the sender and receiver to be ready at the same time
// Buffered channels allow sending and receiving independently until the buffer is full or empty.

func Chan_buffer1() {
	defer handlePanic()
	channel := make(chan int,2)
	monitor := make(chan bool)
	channel <- 1
	channel <-2
	go monitor_channel(channel,monitor)
	//channel <-3
	//<-monitor
	go func ()  {
		for i := range monitor {
			fmt.Println(i)
		}	
	}()
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	//fmt.Println(<-channel)
}

func handlePanic() {
	a := recover()

	if a != nil {
		fmt.Println("Panic occured ",a)
	}
}

func monitor_channel(source <-chan int,monitor chan bool) {
	for i := range source {
		fmt.Println(i)
		monitor <- true
	}
	close(monitor)
}