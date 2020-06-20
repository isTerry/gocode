package main

import (
	"bytes"
	"fmt"
	"io"
)

//interfaces is the reason why go is maintainable and scalable
/*An interface type is defined as a set of method signatures.
A value of interface type can hold any value that implements those methods.
即实现接口就要实现它所有的方法*/
func main() {
	//empty interfaces, to handle multiple type conversions
	var myObj interface{} = NewBufferedWriterCloser()
	//var myObj Empty = NewBufferedWriterCloser()
	if wc, ok := myObj.(WriterCloser); ok {
		wc.Write([]byte("Hello this is Terry and Lisa"))
		wc.Close()
	}
	r, ok := myObj.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}

	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello this is Terry and Lisa"))
	wc.Close()

	//type conversion
	//bwc can access the internal filed of BufferedWriterCloser that wc can't
	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)

	//Impossible type assertion: '*io.Reader' does not implement 'WriterCloser'
	//ior := wc.(*io.Reader)
	ior, ok := wc.(io.Reader)
	if ok {
		fmt.Println(ior)
	} else {
		fmt.Println("Conversion failed")
	}

	//myInt := IntCounter(0)
	//var inc Incrementer = &myInt
	//for i := 0; i < 10; i++ {
	//	fmt.Println(inc.Increment())
	//}

	//polymorphic behaviour
	/*Cannot use 'ConsoleWriter{}' (type ConsoleWriter)
	as type WriterCloserType does not implement 'WriterCloser'
	as 'Close' method has a pointer receiver
	如果使用value类型，要求接口所有方法被value receiver实现*/
	//var w WriterCloser = ConsoleWriter{}
	/*如果使用pointer类型，接口所有方法可以被value 或者 pointer receiver实现*/
	//TODO why???
	var w WriterCloser = &ConsoleWriter{}
	w.Write([]byte("hello Terry"))
	fmt.Println(w)
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

//-----------------------------------------
//"er"
type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

//embedding interfaces
type WriterCloser interface {
	Writer
	Closer
}

//-----------------------------------------
//implement with values vs pointers
type ConsoleWriter struct {
}

//implicit implement
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

//func (cw ConsoleWriter) Close() error  {
func (cw *ConsoleWriter) Close() error {
	return nil
}

//-----------------------------------------
type BufferedWriterCloser struct {
	buff *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buff.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buff.Len() > 8 {
		_, err := bwc.buff.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, nil
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	//flush
	for bwc.buff.Len() > 8 {
		data := bwc.buff.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{buff: bytes.NewBuffer([]byte{})}
}

type Empty interface {
}
