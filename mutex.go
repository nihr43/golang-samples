package main

import (
	"fmt"
	"sync"
)

func printThing(key string, things map[string]string) {
	fmt.Println(things[key])
}

func mutateThing(key string, things map[string]string, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	things[key] = "yellow"
	mu.Unlock()
	fmt.Println("thing", key, "has been mutated")
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	things := map[string]string{
		"A": "green",
		"B": "pink",
	}

	for k, _ := range things {
		printThing(k, things)
	}

	wg.Add(2)
	go mutateThing("A", things, &wg, &mu)
	go mutateThing("B", things, &wg, &mu)
	wg.Wait()

	for k, _ := range things {
		printThing(k, things)
	}
}
