package main

import (
	"fmt"
	"learn_go/dependencyinjection"
	"learn_go/stdlibnet"
	"log"
	"math"
	"net/http"
)

func main() {
	stdlibnet.Serve()
}

func httpCall() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependencyinjection.MyGreeterHandler)))
}

func slice() {
	c := make([]string, 4)
	fmt.Println(c)

	c[3] = "4"
	fmt.Println(c)
	fmt.Println(len(c))
}

func detect_type(j interface{}) {
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("i'm an int")
		case string:
			fmt.Println("i'm a string")
		default:
			fmt.Println("unknown type", t)
		}
	}
	whatAmI(j)
}

func loops() {
	for i := range 3 {
		fmt.Println("range", i)
	}
}

func constants() {
	const n = 50000
	const d = 3e20 / n

	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sincos(n))

}

func var_syntax() {
	f := "apples"
	fmt.Println(f)
}

func variables() {
	var b, c int = 1, 2
	fmt.Println(b, c)
	var e int
	fmt.Println(e)
}

func values() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println(true && false)
	fmt.Println(true || false)
}
