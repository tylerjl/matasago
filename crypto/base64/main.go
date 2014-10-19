package main

import (
    "fmt"
    "github.com/tylerjl/matasago/crypto"
)

func main() {
    hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    str, err := matasago.Decode64(hex)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("%q\n", str)
    }
}
