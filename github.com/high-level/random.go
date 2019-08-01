package main

import (
    "fmt"
)

func main() {
    ch := make(chan int)
    go func(){
        for{
            select{
                case ch <- 0:
                case ch <- 1:
                case ch <- 2:
                case ch <- 3:
                case ch <- 4:
                case ch <- 5:
                case ch <- 6:
                case ch <- 7:
                case ch <- 8:
                case ch <- 9:

            }
        }
    }()

    for v := range ch{
        fmt.Println(v)
    }
}