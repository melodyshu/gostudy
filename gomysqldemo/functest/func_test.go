package functest

import (
    "fmt"
    "testing"
)

func TestWriteInitialHandshake(t *testing.T) {
    data := make([]byte, 4)
    data = append(data, byte(0))
    fmt.Printf("%x\n", data)
    fmt.Println(len(data))
}

func HexToOctet(s string) []byte {
    old := []byte(s)
    var newpwd []byte
    for i := 1; i < len(old); i++ {
        tmp1 := charVal(old[i]) << 4
        i++
        tmp2 := charVal(old[i])
        newpwd = append(newpwd, tmp1|tmp2)
    }
    return newpwd
}

func charVal(X byte) byte {
    if X >= '0' && X <= '9' {
        return byte(X - '0')
    } else {
        if X >= 'A' && X <= 'Z' {
            return byte(X - 'A' + 10)
        } else {
            return byte(X - 'a' + 10)
        }
    }
}

func TestHex(t *testing.T) {
    s := "native"
    if "native" != s {
        fmt.Println("!=")
    }
}
