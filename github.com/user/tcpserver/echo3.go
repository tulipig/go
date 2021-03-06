package main

import (
    "fmt"
    "log"
    "time"
    "bufio"
    "strings"
    "net"
    "sync"
)

func main() {
    listenner,err := net.Listen("tcp","localhost:8000")
    if err!=nil{
        log.Fatal(err)
    }

    for{
        conn,err := listenner.Accept()
        if err!=nil{
            log.Print(err)
            continue
        }
        handleConn(conn)
        //go handleConn(conn)
    }
}

func echo(c net.Conn, shout string, delay time.Duration) {
    fmt.Fprintln(c, "\t", strings.ToUpper(shout))
    time.Sleep(delay)
    fmt.Fprintln(c, "\t", shout)
    time.Sleep(delay)
    fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
    input := bufio.NewScanner(c)
    var wg sync.WaitGroup
    for input.Scan() {
        wg.Add(1)
        go echo(c, input.Text(), 1*time.Second)
    }
    // NOTE: ignoring potential errors from input.Err()

    go func(){
        wg.Wait()
        c.Close()
    }()
}