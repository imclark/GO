package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	names := []string{"Phil", "Noodles", "Barbaros"}
	var wg sync.WaitGroup
	wg.Add(len(names))
	for _, name := range names {
		go printName(name, &wg)
	}
	wg.Wait()
}

func printName(name string, wg *sync.WaitGroup) {
	result := 0.0
	for i := 0; i < 10000000; i++ {
		result += math.Pi * math.Sin(float64(len(name)))
	}
	fmt.Println("Name: ", name)
	wg.Done()
}
