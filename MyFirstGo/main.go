package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
	"math/rand"
	"math"
	"testing")

func main(){
	fmt.Println(uuid.New().String())
	fmt.Println("The time is: ", time.Now())
	fmt.Println("My favorite number is: ", rand.Intn(10))
	fmt.Printf("Now you have %g problems. \n ",math.Sqrt(7))
}