package main

import ("fmt";"os";"bufio")


// Keys can take any type whose values can be compared to "=="


func main() {
	counts := make(map[string]int) //creates an empty map with key string and value int
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { // reads new line from std input
		//*** input.Scan() also returns treue for empty line
		if input.Text() == "end" { // To exit the loop
			break
		}
		counts[input.Text()]++	
	}
	for line, n := range counts {
		fmt.Println(line, n)
	}

	// for line, n := range counts {
	// 	fmt.Println("%d\t%s\n",n,line)
	// }
}


// format specified are also known as verbs in Golang