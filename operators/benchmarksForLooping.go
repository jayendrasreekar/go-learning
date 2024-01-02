package main
import ("fmt"; "os";"strings";"time")

func withForLoop(){
	sep,s := " ",""
	for _,arg := range os.Args[1:]{
		s += sep + arg
	}
	fmt.Println(s)
}

func withStringsJoin(){
	fmt.Println(strings.Join(os.Args[1:]," "))
}

func main(){
	ta_start := time.Now()
	withForLoop()
	ta_end := time.Now()
	fmt.Println(ta_end.Sub(ta_start)) 
	tb_start := time.Now()
	withStringsJoin()
	tb_end := time.Now()
	fmt.Println(tb_end.Sub(tb_start))
}