package main

import "fmt"

func main() {
	switchTemplate()
}

func typeSwitch() {
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("int")
		if true {
			break
		}
		fmt.Println("print depends")
	case float32:
		fmt.Println("float32")
	case string:
		fmt.Println("string")
	case []int:
		fmt.Println("[]int")
	case [2]int:
		fmt.Println("[2]int")
	default:
		fmt.Println("---")
	}
}

func switchTemplate() {
	//switch with no tag, more powerful, cases can overlap
	//keyword fallthrough to make falling through
	//careful! fallthrough it's logic less, coder is taking responsibility
	i := 10
	switch {
	case i <= 10:
		fmt.Println("i<=10")
		fallthrough
	case i >= 20, i <= 30:
		fmt.Println("i<=20,i<=30")
	case i <= 40:
		fmt.Println("i<=40")
	default:
		fmt.Println("default")
	}

	//switch on a tag, multiple cases
	//initializer
	switch a := 1 + 1; a {
	case 1:
		fmt.Println("1")
	case 2, 3:
		fmt.Println("2,3")
	default:
		fmt.Println("---")
	}
	//will not falling through, as implicit "break"
	//so case 2 will not print "3", but do nothing
	a := 2
	switch a {
	case 1:
		fmt.Println("11")
	case 2:
	case 3:
		fmt.Println("33")
	default:
		fmt.Println("------")
	}
}

func shortCircuit() {
	//If the first is fulfilled or not fulfilled, may not execute the next
	if true || false || false {
		fmt.Println("always print")
	}
	if false && true && true {
		fmt.Println("always not print")
	}
}

func ifTemplate() {
	stateNumber := map[string]int{
		"aa": 1,
		"bb": 2,
		"cc": 3,
	}
	//Always have curly braces({}花括号) even only one line to execute
	if true {
		fmt.Println("true print")
	}

	//before semicolon is the initializer, after is condition
	if pop, ok := stateNumber["aa"]; ok {
		fmt.Println(pop)
	}
	//Unresolved reference 'pop'
	//fmt.Println(pop)
}
