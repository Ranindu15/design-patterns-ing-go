package main

import (
	"fmt"
)

func main (){
	GetObject("Test Name 1")
	GetObject("Test Name 2")
	GetObject("Test Name 3")
	//fmt.Println(GetObject("Test Name 1").name)
	//fmt.Println(GetObject("Test Name 2").name)
	//fmt.Println(GetObject("Test Name 3").name)
}

type singleton struct{
	name string
}

var object *singleton

func GetObject(name string) *singleton{
	if object == nil {
		fmt.Println("creating")
			object = &singleton{
				name: name,
			}
	}else {
		fmt.Println("Created one instance", &object)
	}
	return object
}