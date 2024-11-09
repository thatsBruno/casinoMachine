package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}
	//mulipliers := map[string]uint{
	//	"A": 20,
	//	"B": 10,
	//	"C": 5,
	//	"D": 2,
	//}
	balance := uint(200)

	symbolArr := generateSymbolsArray(symbols)

	getSpin(symbolArr, 3, 3)
	fmt.Println("Welcome to the Casino...")
	name := getName()
	fmt.Printf("Welcome to the Casino, %s!", name)

	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}
		balance -= bet
		fmt.Printf("You left with, $%d.\n", bet)
	}
}

func getRandomNumber(min, max int) int {
	randoNumber := rand.Intn(max-min+1) + min
	return randoNumber
}

func getSpin(reel []string, rows int, cols int) [][]string {
	result := make([][]string, 0)

	for i := 0; i < rows; i++ {
		result = append(result, []string{})
	}

	for col := 0; col < cols; col++ {
		selected := map[int]bool{}
		for row := 0; row < rows; row++ {
			for {
				randomIndex := getRandomNumber(0, len(reel)-1)
				_, exists := selected[randomIndex]
				if !exists {
					selected[randomIndex] = true
					result[row][col] = reel[randomIndex]
					break
				}
			}
		}
	}
	return result
}

func generateSymbolsArray(symbols map[string]uint) []string {
	var symbolArr []string

	for symbol, count := range symbols {
		for i := 0; i < int(count); i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}
	return symbolArr
}

func getName() string {
	name := ""

	fmt.Printf("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		log.Fatal(err)
	}
	return name
}

func getBet(balance uint) uint {
	var bet uint

	for true {
		fmt.Printf("Enter your bet: (balance = $%d ", balance)
		_, err := fmt.Scanln(&bet)
		if err != nil {
			log.Fatal(err)
		}
		if bet > balance {
			fmt.Println("Bet cannot be bigger than balance")
		} else {
			break
		}
	}
	return bet
}
