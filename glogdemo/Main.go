package main

import (
    "flag"
    "fmt"
    "github.com/golang/glog"
)

func main2() {
    flag.Set("logtostderr", "true")
    flag.Parse()
    glog.Info("This is an info log.")
    glog.Warning("This is a warning log.")
    glog.Error("This is an error log.")
    //glog.Fatal("This is a fatal log.")
    glog.Exit("This is an exit log.")
}

func main() {
    defer glog.Flush()

    glog.Info("This is an info log.")
    glog.Warning("This is a warning log.")
    glog.Error("This is an error log.")
    if glog.V(0) {
        fmt.Println("This is a fatal log.")
    }
    if glog.V(1) {
        fmt.Println("This is a error log....")
    }

}
