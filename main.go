package main

import (
	"fmt"
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
	name := GetName()
	fmt.Printf("Welcome to the Casino, %s!", name)

	symbolArr := GenerateSymbolsArray(symbols)
	for balance > 0 {
		bet := GetBet(balance)
		if bet == 0 {
			fmt.Printf("You left with, $%d.\n", bet)
			break
		}
		balance -= bet
		spin := GetSpin(symbolArr, 3, 3)
		PrintSpin(spin)
		winningLines := CheckSpin(spin, multipliers)
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
