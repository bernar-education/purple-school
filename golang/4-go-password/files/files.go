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

func ReadFromFile() {
    data, err := os.ReadFile("./file.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(data)
}
