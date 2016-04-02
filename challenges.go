package main

import (
    "encoding/hex"
    "encoding/base64"
    "fmt"
    "os"
    "flag"
    "regexp"
    "bufio"
)

func main() {

    flag.Parse()
    if len(flag.Args()) == 0 {
        fmt.Println("Challenge number required.")
        os.Exit(1)
    } else {
        switch flag.Args()[0] {
        case "1.1":
            set1_challenge1()
        case "1.2":
            bitwise_xor()
        case "1.3":
            xor_cipher()
        case "1.4":
            find_byte()
        }
    }
}

func find_byte() {
    var (
        decoded []byte
        cracked []byte
        score int
        previous int
        best []byte
    )

    lines, readerr := readLines("4.txt")

    if readerr != nil {
        fmt.Println(readerr)
        return
    }

    for i := range lines {
        decoded, _ = hex.DecodeString(lines[i])
        cracked = CrackString(decoded)
        score = ScoreString(cracked)
        if score > previous {
            best = make([]byte, len(cracked))
            copy(best, cracked)
            previous = score
        }
    }

    fmt.Printf("%q\n", best)
}

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func xor_cipher() {
    ciphertext := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
    bytes, _ := hex.DecodeString(ciphertext)

    plaintext := CrackString(bytes)

    fmt.Printf("%q\n", plaintext)
}

func CrackString(s []byte) (plaintext []byte) {

    var (
        key byte
        score int
        previous int
    )

    plaintext  = make([]byte, len(s))
    guesswork := make([]byte, len(s))

    for {
        score = 0
        for i := range s {
            guesswork[i] = s[i] ^ key
        }
        score = ScoreString(guesswork)
        if score > previous {
            previous = score
            copy(plaintext, guesswork)
        }
        if key == 0xff { break } else { key += 1 }
    }

    return
}

func ScoreString(s []byte) (score int) {
    words, _  := regexp.Compile(" +[a-zA-Z]+ +")
    return len(words.FindAllString(string(s[:]), -1))
}

func bitwise_xor() {
    input := "1c0111001f010100061a024b53535009181c"
    key   := "686974207468652062756c6c277320657965"
    input_bytes, _ := hex.DecodeString(input)
    key_bytes, _ := hex.DecodeString(key)
    result := make([]byte, len(input_bytes))
    for i := range(result) {
        result[i] = input_bytes[i] ^ key_bytes[i]
    }
    fmt.Printf("%q\n", result)
}

func set1_challenge1() {
    str := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    decoded, err := hex.DecodeString(str)
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
    encoded := base64.StdEncoding.EncodeToString(decoded)
    fmt.Printf("%q\n", encoded)
}
