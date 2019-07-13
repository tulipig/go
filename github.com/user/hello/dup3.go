package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "strings"
)

func main() {
    counts := make(map[string] int)
    files := os.Args[1:]

    for _,arg := range files {
        data,err := ioutil.ReadFile(arg)
        if err!=nil {
            fmt.Println(os.Stderr, "dup3: %v\n", err)
            continue;
        }
        
        for _,line := range strings.Split(string(data),"\n"){
            counts[line]++
        }
    }

    for line, n := range counts{
        if(n>1){
           fmt.Printf("%d\t%s\n", n, line) 
        }
        
    }
}
