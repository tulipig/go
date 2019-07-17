package main

import(
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "time"
    "path/filepath"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
    flag.Parse()
    roots := flag.Args()
    if len(roots)==0{
        roots =[]string{"."}
    }

    fileSizes := make(chan int64)
    go func(){
        for _,root := range roots{
            walkDir(root,fileSizes)
        }
        close(fileSizes)
    }()

    var tick <-chan time.Time
    if *verbose{
        tick = time.Tick(500*time.Millisecond)
    }

    var nfiles, nbytes int64
loop:
    for{
        select{
            case size,ok := <-fileSizes:
                if !ok{
                    break loop //fileSizes was close
                }
                nfiles++;
                nbytes += size
            case <-tick:
                printDiskUsage(nfiles, nbytes)   
        }
    }

    printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
    fmt.Printf("%d files  %.1f MB\n", nfiles, float64(nbytes)/1e6)
}

func walkDir(dir string, fileSizes chan<- int64) {
    
    for _, entry := range directs(dir){
        if entry.IsDir(){
            subdir := filepath.Join(dir,entry.Name())
            walkDir(subdir, fileSizes)
        } else{
            //fmt.Println(entry.Name(), ": ", entry.Size())
            fileSizes <- entry.Size()
        }
    }
}

func directs(dir string) []os.FileInfo {
    entries, err := ioutil.ReadDir(dir)
    if err != nil{
        fmt.Fprintf(os.Stderr, "du1: %v\n", err)
        return nil
    }
    return entries
}