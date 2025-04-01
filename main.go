package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func AddStars(matrix []string, row int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	for i := 0; i < 30; i++ {
		matrix[row] += "*"
		time.Sleep(time.Duration(rand.Intn(250-150+1)+150) * time.Millisecond)
	}

	ch <- row
}

func SetPlace(matrix []string, ch chan int) {
	for i := 0; i < len(matrix); i++ {
		select {
		case row := <-ch:
			matrix[row] = fmt.Sprintf("Goroutine №%d ended on the %d place!", row, i+1)
		}
	}
}

func PrintStars(matrix []string) {
	fmt.Print("\033[H\033[2J")

	for i := range matrix {
		fmt.Println(matrix[i])
	}

	time.Sleep(150 * time.Millisecond)
}

func main() {
	const rows = 5
	var wg sync.WaitGroup
	matrix := make([]string, rows)
	ch := make(chan int)

	for i := 0; i < rows; i++ {
		wg.Add(1)
		matrix[i] = fmt.Sprintf("%d: ", i)
		go AddStars(matrix, i, &wg, ch)
	}

	go func() {
		go SetPlace(matrix, ch)
		for {
			PrintStars(matrix)
		}
	}()

	wg.Wait()
	PrintStars(matrix)
}
