package files

import (
    "fmt"
    "os"
)


func WriteToFile(content []byte, name string) {
    file, err := os.Create(name)
    if err != nil {
        fmt.Println(err)
    }
    _, err = file.Write(content)

    defer file.Close()

    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("File created successfully")
}

func ReadFromFile(name string) ([]byte, error) {
    data, err := os.ReadFile(name)
    if err != nil {
        return nil, err
    }
    return data, nil
}
