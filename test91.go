package main

import (
	"fmt";
	"math/rand";
	"time";
  "bufio"
  "os"
  "strings"
)

// START0 OMIT
type Message struct {
	str string
  msg string
	wait chan bool // HL
}
// STOP0 OMIT

func main() {

//	c := fanIn(boring("Joe"), boring("Ann"), boring("Bob") // HL
	c := fanIn(boring("Joe"))	 // HL

// START1 OMIT


	for i := 0; i < 3; i++ {
		msg1 := <-c; fmt.Println(msg1.str, msg1.msg)
//		msg2 := <-c; fmt.Println(msg2.str, msg2.msg)
//    msg3 := <-c; fmt.Println(msg3.str, msg3.msg)
		msg1.wait <- false
//		msg2.wait <- false
//    msg3.wait <- false
	}


/*
	var j int = 0

	for {
			if len(c) > 0 {
					msg1 := <-c; fmt.Println(msg1.str, msg1.msg)
					msg1.wait <- false
					j++
					if j > 3 {  break }
			}

				// q に溜まるまで他の事をしたい
	//			time.Sleep(1 * time.Second)
	//			fmt.Println("何か")
		}
*/

// STOP1 OMIT
	fmt.Println("You're all boring; I'm leaving.")
}

func StrStdin() (stringInput string) {
    fmt.Printf("Input your name : ")
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    stringInput = scanner.Text()

    stringInput = strings.TrimSpace(stringInput)
    return
}


func boring(msg string) <-chan Message { // Returns receive-only channel of strings. // HL
	c := make(chan Message)
// START2 OMIT
	waitForIt := make(chan bool) // Shared between all messages.
// STOP2 OMIT
	go func() { // We launch the goroutine from inside the function. // HL
		for i := 0; ; i++ {
        st_w := StrStdin()
        c <- Message{ fmt.Sprintf(" %s: %d", msg, i),st_w, waitForIt }
        time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
         <-waitForIt
    }
	}()
	return c // Return the channel to the caller. // HL
}


// START3 OMIT
func fanIn(inputs ... <-chan Message) <-chan Message { // HL
	c := make(chan Message)
	for i := range inputs {
		input := inputs[i] // New instance of &#39;input&#39; for each loop.
		go func() { for { c <- <-input } }()
	}
	return c
}
