package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	hourOfDay := time.Now().Hour()
	greeting, err := getGreeting(hourOfDay)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Print(greeting)
}

func getGreeting(hour int) (string, error) {
	var message string
	if hour < 7 {
		err := errors.New("too early for greetings")
		return message, err
	}
	if hour < 12 {
		message = "Good morning"
	} else if hour > 12 {
		message = "Good Afternoon"
	} else {
		message = "Good Evening"
	}
	return message, nil
}
