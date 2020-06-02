package main

import (
	"fmt"
	"strconv"
)

func main() {
	//	TODO add
	//	FIXME change
}

func stringIntConvert() {
	fmt.Println("gobasic, world")
	var a int = 2147483648
	str := strconv.FormatInt(int64(a), 10)
	fmt.Printf("%v,%T\n", str, str)
	c, _ := strconv.ParseInt(str,10,64)
	fmt.Printf("%v,%T\n", c, c)
}
