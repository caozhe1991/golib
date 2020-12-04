package main

import (
	"fmt"
	"xiandai/lib/vector3"
)

func main(){

	v1 := vector3.New(3,2,0)
	fmt.Println(vector3.New(2,5,0).Subtract(v1).Norm())
	//fmt.Println(vector3.New(5,2,0).Normalize().Norm())
}
