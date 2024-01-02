package main

import ("fmt";"os")

func main() {
	sep,s := " ",""
	for _,arg := range os.Args[1:]{
		s +=  sep + arg
		// underscore aka 'Blank Identifier' in golan is used to ignore any unused vars
		// as its illegal in Go to have vars that are not used 	
	}
	fmt.Println(s)
}


//var declaration
// s:= "" can only be used in fucntions
// var s string - takes default value as per type
// var s = "" - type is implicitly taken from value
// var s string = "" 


// fmt.Println(strings.Join(os.Args[1:],” “))

