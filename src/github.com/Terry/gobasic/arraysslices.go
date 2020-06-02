package main

import "fmt"

func main() {
	slices()
}

func slices() {
	//slices define:empty[]
	//a := []int{1,2,3,4,5,6}
	//same for array
	a := [...]int{1, 2, 3, 4, 5, 6}

	//基于零开始的index, slice of
	//b := a //array, copy
	//!!!not equals to
	b := a[:] //slice of, point to

	//左闭inclusive
	c := a[1:]
	//右开exclusive
	d := a[:5]
	e := a[1:5]
	/*a:[1 2 3 4 5 6]
	,b:[1 2 3 4 5 6]
	,c:[2 3 4 5 6]
	,d:[1 2 3 4 5]
	,e:[2 3 4 5]*/
	fmt.Printf("a:%v\n,b:%v\n,c:%v\n,d:%v\n,e:%v\n", a, b, c, d, e)

	//point to,could change source
	b[0] = 0
	/*a:[0 2 3 4 5 6]
	,b:[0 2 3 4 5 6]
	,c:[2 3 4 5 6]
	,d:[0 2 3 4 5]
	,e:[2 3 4 5]*/
	fmt.Printf("a:%v\n,b:%v\n,c:%v\n,d:%v\n,e:%v\n", a, b, c, d, e)
}

func arrayPointAndCopy() {
	a := [...]int{1, 2, 3}

	//point to
	//b := &a
	//copy, if it's big,cost time
	b := a

	b[0] = 0
	//a: [1 2 3] b: [0 2 3]
	fmt.Println("a:", a, "b:", b)
}

func matrix() {
	//an array of arrays
	//var numberMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	//delete redundant
	var numberMatrix = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(numberMatrix)
}

func arrayDefine() {
	//omitted
	//var grades [3]int = [3]int{1, 2, 3}
	//var grades = [3]int{1,2,3}
	//grades := [3]int{1,2,3}
	grades := [...]int{1, 2, 3}
	fmt.Printf("%v\n", grades)

	//var students [...]string
	//use of [...] array outside of array literal
	//fmt.Printf("%v,%T", students, students)

	var ids [3]int
	var names [6]string
	var colors [0]string
	//[  ],[3]string,3,3
	fmt.Printf("%v,%T,%v,%v\n", ids, ids, cap(ids), len(ids))
	//[     ],[6]string
	fmt.Printf("%v,%T\n", names, names)
	//[],[0]string
	fmt.Printf("%v,%T\n", colors, colors)

	//invalid array index 0 (out of bounds for 0-element array)
	//colors[0] = "ccc"

	//[4 0 0],[3]int,3,3
	ids[0] = 4
	fmt.Printf("%v,%T,%v,%v", ids, ids, cap(ids), len(ids))
}
