package package1

import (
	"bytes"
	"fmt"
)

func Buffer_test1() {
	var bf bytes.Buffer

	bf.WriteString("Deepak Darshan")
	bf.Write([]byte("Darshan"))

	fmt.Println(bf.String())
	fmt.Println(bf.Len())
	p := make([]byte,5)
	bf.Read(p)
	fmt.Println("From buffer Read : ",p)
}

