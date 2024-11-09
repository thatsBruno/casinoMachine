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
	multipliers := map[string]uint{
		"A": 20,
		"B": 10,
		"C": 5,
		"D": 2,
	}
	balance := uint(200)

	fmt.Println("Welcome to the Casino...")
	name := getName()
	fmt.Printf("Welcome to the Casino, %s!", name)

	symbolArr := generateSymbolsArray(symbols)
	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			fmt.Printf("You left with, $%d.\n", bet)
			break
		}
		balance -= bet
		spin := getSpin(symbolArr, 3, 3)
		printSpin(spin)
		winningLines := checkSpin(spin, multipliers)
		fmt.Println(winningLines)
		for i, multi := range winningLines {
			win := multi * bet
			balance += win
			if multi > 0 {
				fmt.Printf("You win %d, (%dx) on line #%d\n", win, multi, i+1)
			}
		}
	}
}

// methods
func checkSpin(spin [][]string, multiplier map[string]uint) []uint {
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

func printSpin(spin [][]string) {
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
					result[row] = append(result[row], reel[randomIndex])
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
