package main

import 
(
        "fmt"
        "strings"
)

func main() {
    fmt.Println("Introduzca su nombre:  ")
    var nombre string
    fmt.Scanln(&nombre)
    fmt.Printf("Hola %s, Bienvenido!", nombre)
    nombre = strings.TrimSpace(nombre)
}