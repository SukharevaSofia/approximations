package main

import "math"

func powerApprox(n int, dataX, dataY []float64) Data {
	logX := make([]float64, n)
	logY := make([]float64, n)
	for i := 0; i < n; i++ {
		logX[i] = math.Log(dataX[i])
		logY[i] = math.Log(dataY[i])
	}
	_, linearApproximation := linealApprox(n, logX, logY)
	bLin, aLin := linearApproximation.coefficents[0], linearApproximation.coefficents[1]
	a := math.Exp(aLin)
	b := bLin
	f := func(element float64) float64 { return a * math.Pow(element, b) }
	coefficents := []float64{a, b}
	phi := getPhi(dataX, f)
	eps := getEps(phi, dataY)
	meanSqDev := meanSquareDeviation(f, dataX, dataY)

	return Data{f, coefficents, meanSqDev,
		dataX, dataY, phi, eps, POWER}
}
