// Arquivo principal: main.go
package main

import (
	"fmt"
	"cachorro"
)

func main() {
	fmt.Println(cachorro.Idade(10))
}

// Arquivo cachorro.go
// package cachorro
//
// import "fmt"
//
// func Idade(anos int) string {
// 	return fmt.Sprintf("A idade do cachorro é %d anos", anos)
// } 