package main

import "fmt"

func main() {
	var texto string = "Ola" //Definição de variável, DEFINIÇÃO CURTA DE VARÍAVEL COM ":="
	var text bool            //Definição de variável
	texto = "Tchau"          //O script assume a última variável, ao inves de ler "Ola" ele le "Tchau"
	fmt.Println(texto)       //Exemplo de variável vazia, nesse caso nada
	fmt.Println(text)        //Exemplo de variável vazia, neste caso Falsez
}
