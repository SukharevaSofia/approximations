package main

import "math"

func logApprox(n int, dataX, dataY []float64) Data {
	logX := make([]float64, n)
	for i := 0; i < n; i++ {
		logX[i] = math.Log(dataX[i])
	}
	_, linearApproximation := linealApprox(n, logX, dataY)
	bLin, aLin := linearApproximation.coefficents[0], linearApproximation.coefficents[1]
	a := aLin
	b := bLin
	f := func(element float64) float64 { return a*math.Log(element) + b }
	coefficents := []float64{a, b}
	phi := getPhi(dataX, f)
	eps := getEps(phi, dataY)
	meanSqDev := meanSquareDeviation(f, dataX, dataY)

	return Data{f, coefficents, meanSqDev,
		dataX, dataY, phi, eps, LOG}
}
