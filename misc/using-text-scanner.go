package main

import "bytes"
import "fmt"
import "text/scanner"


func main() {
    const src = "hello 1"
    var s scanner.Scanner
    s.Init(bytes.NewBufferString(src))
    tok := s.Scan()
    for tok != scanner.EOF {
        fmt.Println(scanner.TokenString(tok))
        // do something with tok
        tok = s.Scan()
    }
}
