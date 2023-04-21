package main

import (
    "fmt"
    "github.com/shopspring/decimal"
)

func main() {
    var x1 float64 = 0.1
    var y1 float64 = 0.2
    var z1 float64 = x1 + y1
    fmt.Println(z1)
    // 创建一个十进制数
    x := decimal.NewFromFloat(1.23)
    // 加法
    y := decimal.NewFromFloat(4.56)
    z := x.Add(y)  // z = x + y
    fmt.Println(z) // 输出：5.79

    // 减法
    z = x.Sub(y)   // z = x - y
    fmt.Println(z) // 输出：-3.33

    // 乘法
    z = x.Mul(y)   // z = x * y
    fmt.Println(z) // 输出：5.6088

    // 除法
    z = x.Div(y)   // z = x / y
    fmt.Println(z) // 输出：0.269736842105263157894736842105263157894736842105263

    // 自定义精度
    a := decimal.NewFromFloat(1.0)
    b := decimal.NewFromFloat(3.0)
    c := a.Div(b).Round(3) // 保留 3 位小数
    fmt.Println(c)         // 输出：0.333

    // 比较操作
    d := decimal.NewFromFloat(1.23)
    fmt.Println(x.Equal(d))    // 输出：true
    fmt.Println(x.LessThan(y)) // 输出：true
}
