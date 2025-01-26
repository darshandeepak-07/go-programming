package main

import (
	"fmt"
	"modules/interfaces"
	"modules/package1"
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