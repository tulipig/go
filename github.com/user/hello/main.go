package main

import (
    "fmt"
    //"log"
    //"github.com/user/stringutil"
    //"encoding/json"
    //"os"
)

// type Movie struct{
//     Title  string
//     Year   int `json:"released"`
//     Color  bool `json:"color,omitempty"`
//     Actors []string
// }
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

    // fmt.Println(stringutil.Comma("6123456789101010145666777"))
    // fmt.Println(stringutil.Comma3("6123456789101010145666777"))
    // fmt.Println(stringutil.Comma3("-6123456789101010145666777"))
    // fmt.Println(stringutil.Comma3("+6123456789101010145666777.124455676"))


    // var movies = []Movie{
    //     {Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
    //     {Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
    //     {Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
    // }

    // data,err := json.Marshal(movies)
    // if err!=nil{
    //     log.Fatalf("JSON Marshaling failed:%s", err)
    // }
    // fmt.Printf("%s\n", data)

    // data,err = json.MarshalIndent(movies,"","   ")
    // if err!=nil{
    //     log.Fatalf("JSON Marshaling failed:%s",err)
    // }
    // fmt.Printf("%s\n", data)

    // var titles []struct{ Title string }
    // if err := json.Unmarshal(data, &titles); err != nil {
    //     log.Fatalf("JSON unmarshaling failed: %s", err)
    // }
    // fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"


    // result,err := stringutil.SearchIssues(os.Args[1:])
    // if err != nil{
    //     log.Fatal(err)
    // }
    // fmt.Printf("%d issues:\n", result.TotalCount)
    // for _,item := range result.Items{
    //     fmt.Printf("%#-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
    // }


    // var intset stringutil.IntSet
    // fmt.Println(intset.Has(127))
    // intset.Add(127)
    // fmt.Println(intset.Has(127))
    // intset.Add(257)
    // fmt.Println(intset.Has(257))
    // intset.Add(1257)
    // fmt.Println(intset.String())

    // fmt.Println(intset.Len())

    // intset.Remove(257)
    // fmt.Println(intset.String())

    // intset2 := intset.Copy()
    // fmt.Println(intset2.String())

    // intset2.Clear()
    // fmt.Println(intset2.String())
    // fmt.Println(intset.String())

    // intset.AddAll(888,999,1000)
    // fmt.Println(intset.String())

    // intset.Clear()
    // intset2.Clear()

    // intset.AddAll(12,123,1234,12345)
    // intset2.AddAll(11,123,12345,123456)
    // // intset.UnionWith(intset2)
    // // fmt.Println(intset.String())

    // // intset.IntersectWith(intset2)
    // // fmt.Println(intset.String())

    // intset.DifferenceWith(intset2)
    // fmt.Println(intset.String())


    var c ByteCounter
    c.Write([]byte("hello wrold"))
    fmt.Println(c)

    c=0
    var name = "Dolly."
    fmt.Fprintf(&c, "hello,%s", name)
    fmt.Println(c)

}

type ByteCounter int
func (c *ByteCounter) Write(p []byte) (int,error) {
    *c += ByteCounter(len(p))
    return len(p), nil
}