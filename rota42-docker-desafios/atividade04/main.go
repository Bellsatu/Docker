package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Uso: ./videoconverter <arquivo de entrada> <formato de saída>")
        os.Exit(1)
    }
    input := os.Args[1]
    format := os.Args[2]
    // Aqui você chamaria o ffmpeg ou seu processamento
    fmt.Printf("Convertendo %s para o formato %s...\n", input, format)
    // Simulação de processamento
}
