package main

import (
	"fmt"
	"reflect"
)

//collections of disparate data types
type doctor struct {
	number     int
	doctorName string
	episodes   []string
	companions []string
}

type Animal struct {
	//backtick反引号`  semicolon分号;  colon冒号:  quotation mark引号""
	Name   string `required max:"100"`
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

type outer struct {
	numbers []int
	in      inner
	ptr     *inner
}

type inner struct {
	name string
}

func main() {
	//使用new得到的struct或者基本类型的变量，为指针类型
	ptr := new(inner)
	//并且可以通过语法糖直接访问
	fmt.Println(ptr.name)

	a := outer{
		numbers: []int{1, 2, 3},
		in:      inner{name: "Terry"},
		ptr:     &inner{name: "Jack"},
	}
	fmt.Printf("in val:%v\n", a)
	fmt.Printf("in ptr:%p\n", &a)

	//structPassByValue(a)
	/*in val:{[1 2 3] {Terry} 0xc0000421f0}
	in ptr:0xc000068330
	in func val:{[1 2 3] {Terry} 0xc0000421f0}
	in func ptr:0xc000068390
	amended val:{[1 2 3 1] {amended} 0xc000042200}
	amended ptr:0xc000068390
	out val:{[1 2 3] {Terry} 0xc0000421f0}
	out ptr:0xc000068330
	*/

	//a = structPassByValue(a)
	/*in val:{[1 2 3] {Terry} 0xc0000421f0}
	in ptr:0xc000068330
	in func val:{[1 2 3] {Terry} 0xc0000421f0}
	in func ptr:0xc000068390
	amended val:{[1 2 3 1] {amended} 0xc000042200}
	amended ptr:0xc000068390
	out val:{[1 2 3 1] {amended} 0xc000042200}
	out ptr:0xc000068330
	*/
	//fmt.Printf("out val:%v\n", a)
	//fmt.Printf("out ptr:%p\n", &a)

	out := structPassByValue(a)
	/*in val:{[1 2 3] {Terry} 0xc0000421f0}
	in ptr:0xc000068330
	in func val:{[1 2 3] {Terry} 0xc0000421f0}
	in func ptr:0xc0000683c0
	amended val:{[1 2 3 1] {amended} 0xc000042200}
	amended ptr:0xc0000683c0
	out val:{[1 2 3 1] {amended} 0xc000042200}
	out ptr:0xc000068390
	*/
	fmt.Printf("out val:%v\n", out)
	fmt.Printf("out ptr:%p\n", &out)
}

//struct值传递，内部相应地也为值传递，深拷贝
func structPassByValue(a outer) outer {
	fmt.Printf("in func val:%v\n", a)
	fmt.Printf("in func ptr:%p\n", &a)
	a.numbers = append(a.numbers, 1)
	a.in.name = "amended"
	a.ptr = &inner{name: "Lisa"}
	fmt.Printf("amended val:%v\n", a)
	fmt.Printf("amended ptr:%p\n", &a)
	return a
}

func tagsHandle() {
	//get by reflect
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
	//parse and deal with it
}

//embedding(composition,independent), "has-a" relationship
//, such as embed base controller into custom controller
//different from
//common behavior use inheriting, "is" relationship
func birdAnimal() {
	b := Bird{
		Animal: Animal{
			Name:   "J",
			Origin: "A",
		},
		SpeedKPH: 100,
		CanFly:   false,
	}
	b.SpeedKPH = 99
}

func structAnonymous() {
	aDoctor := struct {
		number     int
		companions []string
	}{number: 1, companions: []string{"Lisa", "Jack"}}
	fmt.Println(aDoctor)

	//struct without & is passed by value
	sc := aDoctor
	sc.number = 2
	fmt.Println(aDoctor, sc)

	//&取地址
	sp := &aDoctor
	sp.number = 3
	fmt.Println(aDoctor, sp)
}

func strutsOperation() {
	//position syntax, not recommended
	//ask for all ordered elements, so can cause maintenance problems for same types
	/*aDoctor := doctor{
		1,
		"Terry",
		[]string{"Lisa", "Zoe", "Tom"},
		[]string{"chapter1", "chapter2"},
	}*/
	aDoctor := doctor{
		number:     1,
		doctorName: "Terry",
		companions: []string{"Lisa", "Zoe", "Tom"},
	}
	fmt.Println(aDoctor.doctorName, aDoctor.episodes, aDoctor.companions[1])
	fmt.Println(len(aDoctor.companions), cap(aDoctor.companions))
}

func mapOperation() {
	nameId := map[string]string{
		"Terry": "sa1",
		"John":  "sa2",
	}
	//Invalid map key type: the comparison operators
	//== and != must be fully defined for key type
	//m := map[[]string]int{}

	//map[John:sa2 Terry:sa1]
	fmt.Println(nameId, nameId["Terr"])
	str, ok := nameId["Terr"]
	// false
	fmt.Println(str, ok)

	/*Map: An empty map is allocated with enough space to hold the
	specified number of elements. The size may be omitted, in which case
	a small starting size is allocated.*/
	statePopulations := make(map[string]int, 10)
	statePopulations = map[string]int{
		"California":   12234560,
		"Texas":        13543678,
		"Florida":      12484762,
		"Pennsylvania": 15124532,
	}
	//map[California:12234560 Florida:12484762 Pennsylvania:15124532 Texas:13543678]
	fmt.Println(statePopulations)
	statePopulations["Georgia"] = 12357347
	//map[California:12234560 Florida:12484762 Georgia:12357347 Pennsylvania:15124532
	//Texas:13543678]   order is not guaranteed
	fmt.Println(statePopulations)

	delete(statePopulations, "Georgia")
	//delete a key that is not existed,
	//If m is nil or there is no such element, delete is a no-op.
	delete(statePopulations, "Georgia")

	//we get value and true, or throw away value just to check existence
	_, okTexas := statePopulations["Texas"]
	fmt.Println(okTexas)
	//print a key that is not existed, we get (0 or "") and false
	nonExistedGeorgia, okGeorgia := statePopulations["Georgia"]
	fmt.Println(nonExistedGeorgia, okGeorgia)

	//map is passed by reference
	sp := statePopulations
	delete(sp, "Texas")
	fmt.Println(sp, len(sp))
	fmt.Println(statePopulations, len(statePopulations))
}

//TODO sync.Map
