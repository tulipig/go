package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    abort := make(chan struct{})
    go func(){
        for{
            os.Stdin.Read(make([]byte,1))
            abort <- struct{}{}
        }
    }()

    fmt.Println("commencing countdown. Press return to abort")
    select{
    case <- time.After(3*time.Second):
        //Do nothing
    case <- abort:
        fmt.Println("Lanch aborted!")
        return
    }

    func(){fmt.Println("Have Lanching...")}()

    tick := time.Tick(10*time.Second)
    fmt.Println(111, time.Now())
    for{
        select{
        case <- tick:
            fmt.Println(222, time.Now())
            return
        case <- abort:
            fmt.Println(333, time.Now())
            tick = time.Tick(10*time.Second)
        }
    }
}