package main

import (
    "fmt"
    "github.com/pkg/errors"
    "os"
)

func main() {
    err := process()
    if err != nil {
        fmt.Printf("error: %+v\n", err)
    }
}

func process() error {
    file, err := openFile("test.txt")
    if err != nil {
        return errors.Wrap(err, "failed to open file")
    }
    defer file.Close()

    // process the file

    return nil
}

func openFile(name string) (*os.File, error) {
    _, err := os.Stat(name)
    if os.IsNotExist(err) {
        return nil, errors.Wrap(err, "file does not exist")
    }
    if err != nil {
        return nil, errors.Wrap(err, "failed to stat file")
    }

    file, err := os.Open(name)
    if err != nil {
        return nil, errors.Wrap(err, "failed to open file")
    }

    return file, nil
}
