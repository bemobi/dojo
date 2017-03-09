package main

import (
	"fmt"

	"github.com/bemobi/dojo/go/codec"
)

func main() {
	fmt.Println(codec.Encode("1"))
	fmt.Println(codec.Encode("0"))
	fmt.Println(codec.Encode("939393"))
	fmt.Println(codec.Encode("939394"))
	fmt.Println(codec.Encode("14776336"))
	fmt.Println(codec.Encode("14776335"))
}
