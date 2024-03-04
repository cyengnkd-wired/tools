package main

import (
    "fmt"
    "bytes"
    "os"
    "strings"
)

func main() {
    outputFile, err := os.Create("output.txt")
    if err != nil {
        panic(err)
    }
    defer outputFile.Close()

    var outputBuffer bytes.Buffer

    hexContentOutput := func(hexContent string) {
        outputBuffer.WriteString(hexContent)
	fmt.Printf("%v", hexContent)
    }

    fmt.Printf("[+] File: output.txt created [+]\n\n")
    
    for x := range 256 {
        hex := strings.Replace(fmt.Sprintf("%#x", x), "0x", "\\x", -1)
	hexContentOutput(hex)
    }

    outputFile.WriteString(outputBuffer.String())
}
