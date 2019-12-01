package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calcAllFuel(mass int) int {

	var fuel = (mass / 3) - 2

	if fuel > 0 {
		return fuel + calcAllFuel(fuel)
	} else {
		return 0
	}
}

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var reqFuelPartOne = 0
	var reqFuelPartTwo = 0

	for scanner.Scan() {
		toI, _ := strconv.Atoi(scanner.Text())
		reqFuelPartOne += ((toI / 3) - 2)
		reqFuelPartTwo += calcAllFuel(toI)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part One Answer:")
	fmt.Println(reqFuelPartOne)

	fmt.Println("Part Two Answer:")
	fmt.Println(reqFuelPartTwo)
}
