package matasago

import (
    "encoding/hex"
    "encoding/base64"
)

func decode64(hex_s string) (string, error) {

    decoded, err := hex.DecodeString(hex_s)
    if err != nil {
        return "", err
    }
    encoded := base64.StdEncoding.EncodeToString(decoded)
    return encoded, nil
}
