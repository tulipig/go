package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
)

type dollars float32

func (d dollars) String() string {return fmt.Sprintf("$%.2f",d)}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
    for item,price := range db{
        fmt.Fprintf(w,"%s: %s\n", item,price)
    }
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
    item := req.URL.Query().Get("item")
    price,ok := db[item]
    if !ok{
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w,"no such item:%q\n", item)
        return
    }
    fmt.Fprintf(w,"%s\n",price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
    item := req.URL.Query().Get("item")
    oldprice,ok := db[item]
    if !ok{
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w,"no such item:%q\n", item)
        return
    }

    newprice := req.URL.Query().Get("price")
    if len(newprice)==0{
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w,"cannot parse %s.price\n", item)
        return
    }

    fmt.Println(newprice)
    fprice,_ := strconv.ParseFloat(newprice,32)
    db[item] = dollars(fprice)
    fmt.Fprintf(w,"have update %s.price from %s to %s\n",item, oldprice,db[item])
}

func main() {
    db := database{"shoes":50, "socks":5}
    //mux := http.NewServeMux()
    // mux.Handle("/list", http.HandleFunc(db.list)) //compile error
    // mux.Handle("/price", http.HandleFunc(db.price)) // compile error

    // mux.HandleFunc("/list",db.list)
    // mux.HandleFunc("/price",db.price)
    // log.Fatal(http.ListenAndServe("localhost:8000",mux))

    http.HandleFunc("/list",db.list)
    http.HandleFunc("/price",db.price)
    http.HandleFunc("/update",db.update)
    log.Fatal(http.ListenAndServe("localhost:8000",nil))

}