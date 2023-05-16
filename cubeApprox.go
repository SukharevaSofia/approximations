package main

import "math"

func cubeApprox(n int, dataX, dataY []float64) Data {
	sx, sx2, sx3, sx4, sx5, sx6, sy, sxy, sx2y, sx3y := 0., 0., 0., 0., 0., 0., 0., 0., 0., 0.
	for i := 0; i < len(dataX); i++ {
		sx += dataX[i]
		sx2 += dataX[i] * dataX[i]
		sx3 += math.Pow(dataX[i], 3)
		sx4 += math.Pow(dataX[i], 4)
		sx5 += math.Pow(dataX[i], 5)
		sx6 += math.Pow(dataX[i], 6)
		sy += dataY[i]
		sxy += dataX[i] * dataY[i]
		sx2y += dataX[i] * dataX[i] * dataY[i]
		sx3y += math.Pow(dataX[i], 3) * dataY[i]
	}
	matrix := [][]float64{
		{sx6, sx5, sx4, sx3},
		{sx5, sx4, sx3, sx2},
		{sx4, sx3, sx2, sx},
		{sx3, sx2, sx, float64(n)}}
	matrixB := []float64{sx3y, sx2y, sxy, sy}
	answKramer := GetKramer(matrix, matrixB)
	a, b, c, d := answKramer[0], answKramer[1], answKramer[2], answKramer[3]
	f := func(element float64) float64 { return a*a*a*element + b*b*element + c*element + d }
	coefficents := []float64{a, b, c, d}
	meanSqDev := meanSquareDeviation(f, dataX, dataY)
	phi := getPhi(dataX, f)
	eps := getEps(phi, dataY)
	return Data{f, coefficents, meanSqDev, dataX,
		dataY, phi, eps, CUBIC}

}
