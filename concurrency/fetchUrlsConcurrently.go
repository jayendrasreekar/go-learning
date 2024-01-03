//************* A goroutine is a concurrent function execution.
//************** A channel is a communication mechanism that allows
// one goroutine to pass values of a specified type to another goroutine.

//the main() function is a Go routine itself, we can spawn new Goroutines by using keyword “go”

package main 

import ( "fmt"
	"os"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	fmt.Println("****** Print once")

	start := time.Now()
	var ch = make(chan string) //declare a channel

	for _, url := range os.Args[1:]{
		go fetch(url,ch) // starts a go routine, makes async calls to fetch
	}

	for range os.Args[1:] {
		fmt.Println("?????receive from channel")
		fmt.Println(<-ch) // receive from channel
	}
	fmt.Printf("@@@@@@@@@@@@ %.2fs elapsed/n",time.Since(start).Seconds())
}


func fetch(url string, ch chan<- string){
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil{
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	fmt.Println("========send to channel")
	ch <- fmt.Sprintf("%.sfs %7d %s",secs,nbytes,url)

}



///////////********** Channel syntax
//var ch = make(chan string) // create a channel
// <-ch // receive from channel
// ch<- //Write to channel

//The execution is blocked when a goroutine either requests
// send or receive to the channel untill the corresponding receive or send is not received.

