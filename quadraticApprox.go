package main

import "math"

func quadraticApprox(n int, dataX, dataY []float64) Data {
	sx, sx2, sx3, sx4, sy, sxy, sxxy := 0., 0., 0., 0., 0., 0., 0.
	for i := 0; i < len(dataX); i++ {
		sx += dataX[i]
		sx2 += dataX[i] * dataX[i]
		sx3 += math.Pow(dataX[i], 3)
		sx4 += math.Pow(dataX[i], 4)
		sy += dataY[i]
		sxy += dataX[i] * dataY[i]
		sxxy += dataX[i] * dataX[i] * dataY[i]
	}
	matrix := [][]float64{{float64(n), sx, sx2}, {sx, sx2, sx3}, {sx2, sx3, sx4}}
	bMatrix := []float64{sy, sxy, sxxy}
	answKramer := GetKramer(matrix, bMatrix)
	a0, a1, a2 := answKramer[2], answKramer[1], answKramer[0]
	f := func(element float64) float64 { return a0*element*element + a1*element + a2 }
	coefficents := []float64{a2, a1, a0}
	meanSqDev := meanSquareDeviation(f, dataX, dataY)
	phi := getPhi(dataX, f)
	eps := getEps(phi, dataY)
	return Data{f, coefficents, meanSqDev,
		dataX, dataY, phi, eps, QUADRATIC}
}
