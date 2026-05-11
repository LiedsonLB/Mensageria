// cmd/convert_logo/main.go
package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Ler o arquivo da logo
	imageData, err := ioutil.ReadFile("assets/neoshop_logo.png")
	if err != nil {
		log.Fatal(err)
	}

	// Converter para base64
	base64Str := base64.StdEncoding.EncodeToString(imageData)
	
	// Gerar o código HTML
	htmlImg := fmt.Sprintf(`<img src="data:image/png;base64,%s" alt="NeoShop Logo" style="height: 50px; width: auto; margin-bottom: 15px;">`, base64Str)
	
	fmt.Println("=== COPIE O CÓDIGO ABAIXO ===")
	fmt.Println()
	fmt.Println(htmlImg)
}