package main

import (
        "fmt"
        "io"
        "bufio"
        "strings"
)

func NewReader(s *string) []byte{
    // string-->[]byte
    return []byte(*s)
}

func LimitReader(r io.Reader, n int64) io.Reader {
    //ioReader-->Reader
    br := bufio.NewReader(r)
    b := make([]byte,n)
    br.Read(b)

    // for i:=int64(0);i<n;i++{
    //     fmt.Printf("%c",b[i])
    // }
    // fmt.Printf("\n")

    //string-->ioReader; []byte->string
    return  strings.NewReader(string(b[:n]))
}

func sqlQuote(x interface{}) string {
    switch x := x.(type){
    case nil:
        return "NULL"
    case int,uint:
        return fmt.Sprintf("%d",x)

    case bool:
        if x{
            return "TRUE"
        }
        return "FALSE"
    case string:
        return x
    default:
        panic(fmt.Sprintf("unexpected type %T: %v", x, x))
    }
}

func main() {
    s := "hello world!"
    fmt.Println(string(NewReader(&s)))

    //string-->ioReader
    ss := strings.NewReader(s)
    out := LimitReader(ss,8)
    LimitReader(out,20)

    fmt.Println(sqlQuote("testing"))
    fmt.Println(sqlQuote(true))
    fmt.Println(sqlQuote(86))
    fmt.Println(sqlQuote([]byte(s)))
}
