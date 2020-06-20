package main

import "fmt"

func main() {
	loopCollections()
}

func loopCollections() {
	//s := "hello world"
	//s := []int{1, 2, 3}
	//s := [...]int{1, 2, 3}
	s := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for k, v := range s {
		fmt.Println(k, v)
	}

	//only key
	//for k := range s {
	//	fmt.Println(k)
	//}

	//only value, this is the wrong way
	//for v := range s {
	//	fmt.Println(v)
	//}

	for _, v := range s {
		fmt.Println(v)
	}
}

func loopLabel() {
OutLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			k := i * j
			fmt.Println(k)
			if k >= 6 {
				break OutLoop
			}
		}
	}
}

func loopingCreate() {
	/*An expression, anything that can be evaluated to produce a value.
	any data by itself is an expression because data always evaluates to itself
	expressions can be made up of expressions*/

	/*A statement is a complete line of code that performs some action,
	does not return anything, the statement itself does not evaluate to anything.*/

	//statement 包含 expression, expression is part of statement
	/*Every expression can be used as a statement
	(whose effect is to evaluate the expression and ignore the resulting value),
	but most statements cannot be used as expressions.*/

	//,逗号comma, we can only use comma to separate elements in go
	//we can't use comma to separate two statements like this in go
	//for i := 0, j:=0; i < 5; i++, j++ {
	//	fmt.Println(i)
	//}

	//i++ is a statement in go, so this is wrong
	//for i,j:= 0,0; i < 5; i,j = i++,j++ {
	//	fmt.Println(i,j)
	//}

	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	//initializer out, increment in
	//substitute for keyword "while"
	i, j := 0, 0
	//for ; i < 5; {
	for i < 5 {
		fmt.Println(i, j)
		i, j = i+1, j+2
	}

	//substitute for keyword "do while"
	//break continue
	k := 0
	for {
		if k%2 == 0 {
			fmt.Println(k)
		}
		k++
		if k == 5 {
			break
		}
	}
}
