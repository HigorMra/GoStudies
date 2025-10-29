package basic

import (
	"fmt"
	"strings"
)

func String1() {
	var hello string = "Olá, mundo!"
	var question string = "Como vai?"

	//Concatenação
	var meet = hello + question
	fmt.Println(meet)
	fmt.Println(strings.ToUpper(meet))         //pacote strings
	fmt.Println(strings.Contains(meet, "vai")) //pacote strings
}
