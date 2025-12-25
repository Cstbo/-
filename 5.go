package main

import "fmt"

func main() {
    var x string = 123
    fmt.Println(notDefined)

    var y int = "abc"
    z := []int{"a", "b"}
    m := map[string]int{ "x": "1" }

    var a [3]int
    a[5] = 10

    const c string = 1.23

    var s struct{ A int }
    s.B = 10

    type MyInt string
    var n MyInt = 1

    var ch chan int
    close(nil)

    len(123)
    cap("hello")

    make(int, 10)




    if 1 {
        fmt.Println("hi")
    }

    for range 123 {
    }

    var _ = []int{1}[10]
    var _ = "x"[2]

    var _ = make([]int, -1)
    var _ = len(3.14)
}
