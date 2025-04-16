package cloud

import (
    "fmt"
    "os"
)

type CloudDb struct {
    url string
}

func NewCloudDb(url string) *CloudDb {
    return &CloudDb{
        url: url,
    }
}


func (db *CloudDb) Read() ([]byte, error) {
    return []bytes{}, nil
}


func (db *CloudDb) Write(content []byte) {
}
