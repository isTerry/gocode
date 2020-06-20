package main

import "fmt"

type myStruct struct {
	foo int
}

func main() {
	//pointerCreate()
	//pointToStruct()
	initialPointer()
}

//slices and maps, are projections of underlying data,
//contain the pointer, not the data

func initialPointer() {
	var ms *myStruct
	//<nil>
	fmt.Println(ms)

	//panic: runtime error: invalid memory address or nil pointer dereference
	//fmt.Println(ms.foo)

	//"new" keyword, can't initialize fields at the same time
	ms = new(myStruct)
	//&{0}, this is go initialize
	fmt.Println(ms)

	//dereference operator(asterisk*) has a lower precedence than dot(.) operator
	(*ms).foo = 1
	fmt.Println((*ms).foo)
	//here is syntax sugar:
	//the pointer doesn't have the field, compiler helps interpret it
	ms.foo = 2
	fmt.Println(ms.foo)
}

func pointToStruct() {
	//var ms myStruct
	//ms = myStruct{foo: 1}
	//{1}
	//fmt.Println(ms)

	var ms *myStruct
	//use & before initializer
	ms = &myStruct{foo: 1}
	//&{1} {1} 0xc000006028
	fmt.Println(ms, *ms, &ms)
}

func pointerCreate() {
	//preceding asterisk(*) keyword
	/*1 0xc0000140c0 1 0xc0000140c0 0xc000006028
	2 0xc0000140c0 2 0xc0000140c0 0xc000006028
	3 0xc0000140c0 3 0xc0000140c0 0xc000006028
	*/
	a := 1
	//& address of operator
	//var b = &a
	var b *int = &a
	fmt.Println(a, &a, *b, b, &b)
	a = 2
	fmt.Println(a, &a, *b, b, &b)
	*b = 3
	fmt.Println(a, &a, *b, b, &b)

	//pointer math not allowed in go
	i := []int{1, 2, 3}
	j := &i[0]
	k := &i[1]
	//not allowed, if you really need, check unsafe package
	//k := &i[1]  -8
	//[1 2 3],0xc00006c160,0xc00006c168,1,2
	fmt.Printf("%v,%p,%p,%v,%v\n", i, j, k, *j, *k)
}
