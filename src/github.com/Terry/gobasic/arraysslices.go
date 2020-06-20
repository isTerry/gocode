package main

import (
	"fmt"
)

func main() {
	cutSlices()
}

func cutSlices() {
	//移出中间以及最后一个
	a := []int{1, 2, 3, 4, 5}

	/*It is therefore necessary to store the result of append,
	often in the variable holding the slice itself*/
	//相当于slices a进行append空，result仍然是a，返回给b
	//b := append(a)
	//b[0] = 2
	//fmt.Println(a, "/", b)

	//FIXME need to allocate to append, b is on new array
	b := append(a[:2], a...)

	//FIXME why a changed into that? a[:2] is 1 2, a[3:len(a)-1] is 4
	//	1 2 3 4 5==>1 2 [3] 4 5==>1 2 4 4 5
	//	append here is moving 4 into the existing 3's location
	//	no need to allocate, b is pointing to a
	//	so be careful if there is another reference to a
	//	if don't wanna change a, use loop and do a copy
	//b := append(a[:2], a[3:len(a)-1]...)

	//[1 2 3 4 5] / [1 2 1 2 3 4 5]
	//[1 2 4 4 5] / [1 2 4]
	fmt.Println(a, "/", b)

	b[0] = 0
	//[1 2 3 4 5] / [0 2 1 2 3 4 5]
	//[0 2 4 4 5] / [0 2 4]
	fmt.Println(a, "/", b)

	//len: 7 cap: 10
	//len: 3 cap: 5
	fmt.Println("len:", len(b), "cap:", cap(b))
}

func stringByteSlices() {
	/*As a special case, it is legal to append a string to a byte slice, like this:
	slice = append([]byte("hello "), "world"...)
	*/
	a := "hello world"
	b := []byte(a)
	c := a[:]

	d := []byte("hello")
	d = append(d, " world"...)
	d = append(d, a...)
	d = append(d, b...)
	d = append(d, c...)
	fmt.Println(a, b, c, string(d), string(d[0]))
}

func byteRune() {
	//i对应编码值105，未超出byte及rune范围
	//uint8 byte
	a := []uint8{'i', 'i'}
	var b byte = 'i'
	//int32 rune
	c := 'i'
	var d rune = 'i'

	//[105 105] 105 105 i
	fmt.Println(a, b, c, string(d))
	//105,105,1101001,i
	//原样输出值,整型，binary，char
	fmt.Printf("%v,%d,%b,%c", d, d, d, d)
}

func resizeCapacity() {
	//double every time
	var a []int
	fmt.Println("len:", len(a), "cap:", cap(a))
	a = append(a, 0, 1)
	fmt.Println("len:", len(a), "cap:", cap(a))
	for i := 2; i < 33; i++ {
		a = append(a, i)
		fmt.Println("len:", len(a), "cap:", cap(a))
	}
}

func makeSlices() {
	//s := []int{} 空slices定义 替换为nil
	var s []int
	//[],len:0,cap:0
	fmt.Printf("%v,len:%v,cap:%v\n", s, len(s), cap(s))
	//true
	fmt.Println(s == nil)

	s = []int{1, 2, 3, 4}
	s = make([]int, 3)
	//[0 0 0],len:3,cap:3
	fmt.Printf("%v,len:%v,cap:%v\n", s, len(s), cap(s))
	s = make([]int, 3, 100)
	//[0 0 0],len:3,cap:100
	fmt.Printf("%v,len:%v,cap:%v\n", s, len(s), cap(s))
	s = make([]int, 0, 0)
	//[],len:0,cap:0
	fmt.Printf("%v,len:%v,cap:%v\n", s, len(s), cap(s))
	//false
	fmt.Println(s == nil)
}

func slicesOf() {
	//slices define:empty[]
	a := []int{1, 2, 3, 4, 5, 6}
	//same for array
	//a := [...]int{1, 2, 3, 4, 5, 6}

	//基于零开始的index, slice of
	//FIXME if a is array:copy; if a is slices:point to
	b := a
	//!!!not equals to
	//FIXME slice of, point to
	//	as len(b)<=len(a), there is no need to allocate new array, so just point to
	//b := a[:]

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

//TODO list.List
