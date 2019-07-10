package main

import (
    "fmt"
    "github.com/user/stringutil"
)

func main() {
    // fmt.Println(stringutil.PopCount(15))   
    // fmt.Println(stringutil.PopCountFast(15))   

    //fmt.Println(stringutil.Basename("a/b/c.go"))
    //fmt.Println(stringutil.Basename2("a/b/c.go"))

    // fmt.Println(stringutil.HasPrefix("yingping.huang","yingping"))
    // fmt.Println(stringutil.HasPrefix("yingping.huang","ping"))
    // fmt.Println(stringutil.HasSuffix("yingping.huang","huang"))
    // fmt.Println(stringutil.HasSuffix("yingping.huang","ping"))
    // fmt.Println(stringutil.Contains("yingping.huang","ping"))

    fmt.Println(stringutil.Comma("6123456789101010145666777"))
    fmt.Println(stringutil.Comma3("6123456789101010145666777"))
    fmt.Println(stringutil.Comma3("-6123456789101010145666777"))
    fmt.Println(stringutil.Comma3("+6123456789101010145666777.124455676"))
}
