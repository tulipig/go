// package main

// import (
// 	"fmt"
// 	"github.com/user/stringutil"
// )
// //comment
// func main(){
// 	fmt.Printf(stringutil.Reverse("hello,world."))
// 	fmt.Printf("\n")	
// }


//+++++++++++++++++++++++++++++

package main

import (
    "fmt"
)
//comment
func main(){
    go func(){
        for {
            fmt.Println("sub routine contiuing...")
        }
    }()
    fmt.Println("main routine exit")
}
