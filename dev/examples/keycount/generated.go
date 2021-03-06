// The keycount command was automatically generated by Shenzhen Go.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func Count_words(input <-chan string, output chan<- string, result chan<- map[string]uint) {
	const multiplicity = 1

	const instanceNumber = 0

	m := make(map[string]uint)
	for in := range input {
		m[in]++
		if output != nil {
			output <- in
		}
	}
	result <- m
	if output != nil {
		close(output)
	}
	close(result)
}

func Get_words(words chan<- string) {
	const multiplicity = 1

	const instanceNumber = 0
	fmt.Println("Enter a line of text:")
	s, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}
	for _, word := range strings.Fields(s) {
		words <- word
	}
	close(words)

}

func Print_summary(result <-chan map[string]uint) {
	const multiplicity = 1

	const instanceNumber = 0
	fmt.Printf("Got results: %v\n", <-result)

}

func main() {

	results := make(chan map[string]uint, 0)
	words := make(chan string, 0)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		Count_words(words, nil, results)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		Get_words(words)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		Print_summary(results)
		wg.Done()
	}()

	// Wait for the end
	wg.Wait()
}
