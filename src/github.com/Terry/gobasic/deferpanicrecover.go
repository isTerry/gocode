package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//defer类似于finally(执行顺序不同)
//panic类似于throw(可以re throw)
//recover类似于catch(当前func执行完defer后不能正常结束，调用链上方可以)
func main() {
	callPanickier()
}

func callPanickier() {
	fmt.Println("start")
	panickier()
	fmt.Println("end")
}

func panickier() {
	fmt.Println("about to panic")
	//匿名函数
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			//re panic
			panic(err)
		}
	}()
	panic("guys something happened")
	//will not get executed
	fmt.Println("done panicking")
}

func deferPanic() {
	//first defer(because defer is set, in order to close resources)
	//then panic, then returning
	fmt.Println("start")
	defer fmt.Println("this deferred")
	panic("guys something happened")
	fmt.Println("end")

	//first panic, then returning
	fmt.Println("start")
	panic("guys something happened")
	defer fmt.Println("this deferred")
	fmt.Println("end")
}

//panic means can't figure it out what to do, when program cannot continue at all
//什么时候用panic?如除以零（Go不知道如何处理的时候）；端口已经被占用。
//但是如http返回404；文件无法打开（关键文件除外）；即直接返回错误，不作为panic
/*The panic built-in function stops normal execution of the current goroutine.
When a function F calls panic, normal execution of F stops immediately.
Any functions whose execution was deferred by F are run in the usual way,
and then F returns to its caller. To the caller G,
the invocation of F then behaves like a call to panic,
terminating G's execution and running any deferred functions.
This continues until all functions in the executing goroutine have stopped,
in reverse order. At that point, the program is terminated with a non-zero exit code.
This termination sequence is called panicking
and can be controlled by the built-in function recover.*/
func panicSample() {
	//allow parallel run设置
	/*panic: listen tcp :8080: bind: Only one usage of each socket address (protocol/n
	etwork address/port) is normally permitted.

	goroutine 1 [running]:
	main.panicSample()
	        C:/Users/iswei/gocode/src/github.com/Terry/gobasic/deferpanicrecover.go:
	31 +0x135
	main.main()
	        C:/Users/iswei/gocode/src/github.com/Terry/gobasic/deferpanicrecover.go:
	11 +0x27
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from a HandleFunc #1!\n"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

func deferValue() {
	//we get "start", because a is set when it comes to defer,
	//print execute before returning
	a := "start"
	defer fmt.Println(a)
	a = "end"
}

func deferSample() {
	//to associate the open and close of resources right to next each other
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	//if it's in loop, not suggested to do so,
	//because all resources will be open at the same time before get closed
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func deferOrder() {
	//start
	//end
	//middle
	//fmt.Println("start")
	//defer fmt.Println("middle")
	//fmt.Println("end")

	//栈顺序, last in first out, close resources
	//end
	//middle
	//start
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")

}
