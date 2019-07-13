package main
import (
    "io"
    "fmt"
    "log"
    "flag"
    "time"
    "net"
    "strconv"
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
    var port = flag.Int("p",8000,"port")
    flag.Parse()
    fmt.Println(*port)
    client := "localhost:" + strconv.Itoa(*port)

    listenner,err := net.Listen("tcp", client)
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


