package main

import (
	"fmt"
)

const value int8 = 127

/*const (
	a = iota
	b = iota
)*/
const (
	a = iota
	b
)
const (
	c = iota
)

//add errorType
//const (
//	errorSpecialist = iota
//	catSpecialist
//	dogSpecialist
//)
//just throw away zero value
const (
	_ = iota
	catSpecialist
	dogSpecialist
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

func main() {

}

func iotaOperation() {
	fileSize := 40000000.
	//?here division ok with float64/const int
	fmt.Printf("%.2fGB", fileSize/GB)
}

func avoidZeroInitialValueConfused() {
	var specialistType int
	fmt.Printf("%v,%v\n", specialistType == catSpecialist, catSpecialist)
}

func enumConstants() {
	//1,int
	//0,int
	fmt.Printf("%v,%T\n", b, b)
	fmt.Printf("%v,%T\n", a, a)
	//0,int
	fmt.Printf("%v,%T", c, c)
}

func untypedInfer() {
	//Typed
	//const a int = 3
	//Untyped
	const a = 3
	var b int16 = 4
	fmt.Printf("%v,%T", a+b, a+b)
	//equals
	fmt.Printf("%v,%T", 3+b, 3+b)
}

func shadow() {
	const value int = 0
	fmt.Printf("%v,%T", value, value)

	const inName string = "Terry"
	const OutName string = "Green"
	const flag bool = false
	const f float32 = 3.14
	const d = 3.14
	fmt.Printf("%v,%T", d, d)
}
