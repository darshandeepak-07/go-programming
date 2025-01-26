package package1
import (
	"fmt"
	"sync"
)

func PrintOddEven() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func ()  {
		defer wg.Done()
		
		for i:=0;i<=10;i++ {
			if i % 2 == 0 {
				fmt.Println(i)
			}
		}
	}()

	wg.Add(1)

	go func ()  {
		defer wg.Done()
		
		for i := 0; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Println(i)
			}
		}
	}()

	wg.Wait()
}

func PrintUsingChannel() {
	oddChannel := make(chan bool)
	evenChannel := make(chan bool)

	go func ()  {
		for i:=1 ;i<11 ; i+=2 {
			<-oddChannel
			fmt.Println("Odd : ",i)
			evenChannel <- true
		}
		close(evenChannel)	
	}()

	go func ()  {
		for i := 2; i < 11; i+=2 {
			<-evenChannel
			fmt.Println("Even : ",i)
			oddChannel <- true
		}
		close(oddChannel)
	}()

	oddChannel <- true
	<-evenChannel
}