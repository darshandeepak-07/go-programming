package main

import (
	"fmt"
)

func main1() {
    fmt.Println("Hello, World!")
    varibles_eg()
    conditionals_g()

    add,sub,mul,div := operations(23,45)
    
    fmt.Println(add,sub,mul,div)

    result,err := divide(10,0)


    if err != nil {
        fmt.Println("Error: ",err)
    } else {
        fmt.Println("Result: ",result)
    }

    divide_err(3,4)
    divide_err(4,0)
    divide_err(2,6)
}

// Variables

/*
 bool
 string

 int int8 int16 int32 int64
 unit unit8 unit16 unit32 unit64 unitptr

 byte -> alias for unit8

 rune -> alias for int32
    represents a unitcode code point

 float32 float64

 complex64 complex128

*/

func varibles_eg() {
    // bool
    var isCompleted bool = true
    fmt.Println(isCompleted)

    //string
    var name string = "Deepak"
    fmt.Println(name)

    /*  int
    --------------------------------------------------------------------------------
        int	Platform-dependent (32 or 64 bits)	-2³¹ to 2³¹-1 (32-bit) or -2⁶³ to 2⁶³-1 (64-bit)
        int8	8 bits	-128 to 127
        int16	16 bits	-32,768 to 32,767
        int32	32 bits	-2,147,483,648 to 2,147,483,647
        int64	64 bits	-9,223,372,036,854,775,808 to 9,223,372,036,854,775,80
    */

    /*  unit
    -------------------------------------------------------------------------------
        uint	Platform-dependent (32 or 64 bits)	0 to 2³²-1 (32-bit) or 0 to 2⁶⁴-1 (64-bit)
        uint8	8 bits	0 to 255
        uint16	16 bits	0 to 65,535
        uint32	32 bits	0 to 4,294,967,295
        uint64	64 bits	0 to 18,446,744,073,709,551,615
    */

    c := 5
    fmt.Println(c)

    // unitptr 
    // An integer type to gold a pointer's memory address

    // byte
    // Alias for unit8
    var b byte = 'A'
    fmt.Println(b) // Prints ASCII value

    // rune
    // Alias for int32
    // Represents unicode point
    // Useful for working with characters beyond ASCII, eg,EMOJI, non english letters

    var char rune = '#'
    fmt.Println(char)

    // float

    /*
        float32	32 bits	Approx. 6-9 decimal digits
        float64	64 bits	Approx. 15-17 decimal digits

        Default value -> 0.0
    */

    var pi float64 = 3.1415
    fmt.Println(pi)

    // Complex number 

    /*
        complex64	-> 64 bits	32-bit real + 32-bit imaginary
        complex128  -> 128 bits	64-bit real + 64-bit imaginary
    */

    var cmp complex128 = complex(2,4) // 2 + 4i
    fmt.Println(cmp) // (2+4i)
    fmt.Println(real(cmp)) // 2
    fmt.Println(imag(cmp)) // 4
} 

func conditionals_g() {
     for i := 0; i<5; i++ {
        fmt.Println(i)
     }

     // infinite loop

     /*
     for {
        fmt.PrintLn("Infinte loop")
     }
     */

     i := 0
     for {
        if i>=3 {
            break
        }
        fmt.Println(i)
        i++
     }

     // using for like a while loop based on condition
     age := 5

     for age < 10 {
        fmt.Println(age)
        age ++
     }

     numbers := [] int{ 1,2,3,4,5 }
     for index,number := range numbers {
        fmt.Printf("Index : %d, Value: %d\n",index,number)
     } 

     person := map[string] string {
        "name" : "Deepak",
        "age" : "25",
     }

     for key,valule := range person {
        fmt.Printf("Key: %s, Value: %s\n",key,valule)
     }

     name := "Deepak"
     for index,char := range name {
        fmt.Printf("Index: %d, Rune: %c\n",index,char)
     }

     // if-else

     if 5%2 == 0 {
        fmt.Println("5 is even")
     } else {
        fmt.Println("5 is odd")
     }

     if num:=5; num < 0 {
        fmt.Println("num is -ve")
     } else if num < 10 {
        fmt.Println("num has 1 digit")
     } else {
        fmt.Println("num has multiple digits")
     }
}

func func_eg1() string {
     return "Hello"
}

func operations(num1 int, num2 int) (
    add int,
    sub int,
    mul int,
    div float32,
) {
    add = num1 + num2
    sub = num1 - num2
    mul = num1 * num2
    div = float32(num1 / num2)
    return add,sub,mul,div
}


// Errors,Panic and Recover

func divide(a,b int) (int,error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }

    return a/b,nil
}


// custom error

type CustomError struct {
    Code int32
    Message string
}

func(err CustomError) Error() string {
    return fmt.Sprint("Code: %d, Message: %s",err.Code,err.Message)
}

// Error,panic,recover

func handlePanic() {
    var a = recover()

    if(a != nil) {
        fmt.Println(a)
    }
}

func divide_err(n1,n2 int) {
    defer handlePanic()

    if(n2 == 0) {
        err := CustomError{Code:400,Message:"Cannot divide by zero"}
        panic(err.Message)
    } else {
        res := n1 / n2
        fmt.Println(res)
    }
}

