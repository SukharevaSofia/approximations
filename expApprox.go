package main

import "math"

func expApprox(n int, dataX, dataY []float64) Data {
	logY := make([]float64, n)
	for i := 0; i < n; i++ {
		logY[i] = math.Log(dataY[i])
	}
	_, linearApproximation := linealApprox(n, dataX, logY)
	bLin, aLin := linearApproximation.coefficents[0], linearApproximation.coefficents[1]
	a := math.Exp(aLin)
	b := bLin
	f := func(element float64) float64 { return a * math.Exp(b*element) }
	coefficents := []float64{a, b}
	phi := getPhi(dataX, f)
	eps := getEps(phi, dataY)
	meanSqDev := meanSquareDeviation(f, dataX, dataY)

	return Data{f, coefficents, meanSqDev,
		dataX, dataY, phi, eps, EXP}
}
