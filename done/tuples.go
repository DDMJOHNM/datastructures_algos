package main

import (
	"fmt"
)

//Tuples - finite list of sorted items  


func powerSeries(a int)(int,int){

	return a*a, a*a*a
}


func main(){

	var square int
	var cube int 
	
	square, cube = powerSeries(3)
	fmt.Print("square", square, "cube", cube)  

}

