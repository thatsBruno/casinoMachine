package main

import (
	"fmt"
	"log"
	"math/rand"
)

func PrintSpin(spin [][]string) {
	for _, row := range spin {
		for j, symbol := range row {
			fmt.Printf(symbol)
			if j != len(row)-1 {
				fmt.Print(" | ")
			}
		}
		fmt.Println()
	}
}

func GetRandomNumber(min, max int) int {
	randoNumber := rand.Intn(max-min+1) + min
	return randoNumber
}

func GetSpin(reel []string, rows int, cols int) [][]string {
	result := make([][]string, 0)
	for i := 0; i < rows; i++ {
		result = append(result, []string{})
	}

	for col := 0; col < cols; col++ {
		selected := map[int]bool{}
		for row := 0; row < rows; row++ {
			for {
				randomIndex := GetRandomNumber(0, len(reel)-1)
				_, exists := selected[randomIndex]
				if !exists {
					selected[randomIndex] = true
					result[row] = append(result[row], reel[randomIndex])
					break
				}
			}
		}
	}
	return result
}

func GenerateSymbolsArray(symbols map[string]uint) []string {
	var symbolArr []string

	for symbol, count := range symbols {
		for i := 0; i < int(count); i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}
	return symbolArr
}

func GetBet(balance uint) uint {
	var bet uint

	for {
		fmt.Printf("Enter your bet: (balance = $%d) ", balance)
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

func CheckSpin(spin [][]string, multiplier map[string]uint) []uint {
	var lines []uint

	for _, row := range spin {
		win := true
		checksSymbol := row[0]
		for _, symbol := range row[1:] {
			if checksSymbol != symbol {
				win = false
				break
			}
		}
		if win {
			lines = append(lines, multiplier[checksSymbol])
		} else {
			lines = append(lines, 0)
		}
	}
	return lines
}
