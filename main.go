package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"math/rand"
)

var prompts = []string{
	"Yes?", 
	"Tell me something",
	"Enter something fun",
	"How may I be of service?",
	"Are you being served?",
	"What is the location?",
	"Whatchu want?",
	"YO!",
	"Que?",
}
const panicLine = "!!!"
const quitLine = "q"

type fn func()

func main() {
	start()
}

func start() {
	defer recoverMe(start)
	reader := bufio.NewReader(os.Stdin)
	for {
		prompt()
		in, _ := reader.ReadString('\n')
		in = in[:len(in)-1]

		if in == quitLine {
			exit()
		} else if in == panicLine {
			panic()
		} else {
			fmt.Printf("OK, '%s'\n", in)
		}

	}

}

func prompt() {
	fmt.Println(prompts[rand.Intn(len(prompts))])
}
func exit() {
	log.Println("Bye-Bye!")
	os.Exit(0)
}

func panic() {
	log.Panic("AAGGGHH!!!")
}

func panicOnError(err error, msg string) {
	if err != nil {
		log.Panic(msg)
	}
}

func recoverMe(f fn) {
	if r := recover(); r != nil {
		log.Printf("recovering %s (%s)", f, r)
		start()
	}
}
