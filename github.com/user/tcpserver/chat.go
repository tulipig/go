package main

import (
    "fmt"
    "net"
    "log"
    "bufio"
    "time"
)

func main() {
    listener,err := net.Listen("tcp","localhost:8000")
    if err!=nil{
        log.Fatal(err)
    }

    go broadfast()

    for{
        conn,err := listener.Accept()
        if err!=nil{
            log.Print(err)
            continue
        }

        go handleConn(conn)
    }
}

type client chan<- string
var(
    entering = make(chan client)
    leaving = make(chan client)
    messages = make(chan string)
)

func broadfast() {
    clients := make(map[client]bool)
    for{
        select{
        case msg := <-messages:
            // Broadcast incoming message to all
            // clients' outgoing message channels.
            for cli := range clients{
                cli <- msg
            }
        case cli := <-entering:
            clients[cli] = true
        case cli := <-leaving:
            delete(clients,cli)
            close(cli)
        }
    }
}

func handleConn(conn net.Conn) {
    fmt.Println("Now Time: ", time.Now())
    //
    tick := time.Tick(60*time.Second)
    go func(c net.Conn){
        select {
        case <- tick:
            fmt.Println("No Msg is coming in 60s, close conn", time.Now())
            leaving <- ch
            messages <- who + " has left because timeout"
            c.Close()
        }
    }(c)

    ch := make(chan string)
    go clientWriter(conn, ch)

    who := conn.RemoteAddr().String()
    ch <- "You are " + who
    messages <- who + " have arrived"
    entering <- ch

    input := bufio.NewScanner(conn)
    for input.Scan(){
        messages <- who + ": " + input.Text()

        //update tick
        tick = time.Tick(60*time.Second)
    }

    // NOTE: ignoring potential errors from input.Err()

    leaving <- ch
    messages <- who + " has left"
    conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
    for msg := range ch{
        fmt.Fprintln(conn,msg) //NOTE: ignoring network errors
    }
}











