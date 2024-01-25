package main

import (
	"fmt"
	"strings"
	"time"
)

func main(){	
	go mi_nombre_lentoooo("David")
	fmt.Println("Que Aburrido esperar")
	var wait string
	fmt.Scan(&wait)
}

func mi_nombre_lentoooo(name string){
	letras := strings.Split(name,"")
	for _,letra:=range(letras){
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}


}