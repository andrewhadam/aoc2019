package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	// Open the file
	csvfile, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < len(record); i = i + 4 {
			fmt.Println("What is i?")
			fmt.Println(i)
			if val, _ := strconv.Atoi(record[i]); val == 1 {
				fmt.Println("Found a 1")
				x, _ := strconv.Atoi(record[i+1])
				y, _ := strconv.Atoi(record[i+2])

				a, _ := strconv.Atoi(record[x])
				b, _ := strconv.Atoi(record[y])
				var c = a + b
				var dest, _ = strconv.Atoi(record[i+3])

				fmt.Print("The number at i + 1: ")
				fmt.Println(x)

				fmt.Print("The value of i + 1: ")
				fmt.Println(a)

				fmt.Print("The number at i + 2: ")
				fmt.Println(y)

				fmt.Print("The value at i + 2 : ")
				fmt.Println(b)

				fmt.Print("The value of adding the value at i+1 and the value at i+2 together: ")
				fmt.Println(c)

				record[dest] = strconv.Itoa(c)
			} else if val2, _ := strconv.Atoi(record[i]); val2 == 2 {
				fmt.Println("Found a 2")
				x2, _ := strconv.Atoi(record[i+1])
				y2, _ := strconv.Atoi(record[i+2])
				dest2, _ := strconv.Atoi(record[i+3])

				a2, _ := strconv.Atoi(record[x2])
				b2, _ := strconv.Atoi(record[y2])
				var c2 = a2 * b2

				fmt.Print("The number at i + 1: ")
				fmt.Println(x2)

				fmt.Print("The value of i + 1: ")
				fmt.Println(a2)

				fmt.Print("The number at i + 2: ")
				fmt.Println(y2)

				fmt.Print("The value at i + 2 : ")
				fmt.Println(b2)

				fmt.Print("The value of adding the value at i+1 and the value at i+2 together: ")
				fmt.Println(c2)

				fmt.Print("Value of Dest: ")
				fmt.Println(record[dest2])

				record[dest2] = strconv.Itoa(c2)

				fmt.Print("New Value of Dest: ")
				fmt.Println(record[dest2])

			} else if val3, _ := strconv.Atoi(record[i]); val3 == 99 {
				fmt.Println("I should halt")
				break
			} else {
				fmt.Println("Something has gone horribly wrong")

			}
		}

		fmt.Println("Final value: ")
		fmt.Println(record[0])

	}
}
