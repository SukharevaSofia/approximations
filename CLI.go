package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input() (int, []float64, []float64) {
	var inputString string
	var numberOfElements int
	var dataX, dataY []float64
	var lineX, lineY []string

	fmt.Print(INFO)

	fmt.Println(CHOOSE_INPUT)
	for {
		fmt.Scan(&inputString)
		if inputString == "T" || inputString == "t" || inputString == "F" || inputString == "f" {
			break
		}
		fmt.Println(CHOOSE_INPUT)
	}
	if inputString == "F" || inputString == "f" {
		f, _ := os.Open("data.txt")
		scanner := bufio.NewScanner(f)
		scanner.Scan()
		line := scanner.Text()
		numberOfElements, _ = strconv.Atoi(line)
		dataX = make([]float64, numberOfElements)
		dataY = make([]float64, numberOfElements)
		scanner.Scan()
		lineX = strings.Split(scanner.Text(), " ")
		scanner.Scan()
		lineY = strings.Split(scanner.Text(), " ")
		if (len(lineX) != numberOfElements) || (len(lineY) != numberOfElements) {
			fmt.Println(INPUT_ERR)
			os.Exit(0)
		}
		for i := 0; i < numberOfElements; i++ {
			dataX[i], _ = strconv.ParseFloat(lineX[i], 64)
			dataY[i], _ = strconv.ParseFloat(lineY[i], 64)
		}
	} else {
		fmt.Print(NUMBER_OF_DOTS)
		for {
			fmt.Scan(&numberOfElements)
			if numberOfElements <= 12 && numberOfElements >= 7 {
				break
			}
			fmt.Println(NUMBER_OF_DOTS)

		}
		dataX = make([]float64, numberOfElements)
		dataY = make([]float64, numberOfElements)
		for {
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println(INPUT_X)
			scanner.Scan()
			lineX = strings.Split(scanner.Text(), " ")
			fmt.Println(INPUT_Y)
			scanner.Scan()
			lineY = strings.Split(scanner.Text(), " ")
			if (len(lineY) == numberOfElements) || (len(lineX) == numberOfElements) {
				break
			}
			fmt.Println(INPUT_ERR)
		}
		for i := 0; i < numberOfElements; i++ {
			dataX[i], _ = strconv.ParseFloat(lineX[i], 64)
			dataY[i], _ = strconv.ParseFloat(lineY[i], 64)
		}

	}
	fmt.Println("Количество точек: ", numberOfElements)
	fmt.Println("X: ", dataX)
	fmt.Println("Y: ", dataY)
	return numberOfElements, dataX, dataY
}

func output() (*bufio.Writer, *os.File) {
	fmt.Println(CHOOSE_OUTPUT)
	var inputString string
	for {
		fmt.Scan(&inputString)
		if inputString == "T" || inputString == "t" || inputString == "F" || inputString == "f" {
			break
		}
		fmt.Println(CHOOSE_OUTPUT)
	}
	if inputString == "F" || inputString == "f" {
		f, _ := os.Create("output.txt")
		return bufio.NewWriter(f), f
	}
	return bufio.NewWriter(os.Stdout), os.Stdout

}
