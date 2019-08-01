package main

import(
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "time"
    "sync"
    "path/filepath"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")
var done = make(chan struct{})

func canceled() bool {
    select{
    case <-done:
        return true
    default:
        return false
    }
}

func main() {
    flag.Parse()
    roots := flag.Args()
    if len(roots)==0{
        roots =[]string{"."}
    }

    go func() {
        os.Stdin.Read(make([]byte, 1))
        close(done)
    }()

    fileSizes := make(chan int64)

    var n sync.WaitGroup
    for _,root := range roots{
        n.Add(1)
        go walkDir(root, &n, fileSizes)
    }

    go func () {
        n.Wait()
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
        case <-done:
            // Drain fileSizes to allow existing goroutines to finish.
            for range fileSizes{
                //do nothing
            }
            return
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

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {

    defer n.Done()

    if canceled(){
        return
    }

    for _, entry := range directs(dir){
        if entry.IsDir(){
            n.Add(1)
            subdir := filepath.Join(dir,entry.Name())
            go walkDir(subdir, n, fileSizes)
        } else{
            //fmt.Println(entry.Name(), ": ", entry.Size())
            fileSizes <- entry.Size()
        }
    }
}

var sema = make(chan struct{}, 20)

func directs(dir string) []os.FileInfo {
    select{
    case <-done:
        return nil
    case sema <- struct{}{}:
        //do nothing
    }
    
    defer func() {<-sema}()

    entries, err := ioutil.ReadDir(dir)
    if err != nil{
        fmt.Fprintf(os.Stderr, "du1: %v\n", err)
        return nil
    }
    return entries
}