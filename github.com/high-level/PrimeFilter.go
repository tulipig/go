package main

import (
    "fmt"
)

func GenerateNatural() chan int {
    ch := make(chan int)
    go func(){
        for i := 2; ; i++ {
            ch <- i
            fmt.Println("GenerateNatural: ",i)
        }
    }()

    return ch
}

func PrimeFilter(in <-chan int, prime int) chan int {
    out := make(chan int)
    go func(){
        for{
            if i := <-in; i%prime != 0{
                out <- i
                fmt.Println("PrimeFilter: ",i)
            }
        }
    }()

    return out
}

func main() {
    ch := GenerateNatural()
    fmt.Println(cap(ch))
    for i := 0; i < 10; i++ {
        prime := <-ch
        fmt.Printf("%v: %v\n", i+1, prime)
        ch = PrimeFilter(ch,prime)
    }
}