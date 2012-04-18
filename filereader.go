package fileutil

import (
    "bufio"
    "bytes"
    "io"
    "os"
)

func GetLines(filename string) (chan string, error) {
    file, err := os.Open(filename)

    if err != nil {
        return nil, err
    }

    s := make(chan string)
    go _readLines(file, s)

    return s, nil
}

func _readLines(file io.Reader, s chan string) {
    buf := new(bytes.Buffer)
    reader := bufio.NewReader(file)
    var (
        part []byte
        prefix bool
        err error
    )

    for ; err == nil; {
        if part, prefix, err = reader.ReadLine(); err != nil {
            break
        }

        buf.Write(part)

        if !prefix {
            s <- buf.String()
            buf.Reset()
        }
    }

    close(s)
}

