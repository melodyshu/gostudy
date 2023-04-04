package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
)

type Config struct {
    Name string `yaml:"name"`
    Age  int    `yaml:"age"`
}

func main() {
    // 解析YAML格式的数据
    yamlData := `
name: Alice
age: 20
`
    var conf Config
    err := yaml.Unmarshal([]byte(yamlData), &conf)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Name: %s, Age: %d\n", conf.Name, conf.Age)

    // 序列化Go结构体成YAML格式的数据
    conf.Name = "Bob"
    conf.Age = 30
    yamlBytes, err := yaml.Marshal(conf)
    if err != nil {
        panic(err)
    }
    fmt.Printf("YAML data:\n%s", string(yamlBytes))
}
