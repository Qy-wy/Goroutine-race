package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func add_stars(matrix []string, row int, wg *sync.WaitGroup, counters []int) {
	defer wg.Done()

	for i := 0; i < 30; i++ {
		matrix[row] += "*"
		counters[row]++
		time.Sleep(time.Duration(rand.Intn(250-150+1)+150) * time.Millisecond)
	}
}

func print_stars(matrix []string, counters []int, counter *int) {
	fmt.Print("\033[H\033[2J")

	for i := range matrix {
		fmt.Println(matrix[i])
		if counters[i] == 30 {
			matrix[i] = fmt.Sprintf("%d goroutine ended %d", i, *counter)
			*counter++
			counters[i]++
		}
	}
}

func main() {
	rows := 5
	var wg sync.WaitGroup
	matrix := make([]string, rows)
	counters := make([]int, rows)
	counter := 1

	for i := 0; i < rows; i++ {
		wg.Add(1)
		matrix[i] = fmt.Sprintf("%d: ", i)
		go add_stars(matrix, i, &wg, counters)
	}

	go func() {
		for {
			print_stars(matrix, counters, &counter)
			time.Sleep(150 * time.Millisecond)
		}
	}()

	wg.Wait()

	fmt.Print("\033[H\033[2J")
	print_stars(matrix, counters, &counter)
}
