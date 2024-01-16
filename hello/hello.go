package main

import (
	"bufio"
	"example.com/greetings"
	"fmt"
	"log"
	"os"
)

func SayHello(name string) {
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	message, err := greetings.Hello(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

func SayManyHello(names []string) {
	log.SetPrefix("many-hello")
	log.SetFlags(0)

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	for name, greeting := range messages {
		fmt.Println("[name]", name, "\t", greeting)
	}
}

func main() {
	SayHello("Vanya")
	SayManyHello([]string{
		"Vanya",
		"Jenya",
		"Olya",
	})
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
