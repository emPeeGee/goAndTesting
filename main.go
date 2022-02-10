package main

import (
	"fmt"
	"go_and_test/mocking"
	"os"
	"time"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() { time.Sleep(1 * time.Second) }

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("There", "English"))

	sleeper := &DefaultSleeper{}
	mocking.Countdown(os.Stdout, sleeper)
}
