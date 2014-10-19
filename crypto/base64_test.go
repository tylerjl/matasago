package matasago

import "testing"

func TestBase64(t *testing.T) {
    hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

    actual, err := Decode64(hex)

    if actual != expected || err != nil {
        t.Errorf("decode64(%s) = %v, want %v", hex, actual, expected)
    }
}
