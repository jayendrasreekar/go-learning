package main

// import "fmt" single Import
import ("fmt";"os") // Multiple imports import statement


//os.Args - Array of Strings that has all the command line Args. 
// Accessing element form an Array is similar to other programming languages
// os.Args[0] or accessing a range of elements os.Args[m:n] where [m,n). m defaults
// to 0 and n defaults to len(os.Args)


// default value for int is 0 and string is ""
func main() {
	var sep , s string   // var one,two string 
	s = "***************** MESSAGE ***********************\n"
	for i:=1; i<len(os.Args) ; i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	s += "****************  END ***********************\n" 
	fmt.Println(s)
}