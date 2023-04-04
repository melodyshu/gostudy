package main

import (
    "log"
    "os"
    "runtime/pprof"
    "time"
)

func fibonacci(n int) int {
    if n < 2 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    f, err := os.Create("profile.pprof")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    // 开始分析 CPU 使用情况
    if err := pprof.StartCPUProfile(f); err != nil {
        log.Fatal(err)
    }
    defer pprof.StopCPUProfile()

    // 实际的应用程序代码
    for i := 0; i < 100; i++ {
        fibonacci(35)
        time.Sleep(100 * time.Millisecond)
    }
}
