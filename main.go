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

func PrintStars(matrix []string, ch chan int) {
	fmt.Print("\033[H\033[2J")

	for i := range matrix {
		select {
		case row := <-ch:
			fmt.Printf("Горутина %d завершилась!\n", row)
		default:
			fmt.Println(matrix[i])
			time.Sleep(150 * time.Millisecond)
		}
	}
}

func main() {
	rows := 5
	var wg sync.WaitGroup
	matrix := make([]string, rows)
	ch := make(chan int, rows)

	for i := 0; i < rows; i++ {
		wg.Add(1)
		matrix[i] = fmt.Sprintf("%d: ", i)
		go AddStars(matrix, i, &wg, ch)
	}

	go func() {
		for {
			PrintStars(matrix, ch)
		}
	}()

	wg.Wait()
	close(ch)

	for row := range ch {
		fmt.Printf("Горутина %d завершилась!\n", row)
	}
}
