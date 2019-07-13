package main

import (
    "io"
    "log"
    "fmt"
    "net"
    "os"
    "strings"
    "bufio"
)

func parseParam(addrs []string, mp map[string]string) {
    for _,addr := range addrs{
        s := strings.Split(addr,"=")
        if len(s)!=2{
            log.Fatal("err param")
            return
        }

        sp := strings.Split(s[1],":")
        if len(sp)!=2{
            log.Fatal("err param")
            return
        }

        //fmt.Println(s[0],sp[1])
        mp[s[0]] = sp[1]
    }
}

//./clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
func main() {
    ch := make(chan int)
    mp := make(map[string]string) //key:address, value:port
    addrs := os.Args[1:]
    parseParam(addrs, mp)
    if len(mp)<1{
        log.Fatal("err para")
        return
    }

    for addr,port := range mp{
        serv := "localhost:" + port
        Dail(addr, serv)
    } 

    <-ch  
}

func Dail(addr string, serv string) {   
    conn, err := net.Dial("tcp", serv)
    if err != nil {
        fmt.Println(err)
        log.Fatal(err)
    }
    
    //go mustCopy(os.Stdout, conn)
    
    go handleConn(addr, conn)
}

func handleConn(addr string, c net.Conn) {
    defer c.Close()
    reader := bufio.NewReader(c)
    for {
        msg, err := reader.ReadString('\n')
        if err != nil {
            c.Close()
            break
        }
        out := fmt.Sprintf("%10s: %s", addr,msg)
        fmt.Print(out)
    }
}

func mustCopy(dst io.Writer, src io.Reader) {
    if _, err := io.Copy(dst, src); err != nil {
        log.Fatal(err)
    }
}
