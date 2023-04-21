package cachingtest

import (
    "encoding/hex"
    "fmt"
    "testing"
)

var foobarPwdSHA2Hash, _ = hex.DecodeString("244124303035245e3e122f765160382d0b1e4d416f3f6752557c3c4d726e5a68756f4a547547682e536c504e6e76484a7854646b5452486d636e45464d61377a634147714144")

func TestCheckShaPasswordGood(t *testing.T) {
    fmt.Println(string(foobarPwdSHA2Hash))
    pwd := "123456"
    r, _ := CheckHashingPassword(foobarPwdSHA2Hash, pwd, AuthCachingSha2Password)
    fmt.Println(r)
}
