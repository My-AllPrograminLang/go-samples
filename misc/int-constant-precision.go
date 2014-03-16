package main

import "fmt"

// works. but with 63 doesn't work... Is it a problem in fmt?
const bigint = 1<<62

func main() {
    fmt.Println(bigint)
}
