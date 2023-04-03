package main

import (
    "fmt"
    "github.com/blacktear23/go-proxyprotocol"
    "log"
    "net"
)

func handleConn(conn net.Conn) {
    defer conn.Close()

    // Read data from the client
    data := make([]byte, 1024)
    n, err := conn.Read(data)
    if err != nil {
        log.Println(err)
        return
    }

    // Echo the data back to the client
    echo := fmt.Sprintf("Echo: %s", string(data[:n]))
    _, err = conn.Write([]byte(echo))
    if err != nil {
        log.Println(err)
        return
    }
}

func main() {
    l, err := net.Listen("tcp", ":8888")
    // Create a new TCP listener with proxy protocol support
    listener, err := proxyprotocol.NewListener(l, "*", 5, false)
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()

    // Start accepting connections
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println(err)
            continue
        }

        // Handle the connection
        go handleConn(conn)
    }
}
