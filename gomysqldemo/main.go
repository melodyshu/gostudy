package main

import (
    "github.com/go-mysql-org/go-mysql/mysql"
    "github.com/go-mysql-org/go-mysql/server"
    "log"
    "net"
)

func main() {
    // Listen for connections on localhost port 4000
    l, err := net.Listen("tcp", ":3306")
    if err != nil {
        log.Fatal(err)
    }

    // Accept a new connection once
    c, err := l.Accept()
    if err != nil {
        log.Fatal(err)
    }

    // Create a connection with user root and an empty password.
    // You can use your own handler to handle command here.
    //conn, err := server.NewConn(c, "root", "", server.EmptyHandler{})
    srv := server.NewServer("8.0.1", mysql.DEFAULT_COLLATION_ID, mysql.AUTH_CACHING_SHA2_PASSWORD, nil, nil)
    p := server.NewInMemoryProvider()
    p.AddUser("root", "123456")
    conn, err := server.NewCustomizedConn(c, srv, p, server.EmptyHandler{})
    if err != nil {
        log.Fatal(err)
    }

    // as long as the client keeps sending commands, keep handling them
    for {
        if err := conn.HandleCommand(); err != nil {
            log.Fatal(err)
        }
    }
}
