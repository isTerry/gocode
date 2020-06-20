package main

import (
	"fmt"
)

//entry point:always package main, func main takes no parameters and returns no values
func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)
	fmt.Println(v, p)

	//anonymousFunc()

	//s := variantParameters("Terry", 1, 2, 3, 4, 5)
	//fmt.Println(s)

	//p := returnPointer("Terry", 1, 2, 3, 4, 5)

	//variant contains zero parameter
	//p := returnPointer("Terry")
	//fmt.Println(*p)

	//name := "Lisa"
	//passByValue(name)
	////passByPointer(&name)
	//fmt.Println(name)
}

/*为什么用method？
Remember: a method is just a function with a receiver argument.
Receiver in OOP terms would be the “class” that the method is a part of.
Functions are called independently with the arguments specified,
and methods are called on the type of their receiver.
The difference in methods and functions is mostly syntactic,
and you should use the appropriate abstraction depending on the use case
1.Method chaining
p = p.withName("John").withAge(21) 而不是 p = withName(withAge(p, 18), "John")
2.A function reads and modifies a lot of values of a particular type,
it should probably be a method of that type.
3.Semantics(Is the customer an adult?)
// Method
customer.isAdult()
// Function
isAdult(customer)
*/

//定义位置：The receiver is between the func keyword and the method name.

/*can only declare a method with a receiver whose type is defined
in the same package as the method.
Cannot declare a method with a receiver whose type is defined
in another package (which includes the built-in types such as int).
必须用在同一个包定义的type上，但可以 如：type counter int*/
/*Declare methods with pointer receivers.
The literal syntax *T for some type T.
(Also, T cannot itself be a pointer such as *int.)*/

/*Go负责对method中receiver的value、pointer自动双向转换，但function不能转换
functions with a pointer argument must take a pointer,
while methods with pointer receivers take either a value or a pointer as the receiver,
as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5)
since the Scale method has a pointer receiver.
The equivalent thing happens in the reverse direction.
Functions that take a value argument must take a value of that specific type,
while methods with value receivers take either a value or a pointer as the receiver
p := &v
fmt.Println(p.Abs()) // OK
In this case, the method call p.Abs() is interpreted as (*p).Abs().*/

/*要修改值且避免拷贝用pointer receiver:
The first is so that the method can modify the value that its receiver points to
The second is to avoid copying the value on each method call.
This can be more efficient if the receiver is a large struct, for example.*/
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//() means execution
func anonymousFunc() {
	//内部func可以访问外部func变量， 外部不可访问内部
	for i := 0; i < 5; i++ {
		//better practice:pass into the inner func
		//changes in the outer scope are reflected to the inner scope
		func(i int) {
			fmt.Println(i)
		}(i)

		//异步asynchronous时获取i，外部i改变可能和内部不一致
		//variable
		var f func() = func() {
			//f := func() {
			fmt.Println(i)
		}
		f()
	}

	//functions as types, can be variables or arguments or return values
	//divide(5.0, 0.0)
	var divide func(float64, float64) (float64, error)
	//divide(5.0, 0.0)
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("cannot provide zero as second value")
		}
		return a / b, nil
	}
	res, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

//idiomatic way of error return: substitute for panic
func divide(a, b float64) (float64, error) {
	//func divide(a, b float64) float64 {

	//return as soon as possible, no "else"
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot provide zero as second value")
	}
	return a / b, nil

	//if b == 0.0 {
	//	panic("cannot provide zero as second value")
	//}
	//return a / b
}

//name return func, go will initialize it
func nameReturn(values ...int) (result int) {
	for _, v := range values {
		result += v
	}
	return
}

//when func ends, stack removed, go helps promote the pointer memory to the heap
func returnPointer(name string, values ...int) *int {
	fmt.Println(name, values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

//variant parameters, compiler helps change parameters into the slices
//can only have one and has to be the end
//because compiler doesn't know where is the end when it comes into the ... before ...
func variantParameters(name string, values ...int) int {
	fmt.Println(name, values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func passByValue(name string) {
	name = "Terry"
	fmt.Println(name)
}

func passByPointer(name *string) {
	*name = "Terry"
	fmt.Println(*name)
}

//syntax sugar, all one type
//func funcSample(code1, code2 int) {
func funcSample(msg string, code int) {
	fmt.Println(msg, code)
}
