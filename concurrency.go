package main

import (
    "fmt"
    "time"
)

type Result struct {
    Field1 string
    Field2 int
}

func producer(ch chan Result, d time.Duration, name string) {
    i := new(Result)
    i.Field1 = name
    i.Field2 = 0
    
    for {
        ch <- *i
        i.Field2++
        time.Sleep(d)
    }
}

func reader(out chan Result) {
    for x := range out {
        fmt.Println(x.Field1, x.Field2)
    }
}

func main() {
    ch := make(chan Result)
    out := make(chan Result)
    go producer(ch, 100*time.Millisecond, "shorty")
    go producer(ch, 1000*time.Millisecond, "long")
    go reader(out)
    for i := range ch {
        out <- i
    }
}