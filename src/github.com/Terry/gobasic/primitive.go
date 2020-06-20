package main

import (
	"fmt"
	"github.com/Terry/checks"
	"math"
	"strconv"
)

func main() {

}

func scopeIfElse() {
	if a := 100; a != 100 {

	} else {
		fmt.Println(a)
	}
}

//float is the approximation of decimal
//use error value(precision) to check
func approximation() {
	//myNum := 0.1
	myNum := 0.123
	//if myNum==math.Pow(math.Sqrt(myNum),2){
	//make sure the precision(0.001) is
	//sufficiently large to catch all cases, but sufficiently small not to affect result
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println(myNum, "same")
	} else {
		fmt.Println(myNum, "different")
	}

}

func stringIntConvert() {
	fmt.Println("gobasic, world")
	var a int = 2147483648
	str := strconv.FormatInt(int64(a), 10)
	fmt.Printf("%v,%T\n", str, str)
	c, _ := strconv.ParseInt(str, 10, 64)
	fmt.Printf("%v,%T\n", c, c)
}

func checkCapLower() {
	//same package, all lower case
	aDoctor := doctor{
		number:     0,
		doctorName: "Known",
		episodes:   []string{"ss", "bb"},
		companions: nil,
	}
	fmt.Println(aDoctor)

	//not in the same package
	//can't see unless the element is capitalized
	bDoctor := checks.Doctor{
		Number:     1,
		DoctorName: "Emma",
	}
	fmt.Println(bDoctor)
}
