package main
import (
    "io"
    "log"
    "time"
    "net"
)

func handleConn(c net.Conn) {
    defer c.Close()
    for{
        _,err := io.WriteString(c, time.Now().Format("15:04:05\n"))
        if err != nil{
            return
        }
        time.Sleep(1*time.Second)
    }
}

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
        //handleConn(conn)
        go handleConn(conn)
    }
}


