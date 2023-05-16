package main

import (
	"math"
)

func GetKramer(matrix [][]float64, b []float64) []float64 {
	n := len(matrix)
	detMatrix := det(matrix)
	x := make([]float64, 0)
	bMatrix := make([][]float64, n)
	for i := 0; i < n; i++ {
		bMatrix[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		clone(bMatrix, matrix)
		for j := 0; j < n; j++ {
			bMatrix[j][i] = b[j]
		}
		detB := det(bMatrix)
		x = append(x, detB/detMatrix)
	}
	return x
}

func clone(newMatrix, oldMatrix [][]float64) {
	for i := 0; i < len(oldMatrix); i++ {
		for j := 0; j < len(oldMatrix); j++ {
			newMatrix[i][j] = oldMatrix[i][j]
		}
	}
}

func det(matrix [][]float64) float64 {
	n := len(matrix)
	if n == 1 {
		return matrix[0][0]
	}
	sign := 1.0
	detMatrix := 0.0
	for j := 0; j < n; j++ {
		newMatrix := make([][]float64, 0)
		for i := 1; i < n; i++ {
			row := make([]float64, 0)
			for k := 0; k < n; k++ {
				if k != j {
					row = append(row, matrix[i][k])
				}
			}
			newMatrix = append(newMatrix, row)
		}
		detMatrix += sign * matrix[0][j] * det(newMatrix)
		sign *= -1
	}
	return detMatrix
}

func squareDeviationArr(xData, yData []float64) float64 {
	sum := 0.
	for i := 0; i < len(xData); i++ {
		sum += math.Pow(xData[i]-yData[i], 2)
	}
	return sum
}
func squareDeviation(xData []float64, y float64) float64 {
	sum := 0.
	for i := 0; i < len(xData); i++ {
		sum += math.Pow(xData[i]-y, 2)
	}
	return sum
}

func meanSquareDeviation(function func(element float64) float64, xData, yData []float64) float64 {
	data := getPhi(xData, function)
	return math.Sqrt(squareDeviationArr(data, yData) / float64(len(xData)))
}

// gets a list of X values and a function as arguments (e.g. a*x+b for a linear equation)
// calculates the Y value for every X and returns the list of Ys
func getPhi(x []float64, function func(element float64) float64) []float64 {
	phi := make([]float64, len(x))
	for i := 0; i < len(x); i++ {
		phi[i] = function(x[i])
	}
	return phi
}

func getEps(phiX, yData []float64) []float64 {
	eps := make([]float64, 0)
	for i := 0; i < len(phiX); i++ {
		eps = append(eps, phiX[i]-yData[i])
	}
	return eps
}
