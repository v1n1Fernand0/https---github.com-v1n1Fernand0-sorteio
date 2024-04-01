package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	numbers := getNumbersInRange(1000)
	selectedNumbers := getRandonNumbers(numbers, 10)
	countDown()
	fmt.Println("Aqui estão os números sorteados: ", selectedNumbers)
	fmt.Scanln()
}

func getNumbersInRange(numberRange int) []int {

	numbers := make([]int, numberRange)

	for i := 0; i < len(numbers); i++ {
		numbers[i] = i + 1
	}

	return numbers
}

func getRandonNumbers(numbers []int, pickNumber int) []int {

	if pickNumber > len(numbers) {
		pickNumber = len(numbers)
	}

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	shuffled := make([]int, len(numbers))

	copy(shuffled, numbers)

	r.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled[:pickNumber]
}

func countDown() {
	fmt.Println("Preparando os números...")

	for i := 1; i < 10; i++ {
		println("Revelando em... \n", i)
		time.Sleep(1 * time.Second)
	}
}
