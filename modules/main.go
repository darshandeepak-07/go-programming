package main

import (
	"context"
	"fmt"
	"modules/interfaces"
	"modules/package1"
	"sync"
	"time"
)

func main() {
	fmt.Println("Initialising modules...!")

	package1.Hello_Package1()

	rect := interfaces.Rectangle {Width: 34.5,Height: 23.6}

	fmt.Println("Area = ",rect.Area())
	fmt.Println("Perimenter = ",rect.Perimeter())

	printValue("ed")
	printValue(23)
	printValue(34.2343)

	printType("deepak")
	printType(24)

	doc := Document {Title: "Deepak Darshan"}
	image := Image{FileName: "abc.png"}
	printDocumentAndImage(doc)
	printDocumentAndImage(image)

	arr_slice_make_map()
	struct_eg()
	type_assert()
	process("Hello")
	doTimeOut()
	contextWithValues()
	demonstrate_goroutine()
	demonstrate_channel()
	execute_mutex()
}

// Empty interfaces

func printValue(value interface{}) {
	fmt.Println("Value = ",value)
}

// type assertion in interfaces

func printType(i interface{}) {
	value,isTypeMatched := i.(string)

	if isTypeMatched {
		fmt.Println("String value = ",value)
	} else {
		fmt.Println("Not a string")
	}
}

// type assertion in switch

func getType(i interface{}) {

	switch v := i.(type) {

	case string : 
		fmt.Println("String = ",v)

	case int :
		fmt.Println("Integer = ",v)
	
	case float64 :
		fmt.Println("FLoat64 = ",v)
	
	default :
		fmt.Println("Type Unknown")
	}
}

type Printer interface {
	Print() string
}

type Document struct {
	Title string
}

type Image struct {
	FileName string
}

// Methods on struct
func (d Document) Print() string {
	return "Printing document = " + d.Title
}

func (i Image) Print() string {
	return "Printing Image = " + i.FileName
}

func printDocumentAndImage(p Printer) {
	fmt.Println(p.Print())
}

func arr_slice_make_map() {
	arr := [5]int{1,2,3,4,5}
	fmt.Println(arr)

	// Empty array
	var em_ar [4]int
	fmt.Println("Empty array = ",em_ar)
	em_ar[0] = 5
	em_ar[2] = 10
	fmt.Println("array = ",em_ar)

	for i,val := range arr {
		fmt.Printf("Index %d Value %d \n",i,val)
	}

	slice1 := arr[0:4]
	fmt.Println(slice1)

	slice2 := arr[3:5]
	fmt.Println(slice2)

	// changeing 3rd index will also change in slice1 and slice2 and in the array
	slice1[3] = 14
	fmt.Println(arr)
	fmt.Println(slice1)
	fmt.Println(slice2)

	// creating slice with make()

	nums := make([]int,0,10)
	fmt.Println(nums)

	for i:=0;i<6;i++ {
		 // nums[i] = i -> will return index out of range
		 nums = append(nums, i) // will return a new slice when capacity is reached
		 fmt.Printf("Length : %d Capacity %d \n",len(nums),cap(nums))
	}

	fmt.Println(nums)

	// Maps

	var map1 map[int]string
	fmt.Println(map1) // Empty map

	capital := make(map[string]string)

	capital["TamilNadu"] = "Chennai"
	capital["Karnataka"] = "Bangalore"

	for key,val := range capital {
		fmt.Printf("Key = %s and Value = %s \n",key,val)
	}

	// Check if key exists

	value,exist := capital["Kerala"]

	if exist {
		fmt.Println("value exist ",value)
	} else {
		fmt.Println("Value does not exist")
	}

	// Delete a key
	delete(capital,"Karnataka")
	fmt.Println(capital)

	// makes() is used to initialise maps,slice and channels
} 


func struct_eg() {
	// Anonymous struct

	animal := struct {
		Name string
		isDomesticated bool
	} {
		Name: "Lion",
		isDomesticated: false,
	}

	fmt.Println(animal)
}


// Type assertion

type Student struct {
	Name string
	id string
}

type School interface {
	getStudent(id string) Student
}

func (std *Student) getStudent(id string) Student {
	return Student{Name: "Deepak",id: "07"}
}

func (std *Student) listStudents() []Student {
	return []Student{ {Name: "Deepak",id: "07"},{Name: "Darshan",id: "06"}}
}

func type_assert() {
	var instance School = &Student{Name: "Deepak",id: "07"}

	if value,ok := instance.(interface{ listStudents()[]Student });ok {
		students := value.listStudents()

		for  _,student := range students {
			fmt.Println(student.id,student.Name)
		}

	}
}

// type switch

func process[T any](value T) {
	switch v:= any(value).(type) {
	case int :
		fmt.Println("Integer  = ",v)
	case string :
		fmt.Println("String = ",v)
	default :
		fmt.Println("Type not found")
	}
} 


// Context
// for handling timeout , deadlines
// cancelling go routines 
// passing metadata across the go application

func doTimeOut() {
	ctx := context.Background()

	ctxWithTimeout,cancel := context.WithTimeout(ctx,5 * time.Second)

	defer cancel()

	done := make(chan struct{})

	go func() {
		time.Sleep(6 * time.Second)
		close(done)
	}()

	select {
	case <- done:
		fmt.Println("API invoked")
	case <- ctxWithTimeout.Done():
		fmt.Println("APi call timed out : ",ctxWithTimeout.Err())
	}
}

func contextWithValues() {
	ctx := context.Background()

	ctxWithValue := context.WithValue(ctx,"user","123")

	if id,ok := ctxWithValue.Value("user").(string) ; ok {
		fmt.Println("User is ",id)
	} else {
		fmt.Println("User not found")
	}
}

// goroutine

func demonstrate_goroutine() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func ()  {
		defer wg.Done()

		for i := 0 ; i < 20 ; i++ {
			fmt.Println("From function 1 -> ",i)
		}
	}()

	wg.Add(1)
	go func ()  {
		defer wg.Done()
		for i := 0;i<20;i++ {
			fmt.Println("From function 2 -> ",i)
		}	
	}()

	wg.Wait()
	fmt.Println("Process completed")
}


// channels

func demonstrate_channel() {
	c := make(chan int)

	go func()  {
		sum := 0
		for i:=1 ; i<10;i++ {
			sum += i
		}
		c <- sum
	}()

	result := <-c
	fmt.Println("Result from channel ",result)
}


// mutex

type SafeWriter struct {
	mutex sync.Mutex
	Numbers map[string]int
}

func (s *SafeWriter) Add(number int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.Numbers["keyValue"] = number
}

func execute_mutex() {
	s := SafeWriter{Numbers: make(map[string]int) }
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func (i int)  {
			defer wg.Done()
			s.Add(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("Value in Map = ",s.Numbers["keyValue"])
}