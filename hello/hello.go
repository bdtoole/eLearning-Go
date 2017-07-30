package main

import "fmt"

func main() {
    a,b := 3,5
    fmt.Printf("hello, world\n")
    fmt.Printf("%v + %v = %v\n",a,b,add(a,b))
    fmt.Printf("Swap values a=%v, b=%v\n",a,b)
    a,b = swap(a,b)
    fmt.Printf("a=%v, b=%v\n",a,b)
}

func add(n1, n2 int) int {
    return n1 + n2
}

func swap(a, b int) (x, y int){
    x,y = b,a
    return
}