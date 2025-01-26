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

func (d Document) Print() string {
	return "Printing document = " + d.Title
}

func (i Image) Print() string {
	return "Printing Image = " + i.FileName
}

func printDocumentAndImage(p Printer) {
	fmt.Println(p.Print())
}

